package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrCpfExists = errors.New("this cpf already exists")
)

func (a AccountUsecase) CreateAccount(ctx context.Context, person account.Account) (account.Account, error) {
	accounts, _ := a.repo.ListAllAccounts(ctx)

	for _, ac := range accounts {
		if person.Cpf == ac.Cpf {
			return account.Account{}, ErrCpfExists
		}
	}

	acc, err := account.NewAccount(person)

	if err != nil {
		return account.Account{}, err
	}

	err = a.repo.StoreAccount(ctx, acc)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
