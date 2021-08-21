package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrEmptyID = errors.New("id não pode ser vazio")
)

type AccountStore struct {
	accStore map[string]domain.Account
}

func NewAccountStore() AccountStore {
	as := make(map[string]domain.Account)

	return AccountStore{
		accStore: as,
	}
}

func (a AccountStore) StoreAccount(account domain.Account) error {
	if account.id == "" {
		return ErrEmptyID
	}

	a.accStore[account.id] = account
	return nil
}
