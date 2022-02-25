package activity

import (
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	blStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/bills"
	trStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/transfers"
)

type ActivityStore struct {
	activity      []activities.AccountActivity
	billStore     blStore.BillStore
	transferStore trStore.TransferStore
}

func NewAccountActivity() ActivityStore {
	act := make([]activities.AccountActivity, 0)
	bl := blStore.NewBillStore()
	tr := trStore.NewTransferStore()

	return ActivityStore{
		activity:      act,
		billStore:     bl,
		transferStore: tr,
	}
}
