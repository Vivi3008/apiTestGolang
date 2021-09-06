package store

import (
	"errors"
	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrIdNotExists = errors.New("nome inexistente")
)

func (a AccountStore) ListAll() ([]domain.Account, error) {
	var list []domain.Account
	for _, account := range a.accStore {
		list = append(list, account)
	}
	return list, nil
}

func (a AccountStore) ListOne(name string) (domain.Account, error) {
	listAll, _ := a.ListAll()

	var listOne domain.Account

	for _, account := range listAll {
		if account.Name == name {
			listOne = account
		}
	}

	if listOne.Name == "" {
		return domain.Account{}, ErrIdNotExists
	} else {
		return listOne, nil
	}
}
