package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) ListAllAccounts() ([]domain.Account, error) {
	list, err := a.store.ListAll()

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a Accounts) ListAccountById(id string) (domain.Account, error) {
	account, err := a.store.ListOne(id)

	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}
