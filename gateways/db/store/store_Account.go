package store

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrCpfExists = errors.New("this cpf already exists")
)

func (a AccountStore) StoreAccount(ctx context.Context, account account.Account) error {
	if account.Id == "" {
		return ErrEmptyID
	}

	_, err := a.ListAccountByCpf(account.Cpf)

	if err == ErrCpfNotExists {
		a.accStore[account.Id] = account
		return nil
	} else {
		return ErrCpfExists
	}
}
