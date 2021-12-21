package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b Bills) SaveBill(bill bills.Bill) (bills.Bill, error) {
	err := b.storeBill.StoreBill(bill)

	if err != nil {
		return bills.Bill{}, err
	}

	return bill, nil
}
