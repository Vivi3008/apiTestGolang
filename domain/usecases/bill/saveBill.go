package bill

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b BillUsecase) SaveBill(bill bills.Bill) (bills.Bill, error) {
	err := b.blRepo.StoreBill(bill)

	if err != nil {
		return bills.Bill{}, err
	}

	return bill, nil
}
