package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var ErrOpenFile = errors.New("error to open file")

func (a AccountStore) ListAllAccounts(ctx context.Context) ([]account.Account, error) {
	var list = make([]account.Account, 0)

	data, err := store.ReadFile(a.src, "account")

	if err != nil {
		return []account.Account{}, ErrOpenFile
	}

	list = append(list, data.Account...)

	return list, nil
}

func (a AccountStore) ListAccountById(ctx context.Context, accountId string) (account.Account, error) {
	listAll, _ := a.ListAllAccounts(ctx)

	var listOne account.Account

	for _, account := range listAll {
		if string(accountId) == account.Id {
			listOne = account
		}
	}

	if listOne.Id == "" {
		return account.Account{}, ErrIdNotExists
	} else {
		return listOne, nil
	}
}

func (a AccountStore) ListAccountByCpf(ctx context.Context, cpf string) (account.Account, error) {
	listAll, _ := a.ListAllAccounts(ctx)

	var listOne account.Account

	for _, account := range listAll {
		if account.Cpf == cpf {
			listOne = account
		}
	}

	if listOne.Id == "" {
		return account.Account{}, ErrCpfNotExists
	} else {
		return listOne, nil
	}
}
