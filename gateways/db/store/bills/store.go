package bills

import "github.com/Vivi3008/apiTestGolang/domain/entities/bills"

var source = "bills.json"

type BillStore struct {
	blStore []bills.Bill
	src     string
}

func NewBillStore() BillStore {
	bl := make([]bills.Bill, 0)

	return BillStore{
		blStore: bl,
		src:     source,
	}
}
