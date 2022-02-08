package transfers

import "context"

type TransferMock struct {
	OnListAll func(id string) ([]Transfer, error)
	OnStore   func(trans Transfer) error
}

func (m TransferMock) ListTransfer(ctx context.Context, id string) ([]Transfer, error) {
	return m.OnListAll(id)
}

func (m TransferMock) SaveTransfer(ctx context.Context, trans Transfer) error {
	return m.OnStore(trans)
}
