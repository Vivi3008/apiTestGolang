package activity

import (
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	blStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/bills"
)

type ActivityStore struct {
	activity  []activities.AccountActivity
	billStore blStore.BillStore
}

func NewAccountAcitivity() ActivityStore {
	act := make([]activities.AccountActivity, 0)
	bl := blStore.NewBillStore()

	return ActivityStore{
		activity:  act,
		billStore: bl,
	}
}
