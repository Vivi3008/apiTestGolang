package bills

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

func (b BillStore) ListBills(ctx context.Context, accountOriginId string) ([]bills.Bill, error) {
	listBills, err := store.ReadFile(b.src, "bill")
	if err != nil {
		return []bills.Bill{}, err
	}

	for _, bl := range listBills.Bill {
		if bl.AccountId == accountOriginId {
			b.blStore = append(b.blStore, bl)
		}
	}

	return b.blStore, nil
}
