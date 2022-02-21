package bills

import "context"

type BillMock struct {
	OnStore   func(b Bill) error
	OnListAll func(id string) ([]Bill, error)
}

func (b BillMock) ListBills(ctx context.Context, id string) ([]Bill, error) {
	return b.OnListAll(id)
}

func (b BillMock) StoreBill(ctx context.Context, bl Bill) error {
	return b.OnStore(bl)
}
