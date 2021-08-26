package store

import "github.com/Vivi3008/apiTestGolang/domain"

func (a AccountStore) StoreAccount(account domain.Account) error {
	if account.Id == "" {
		return ErrEmptyID
	}

	a.accStore[account.Id] = account
	return nil
}
