package bills

type BillMock struct {
	OnStore   func(b Bill) error
	OnListAll func(id string) ([]Bill, error)
}

func (b BillMock) ListBills(id string) ([]Bill, error) {
	return b.OnListAll(id)
}

func (b BillMock) StoreBill(bl Bill) error {
	return b.OnStore(bl)
}
