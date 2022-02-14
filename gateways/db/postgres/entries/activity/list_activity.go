package activity

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	repoBil "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
)

type DestinyAccount struct {
	AccountDestinationId string
	Name                 string
}

type DescriptionPayment struct {
	Description string
	Status      bills.Status
}

func (r Repository) ListActivity(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	blRepo := repoBil.NewRepository(r.DB)

	listBills, err := blRepo.ListBills(ctx, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	var listActivities = make([]activities.AccountActivity, 0)

	var activity activities.AccountActivity

	for i := range listBills {
		activity.Type = activities.Bill
		activity.Amount = listBills[i].Value
		activity.CreatedAt = listBills[i].CreatedAt
		activity.Details = DescriptionPayment{
			Description: listBills[i].Description,
			Status:      listBills[i].StatusBill,
		}
		listActivities = append(listActivities, activity)
	}

	listTransfersAcitivy, err := r.ListTransfersAccount(ctx, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	listActivities = append(listActivities, listTransfersAcitivy...)

	return listActivities, nil
}

func (r Repository) ListTransfersAccount(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	const statement = `SELECT 
		name, 
		tr.id,
		tr.account_destination_id,
		tr.amount,
		tr.created_at
	FROM accounts AS a JOIN transfers AS tr 
	ON a.id=tr.account_destination_id 
	WHERE tr.account_origin_id=$1
	ORDER BY tr.created_at DESC`

	var listActivities = make([]activities.AccountActivity, 0)

	rows, err := r.DB.Query(ctx, statement, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var details DestinyAccount
		var activity activities.AccountActivity
		err = rows.Scan(&details.Name, &details.AccountDestinationId, &activity.Amount, &activity.CreatedAt)

		if err != nil {
			return []activities.AccountActivity{}, err
		}

		activity.Details = details
		activity.Type = activities.Transfer
		listActivities = append(listActivities, activity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listActivities, nil
}
