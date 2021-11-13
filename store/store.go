package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

type AccountStore struct {
	accStore map[string]domain.Account
}

type TransferStore struct {
	tranStore map[string]domain.Transfer
}

type BillStore struct {
	blStore map[string]domain.Bill
}

func NewAccountStore() AccountStore {
	as := make(map[string]domain.Account)

	return AccountStore{
		accStore: as,
	}
}

func NewTransferStore() TransferStore {
	tr := make(map[string]domain.Transfer)

	return TransferStore{
		tranStore: tr,
	}
}

func NewBillStore() BillStore {
	bl := make(map[string]domain.Bill)

	return BillStore{
		blStore: bl,
	}
}
