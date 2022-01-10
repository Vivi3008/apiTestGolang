package transfers

type TransferMock struct {
	OnListAll func(id string) ([]Transfer, error)
	OnStore   func(trans Transfer) error
	OnCreate  func(trans Transfer) (Transfer, error)
}

func (m TransferMock) ListTransfer(id string) ([]Transfer, error) {
	return m.OnListAll(id)
}

func (m TransferMock) SaveTransfer(trans Transfer) error {
	return m.OnStore(trans)
}

func (m TransferMock) CreateTransfer(trans Transfer) (Transfer, error) {
	return m.OnCreate(trans)
}
