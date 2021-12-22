package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrCpfExists = errors.New("this cpf already exists")
)

func (a Accounts) CreateAccount(person account.Account) (account.Account, error) {
	acc, err := account.NewAccount(person)

	if err != nil {
		return account.Account{}, err
	}

	accounts, _ := a.ListAllAccounts()

	for _, ac := range accounts {
		if acc.Cpf == ac.Cpf {
			return account.Account{}, ErrCpfExists
		}
	}

	err = a.accs.StoreAccount(acc)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
