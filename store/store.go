package store

import (
	"errors"
)

var (
	ErrEmptyID = errors.New("id não pode ser vazio")
)

type AccountStore struct {
	accStore map[string]Account
}

func NewAccount() AccountStore {
	as := make(map[string]Account)

	return AccountStore{
		accStore: as,
	}
}

func (a AccountStore) StoreAccount(account Account) error {
	if account.id == "" {
		return ErrEmptyID
	}

	a.accStore[account.id] = account
	return nil
}
