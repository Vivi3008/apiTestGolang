package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func (b Bills) ListAllBills(accountId string) ([]domain.Bill, error) {
	list, err := b.storeBill.ListBills(accountId)

	if err != nil {
		return nil, errors.New("could not list all bills")
	}

	return list, nil
}
