package store

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b BillStore) StoreBill(bill bills.Bill) error {
	if bill.Id == "" {
		return ErrEmptyID
	}

	b.blStore[bill.Id] = bill
	return nil
}
