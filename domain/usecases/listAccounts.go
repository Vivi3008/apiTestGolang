package usecases

import (
	"fmt"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) ListAllAccounts() ([]domain.Account, error) {
	list, err := a.store.ListAll()

	if err != nil {
		return nil, fmt.Errorf("Could not list all accounts: %v\n", err)
	}

	return list, nil
}

func (a Accounts) ListAccountById(id domain.AccountId) (domain.Account, error) {
	account, err := a.store.ListOne(id)

	if err != nil {
		return domain.Account{}, fmt.Errorf("Could not list account: %v\n", err)
	}

	return account, nil
}
