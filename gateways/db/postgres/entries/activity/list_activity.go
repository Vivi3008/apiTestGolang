package activity

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
)

type DestinyAccount struct {
	AccountDestinationId interface{}
	Name                 interface{}
}

type DescriptionPayment struct {
	Description interface{}
	Status      interface{}
}

func (r Repository) ListActivity(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	/* 	blRepo := repoBil.NewRepository(r.DB)

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
	   	} */

	listActivities, err := r.ListTransfersAccount(ctx, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	return listActivities, nil
}

func (r Repository) ListTransfersAccount(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	const statement = `SELECT 
	name, 
	tr.account_destination_id, 
	tr.amount,tr.created_at, 
	null as description, 
	null as status
FROM accounts AS a JOIN transfers AS tr 
ON a.id=tr.account_destination_id 	
WHERE tr.account_origin_id=$1
UNION ALL
SELECT 
	null as name, 
	tr.account_destination_id, 
	value as amount, 
	b.created_at, 
	description, 
	status 
FROM bills AS b 
JOIN transfers AS tr on b.account_id=tr.account_origin_id
WHERE tr.account_origin_id=$2
ORDER BY created_at DESC`

	var listActivities = make([]activities.AccountActivity, 0)

	rows, err := r.DB.Query(ctx, statement, accountId, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var detailsTransfers DestinyAccount
		var detailsBills DescriptionPayment
		var activity activities.AccountActivity
		err = rows.Scan(&detailsTransfers.Name,
			&detailsTransfers.AccountDestinationId,
			&activity.Amount,
			&activity.CreatedAt,
			&detailsBills.Description,
			&detailsBills.Status)

		if err != nil {
			return []activities.AccountActivity{}, err
		}

		if detailsBills.Status == nil && detailsBills.Description == nil {
			activity.Type = activities.Transfer
			activity.Details = detailsTransfers
		} else {
			activity.Type = activities.Bill
			activity.Details = detailsBills
		}

		listActivities = append(listActivities, activity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listActivities, nil
}
