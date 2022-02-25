package account

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

func (a AccountStore) UpdateAccount(ctx context.Context, balance int, id string) (account.Account, error) {
	var updatedAccounts = make([]account.Account, 0)

	newAccount, err := a.ListAccountById(ctx, id)
	if err != nil {
		return account.Account{}, err
	}
	newAccount.Balance = balance

	accounts, err := a.ListAllAccounts(ctx)
	if err != nil {
		return account.Account{}, err
	}

	for _, ac := range accounts {
		if ac.Id == id {
			ac = newAccount
		}
		updatedAccounts = append(updatedAccounts, ac)
	}

	err = store.StoreFile(updatedAccounts, a.Src)
	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
