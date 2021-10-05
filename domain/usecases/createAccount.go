package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) CreateAccount(person domain.Account) (domain.Account, error) {
	account, err := domain.NewAccount(person)

	if err != nil {
		return domain.Account{}, err
	}

	err = a.store.StoreAccount(account)

	if err != nil {
		return domain.Account{}, nil
	}

	return account, nil
}
