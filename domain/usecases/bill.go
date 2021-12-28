package usecases

import "github.com/Vivi3008/apiTestGolang/store"

type Bills struct {
	storeBill store.BillStore
}

func CreateNewBill(store store.BillStore) Bills {
	return Bills{
		storeBill: store,
	}
}
