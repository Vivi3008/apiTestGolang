package activity

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	repoBil "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
	repoTr "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"
)

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
			actBl.Details = listBills[i].Description
		}
		listActivities = append(listActivities, actBl)
	}

	trRepo := repoTr.NewRepository(r.DB)

	listTransfers, err := trRepo.ListTransfer(ctx, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	if len(listTransfers) != 0 {
		var actTr activities.AccountActivity
		for i := 0; i < len(listTransfers); i++ {
			actTr.Type = activities.Transfer
			actTr.Amount = listTransfers[i].Amount
			actTr.CreatedAt = listTransfers[i].CreatedAt
			actTr.Details = listTransfers[i].AccountOriginId
		}
		listActivities = append(listActivities, actTr)
	}

	return listActivities, nil
}
