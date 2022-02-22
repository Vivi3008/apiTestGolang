package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var (
	ErrIdNotExists  = errors.New("id does not exist")
	ErrCpfNotExists = errors.New("cpf does not exist")
)

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
