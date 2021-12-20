package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrIdNotExists = errors.New("id does not exist")
)

func (a AccountStore) ListAll() ([]account.Account, error) {
	var list []account.Account
	for _, account := range a.accStore {
		list = append(list, account)
	}
	return list, nil
}

func (a AccountStore) ListOne(accountId string) (account.Account, error) {
	listAll, _ := a.ListAll()

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

func (tr TransferStore) ListTransfers(accountOriginId string) ([]account.Transfer, error) {
	transfers := make([]account.Transfer, 0)

	for _, transfer := range tr.tranStore {
		if accountOriginId == transfer.AccountOriginId {
			transfers = append(transfers, transfer)
		} else {
			return nil, ErrIdNotExists
		}
	}

	return transfers, nil
}

func (b BillStore) ListBills(accountOriginId string) ([]domain.Bill, error) {
	var list []domain.Bill

	for _, bill := range b.blStore {
		if accountOriginId == bill.AccountId {
			list = append(list, bill)
		}
	}
	return list, nil
}
