package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

type UsecaseMock struct {
	OnStore   func(b bills.Bill) error
	OnCreate  func(b bills.Bill) (bills.Bill, error)
	OnListAll func(id string) ([]bills.Bill, error)
}

func (b UsecaseMock) ListBills(ctx context.Context, id string) ([]bills.Bill, error) {
	return b.OnListAll(id)
}

func (b UsecaseMock) SaveBill(ctx context.Context, bl bills.Bill) error {
	return b.OnStore(bl)
}

func (b UsecaseMock) CreateBill(ctx context.Context, bl bills.Bill) (bills.Bill, error) {
	return b.OnCreate(bl)
}
