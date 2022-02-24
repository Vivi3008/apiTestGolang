package store

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type TransferStore struct {
	tranStore map[string]transfers.Transfer
}

func NewTransferStore() TransferStore {
	tr := make(map[string]transfers.Transfer)

	return TransferStore{
		tranStore: tr,
	}
}
