package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrCpfExists = errors.New("This cpf already exists")
)

func (a Accounts) CreateAccount(person domain.Account) (domain.Account, error) {
	account, err := domain.NewAccount(person)

	if err != nil {
		return domain.Account{}, err
	}

	accounts, _ := a.ListAllAccounts()

	for _, acc := range accounts {
		if account.Cpf == acc.Cpf {
			return domain.Account{}, ErrCpfExists
		}
	}

	err = a.store.StoreAccount(account)

	if err != nil {
		return domain.Account{}, nil
	}

	return account, nil
}
