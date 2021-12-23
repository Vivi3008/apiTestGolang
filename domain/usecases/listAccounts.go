package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (a AccountUsecase) ListAllAccounts() ([]account.Account, error) {
	list, err := a.accs.ListAllAccounts()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a AccountUsecase) ListAccountById(id string) (account.Account, error) {
	acc, err := a.accs.ListAccountById(id)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
