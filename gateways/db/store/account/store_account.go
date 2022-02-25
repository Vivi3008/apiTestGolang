package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var (
	ErrIdNotExists  = errors.New("id does not exist")
	ErrCpfNotExists = errors.New("cpf does not exist")
	ErrCpfExists    = errors.New("this cpf already exists")
	ErrEmptyID      = errors.New("id not be empty")
)

func (a AccountStore) StoreAccount(ctx context.Context, account account.Account) error {
	if account.Id == "" {
		return ErrEmptyID
	}

	_, err := a.ListAccountByCpf(ctx, account.Cpf)

	switch err {
	case ErrCpfNotExists:
		a.accStore = append(a.accStore, account)
		err = store.StoreFile(a.accStore, a.Src)
		return err
	default:
		return ErrCpfExists
	}
}
