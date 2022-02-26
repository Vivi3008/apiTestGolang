package transfers

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var source = "./storage/transfers.json"

type TransferStore struct {
	tranStore []transfers.Transfer
	Src       string
}

func NewTransferStore() TransferStore {
	tr := make([]transfers.Transfer, 0)

	return TransferStore{
		tranStore: tr,
		Src:       source,
	}
}
