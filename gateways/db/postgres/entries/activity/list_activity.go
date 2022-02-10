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

	var listActivities []activities.AccountActivity

	for i := 0; i < len(listBills); i++ {
		listActivities[i].Type = activities.Bill
		listActivities[i].Amount = listBills[i].Value
		listActivities[i].CreatedAt = listBills[i].ScheduledDate
		listActivities[i].Details = listBills[i].Description
	}

	trRepo := repoTr.NewRepository(r.DB)

	listTransfers, err := trRepo.ListTransfer(ctx, accountId)

	if err != nil {
		return []activities.AccountActivity{}, err
	}

	for i := 0; i < len(listTransfers); i++ {
		listActivities[i].Type = activities.Transfer
		listActivities[i].Amount = listTransfers[i].Amount
		listActivities[i].CreatedAt = listTransfers[i].CreatedAt
		listActivities[i].Details = listTransfers[i].AccountOriginId
	}

	return listActivities, nil
}
