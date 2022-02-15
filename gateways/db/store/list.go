package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var (
	ErrIdNotExists  = errors.New("id does not exist")
	ErrCpfNotExists = errors.New("cpf does not exist")
)

func (a AccountStore) ListAllAccounts() ([]account.Account, error) {
	var list []account.Account
	for _, account := range a.accStore {
		list = append(list, account)
	}
	return list, nil
}

func (a AccountStore) ListAccountById(accountId string) (account.Account, error) {
	listAll, _ := a.ListAllAccounts()

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

func (a AccountStore) ListAccountByCpf(cpf string) (account.Account, error) {
	listAll, _ := a.ListAllAccounts()

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

func (tr TransferStore) ListTransfer(accountOriginId string) ([]transfers.Transfer, error) {
	transfers := make([]transfers.Transfer, 0)

	for _, transfer := range tr.tranStore {
		if accountOriginId == transfer.AccountOriginId {
			transfers = append(transfers, transfer)
		} else {
			return nil, ErrIdNotExists
		}
	}

	return transfers, nil
}

func (b BillStore) ListBills(accountOriginId string) ([]bills.Bill, error) {
	var list []bills.Bill

	for _, bill := range b.blStore {
		if accountOriginId == bill.AccountId {
			list = append(list, bill)
		}
	}
	return list, nil
}
