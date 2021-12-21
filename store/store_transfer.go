package store

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (tr TransferStore) StoreTransfer(transaction transfers.Transfer) error {
	if transaction.Id == "" {
		return ErrEmptyID
	}

	tr.tranStore[transaction.Id] = transaction
	return nil
}
