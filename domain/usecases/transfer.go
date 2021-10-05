package usecases

import "github.com/Vivi3008/apiTestGolang/store"

type Tranfers struct {
	storeTransfer store.TransferStore
}

func SaveNewTransfer(storeTr store.TransferStore) Tranfers {
	return Tranfers{
		storeTransfer: storeTr,
	}
}
