package account

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrIdNotExists      = errors.New("id doesn't exists")
	ErrListAccountEmpty = errors.New("list account is empty")
)

func (a AccountUsecase) ListAllAccounts() ([]account.Account, error) {
	list, err := a.repo.ListAllAccounts()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a AccountUsecase) ListAccountById(id string) (account.Account, error) {
	list, _ := a.repo.ListAllAccounts()

	var accRes account.Account

	if len(list) == 0 {
		return account.Account{}, ErrListAccountEmpty
	}

	for _, acc := range list {
		if acc.Id == id {
			accRes = acc
		}
	}

	if accRes.Id == "" {
		return account.Account{}, ErrIdNotExists
	}

	return accRes, nil
}
