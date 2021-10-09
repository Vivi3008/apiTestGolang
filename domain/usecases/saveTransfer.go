package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (tr Tranfers) SaveTransfer(trans domain.Transfer) (domain.Transfer, error) {
	err := tr.storeTransfer.StoreTransfer(trans)

	if err != nil {
		return domain.Transfer{}, err
	}

	return trans, nil
}
