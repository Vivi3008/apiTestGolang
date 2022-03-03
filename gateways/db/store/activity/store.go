package activity

import (
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	acStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/account"
	blStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/bills"

	trStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/transfers"
)

type ActivityStore struct {
	activity      []activities.AccountActivity
	billStore     blStore.BillStore
	transferStore trStore.TransferStore
	accountStore  acStore.AccountStore
}

func NewAccountActivity() ActivityStore {
	act := make([]activities.AccountActivity, 0)
	bl := blStore.NewBillStore()
	tr := trStore.NewTransferStore()
	ac := acStore.NewAccountStore()

	return ActivityStore{
		activity:      act,
		billStore:     bl,
		transferStore: tr,
		accountStore:  ac,
	}
}
