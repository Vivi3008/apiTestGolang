package transfers

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var (
	ErrIdNotExists  = errors.New("id does not exist")
	ErrCpfNotExists = errors.New("cpf does not exist")
)

func (tr TransferStore) ListTransfer(id string) ([]transfers.Transfer, error) {
	var list = make([]transfers.Transfer, 0)

	listTransfer, err := store.ReadFile(tr.Src, "transfer")
	if err != nil {
		return []transfers.Transfer{}, err
	}

	for _, transfer := range listTransfer.Transfer {
		if id == transfer.AccountOriginId {
			list = append(list, transfer)
		} else {
			return nil, ErrIdNotExists
		}
	}

	return list, nil
}
