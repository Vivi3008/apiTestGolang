package store

import "github.com/Vivi3008/apiTestGolang/domain"

func (b BillStore) StoreBill(bill domain.Bill) error {
	if bill.Id == "" {
		return ErrEmptyID
	}

	b.blStore[bill.Id] = bill
	return nil
}
