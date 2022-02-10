package activities

import (
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
)

type TypeActivity string

const (
	Transfer TypeActivity = "Transferência"
	Bill     TypeActivity = "Pagamento"
)

type AccountActivity struct {
	Type      TypeActivity
	Amount    int
	CreatedAt time.Time
	Details   interface{}
}

type ActivityUsecase struct {
	accRepo account.AccountUsecase
	trRepo  transfers.TranfersUsecase
	blRepo  bill.BillUsecase
}

func NewAccountAcitivyUsecase(acc account.AccountUsecase, tr transfers.TranfersUsecase, bl bill.BillUsecase) ActivityUsecase {
	return ActivityUsecase{acc, tr, bl}
}
