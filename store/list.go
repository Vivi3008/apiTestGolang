package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrIdNotExists = errors.New("Id inexistente")
)

func (a AccountStore) ListAll() ([]domain.Account, error) {
	var list []domain.Account
	for _, account := range a.accStore {
		list = append(list, account)
	}
	return list, nil
}

func (a AccountStore) ListOne(accountId domain.AccountId) (domain.Account, error) {
	listAll, _ := a.ListAll()

	var listOne domain.Account

	for _, account := range listAll {
		if string(accountId) == account.Id {
			listOne = account
		}
	}

	if listOne.Id == "" {
		return domain.Account{}, ErrIdNotExists
	} else {
		return listOne, nil
	}
}

func (tr TransferStore) ListTransfers(accountOriginId domain.AccountOriginId) ([]domain.Transfer, error) {
	var list []domain.Transfer

	for _, transfer := range tr.tranStore {
		if transfer.AccountOriginId == string(accountOriginId) {
			list = append(list, transfer)
		}
	}
	return list, nil
}
