package activity

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (at ActivityStore) ListActivity(ctx context.Context, accountId string) ([]bills.Bill, error) {
	bill, err := at.billStore.ListBills(ctx, accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	return bill, nil
}
