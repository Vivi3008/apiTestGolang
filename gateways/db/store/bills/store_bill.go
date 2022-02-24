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

	b.blStore = append(b.blStore, bill)

	err := store.StoreFile(b.blStore, b.src)
	if err != nil {
		return err
	}
	return nil
}
