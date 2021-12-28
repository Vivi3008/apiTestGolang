package store

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (a AccountStore) StoreAccount(account account.Account) error {
	if account.Id == "" {
		return ErrEmptyID
	}

	a.accStore[account.Id] = account
	return nil
}
