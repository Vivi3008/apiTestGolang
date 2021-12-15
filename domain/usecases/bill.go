package usecases

import "github.com/Vivi3008/apiTestGolang/store"

type Bills struct {
	storeBill store.BillStore
}

type MethodPayment string

const (
	Debit  MethodPayment = "Débito"
	Credit MethodPayment = "Crédito"
)

func CreateNewBill(store store.BillStore) Bills {
	return Bills{
		storeBill: store,
	}
}
