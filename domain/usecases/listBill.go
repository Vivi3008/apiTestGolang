package usecases

import (
	"fmt"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func (b Bills) ListAllBills(accountId domain.AccountId) ([]domain.Bill, error) {
	list, err := b.storeBill.ListBills(accountId)

	if err != nil {
		return nil, fmt.Errorf("Could not list all bills: %v\n", err)
	}

	return list, nil
}
