package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

func (tr TransferStore) SaveTransfer(transaction transfers.Transfer) error {
	if transaction.Id == "" {
		return ErrEmptyID
	}

	tr.tranStore[transaction.Id] = transaction
	return nil
}
