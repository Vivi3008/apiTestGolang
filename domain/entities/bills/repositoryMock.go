package bills

import "context"

type BillMock struct {
	OnStore   func(b Bill) error
	OnCreate  func(b Bill) (Bill, error)
	OnListAll func(id string) ([]Bill, error)
}

func (b BillMock) ListBills(ctx context.Context, id string) ([]Bill, error) {
	return b.OnListAll(id)
}

func (b BillMock) StoreBill(ctx context.Context, bl Bill) error {
	return b.OnStore(bl)
}

func (b BillMock) CreateBill(ctx context.Context, bl Bill) (Bill, error) {
	return b.OnCreate(bl)
}
