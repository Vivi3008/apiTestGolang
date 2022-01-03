package bill

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

type BillUsecase struct {
	acRepo account.AccountUsecase
	blRepo bills.BillRepository
}

func NewBillUseCase(bl bills.BillRepository, ac account.AccountUsecase) BillUsecase {
	return BillUsecase{
		blRepo: bl,
		acRepo: ac,
	}
}
