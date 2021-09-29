package store

import "github.com/Vivi3008/apiTestGolang/domain"

func (tr TransferStore) StoreTransfer(transaction domain.Transfer) error {
	if transaction.Id == "" {
		return ErrEmptyID
	}

	tr.tranStore[transaction.Id] = transaction
	return nil
}
