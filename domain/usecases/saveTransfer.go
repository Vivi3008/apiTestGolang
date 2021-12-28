package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (tr Tranfers) SaveTransfer(trans transfers.Transfer) (transfers.Transfer, error) {
	err := tr.storeTransfer.StoreTransfer(trans)

	if err != nil {
		return transfers.Transfer{}, err
	}

	return trans, nil
}
