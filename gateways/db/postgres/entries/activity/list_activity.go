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

	if len(listBills) != 0 {
		var actBl activities.AccountActivity
		for i := 0; i < len(listBills); i++ {
			actBl.Type = activities.Bill
			actBl.Amount = listBills[i].Value
			actBl.CreatedAt = listBills[i].ScheduledDate
			actBl.Details = DescriptionPayment{
				Description: listBills[i].Description,
				Status:      listBills[i].StatusBill,
			}
			listActivities = append(listActivities, actBl)
		}
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
		tr.account_destination_id,
		tr.amount,
		tr.created_at
	FROM accounts AS a JOIN transfers AS tr 
	ON a.id=tr.account_destination_id 
	WHERE tr.account_origin_id=$1
	ORDER BY tr.created_at ASC`

	var listActivities = make([]activities.AccountActivity, 0)

	rows, err := r.DB.Query(ctx, statement, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	for rows.Next() {
		var details DestinyAccount
		var activity activities.AccountActivity
		err = rows.Scan(&details.Name, &details.AccountDestinationId, &activity.Amount, &activity.CreatedAt)
		activity.Details = details
		activity.Type = activities.Transfer

		if err != nil {
			return []activities.AccountActivity{}, err
		}
		listActivities = append(listActivities, activity)
	}
	return listActivities, nil
}
