package transfers

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

func (tr TransferStore) SaveTransfer(ctx context.Context, transfer transfers.Transfer) error {
	if transfer.Id == "" {
		return ErrEmptyID
	}

	listTransfer, err := store.ReadFile(tr.Src, "transfer")
	if err != nil {
		return err
	}

	listTransfer.Transfer = append(listTransfer.Transfer, transfer)

	err = store.StoreFile(listTransfer.Transfer, tr.Src)
	if err != nil {
		return err
	}

	return nil
}
