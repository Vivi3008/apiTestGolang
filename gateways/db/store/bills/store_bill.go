package bills

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

func (b BillStore) StoreBill(ctx context.Context, bill bills.Bill) error {
	if bill.Id == "" {
		return ErrEmptyID
	}

	listBills, err := store.ReadFile(b.Src, "bill")
	if err != nil {
		return err
	}

	listBills.Bill = append(listBills.Bill, bill)

	err = store.StoreFile(listBills.Bill, b.Src)
	if err != nil {
		return err
	}
	return nil
}
