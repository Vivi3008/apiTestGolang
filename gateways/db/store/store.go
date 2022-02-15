package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

type AccountStore struct {
	accStore map[string]account.Account
}

type TransferStore struct {
	tranStore map[string]transfers.Transfer
}

type BillStore struct {
	blStore map[string]bills.Bill
}

func NewAccountStore() AccountStore {
	as := make(map[string]account.Account)

	return AccountStore{
		accStore: as,
	}
}

func NewTransferStore() TransferStore {
	tr := make(map[string]transfers.Transfer)

	return TransferStore{
		tranStore: tr,
	}
}

func NewBillStore() BillStore {
	bl := make(map[string]bills.Bill)

	return BillStore{
		blStore: bl,
	}
}
