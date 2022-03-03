package bills

import (
	"context"
	"sort"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

func (b BillStore) ListBills(ctx context.Context, accountOriginId string) ([]bills.Bill, error) {
	listBills, err := store.ReadFile(b.Src, "bill")
	if err != nil {
		return []bills.Bill{}, err
	}

	for _, bl := range listBills.Bill {
		if bl.AccountId == accountOriginId {
			b.blStore = append(b.blStore, bl)
		}
	}

	b.OrderListBillsByDateDesc()

	return b.blStore, nil
}

func (b BillStore) OrderListBillsByDateDesc() {
	sort.Slice(b.blStore, func(i, j int) bool {
		return b.blStore[i].CreatedAt.After(b.blStore[j].CreatedAt)
	})
}
