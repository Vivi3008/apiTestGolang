package bill

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b BillUsecase) ListAllBills(accountId string) ([]bills.Bill, error) {
	_, err := b.acRepo.ListAccountById(accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	list, err := b.blRepo.ListBills(accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	return list, nil
}
