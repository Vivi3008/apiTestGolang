package usecases

import "github.com/Vivi3008/apiTestGolang/domain"

func (b Bills) SaveBill(bill domain.Bill) (domain.Bill, error) {
	err := b.storeBill.StoreBill(bill)

	if err != nil {
		return domain.Bill{}, err
	}

	return bill, nil
}
