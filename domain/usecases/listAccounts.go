package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (a Accounts) ListAllAccounts() ([]account.Account, error) {
	list, err := a.ListAllAccounts()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a Accounts) ListAccountById(id string) (account.Account, error) {
	acc, err := a.ListAccountById(id)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
