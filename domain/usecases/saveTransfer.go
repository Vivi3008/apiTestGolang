package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (tr Tranfers) SaveTransfer(trans account.Transfer) (account.Transfer, error) {
	err := tr.storeTransfer.StoreTransfer(trans)

	if err != nil {
		return account.Transfer{}, err
	}

	return trans, nil
}
