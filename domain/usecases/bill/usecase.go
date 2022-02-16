package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

type Usecase interface {
	CreateBill(ctx context.Context, bill bills.Bill) (bills.Bill, error)
	SaveBill(ctx context.Context, bill bills.Bill) error
	ListAllBills(ctx context.Context, accountId string) ([]bills.Bill, error)
}

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
