package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

type Usecase interface {
	CreateBill(ctx context.Context, bill bills.Bill) (bills.Bill, error)
	SaveBill(ctx context.Context, bill bills.Bill) error
	ListAllBills(ctx context.Context, accountId string) ([]bills.Bill, error)
}
