package store

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (tr TransferStore) StoreTransfer(transaction account.Transfer) error {
	if transaction.Id == "" {
		return ErrEmptyID
	}

	tr.tranStore[transaction.Id] = transaction
	return nil
}
