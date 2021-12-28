package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b Bills) ListAllBills(accountId string) ([]bills.Bill, error) {
	list, err := b.storeBill.ListBills(accountId)

	if err != nil {
		return nil, errors.New("could not list all bills")
	}

	return list, nil
}
