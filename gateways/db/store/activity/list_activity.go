package activity

import (
	"context"
	"sort"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
)

func (at ActivityStore) ListActivity(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	listActivities, err := at.ListBillActivity(ctx, accountId)
	if err != nil {
		return []activities.AccountActivity{}, err
	}

	listActivitiesTransfer, err := at.ListTransferActivity(ctx, accountId)
	if err != nil {
		return []activities.AccountActivity{}, err
	}

	listActivities = append(listActivities, listActivitiesTransfer...)

	at.activity = listActivities
	at.OrderListByDateDesc()

	return at.activity, nil
}

func (at ActivityStore) ListBillActivity(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	var listActivities = make([]activities.AccountActivity, 0)

	listBills, err := at.billStore.ListBills(ctx, accountId)
	if err != nil {
		return []activities.AccountActivity{}, err
	}

	for _, bl := range listBills {
		activity := activities.AccountActivity{
			Type:      activities.Bill,
			Amount:    bl.Value,
			CreatedAt: bl.CreatedAt,
			Details: activities.DescriptionPayment{
				Description: bl.Description,
				Status:      bl.StatusBill,
			},
		}
		listActivities = append(listActivities, activity)
	}

	return listActivities, nil
}

func (at ActivityStore) ListTransferActivity(ctx context.Context, accountId string) ([]activities.AccountActivity, error) {
	var listActivities = make([]activities.AccountActivity, 0)

	listTransfers, err := at.transferStore.ListTransfer(ctx, accountId)
	if err != nil {
		return []activities.AccountActivity{}, err
	}

	for _, tr := range listTransfers {
		accountDestiny, err := at.accountStore.ListAccountById(ctx, tr.AccountDestinationId)
		if err != nil {
			return []activities.AccountActivity{}, err
		}

		activity := activities.AccountActivity{
			Type:      activities.Transfer,
			Amount:    tr.Amount,
			CreatedAt: tr.CreatedAt,
			Details: activities.DestinyAccount{
				AccountDestinationId: tr.AccountDestinationId,
				Name:                 accountDestiny.Name,
			},
		}
		listActivities = append(listActivities, activity)
	}
	return listActivities, nil
}

func (at ActivityStore) OrderListByDateDesc() {
	sort.Slice(at.activity, func(i, j int) bool {
		return at.activity[i].CreatedAt.After(at.activity[j].CreatedAt)
	})
}
