package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type UsecaseMock struct {
	OnListAll func(id string) ([]transfers.Transfer, error)
	OnCreate  func(trans transfers.Transfer) (transfers.Transfer, error)
	OnSave    func(trans transfers.Transfer) error
}

func (m UsecaseMock) ListTransfer(ctx context.Context, id string) ([]transfers.Transfer, error) {
	return m.OnListAll(id)
}

func (m UsecaseMock) CreateTransfer(ctx context.Context, trans transfers.Transfer) (transfers.Transfer, error) {
	return m.OnCreate(trans)
}

func (m UsecaseMock) SaveTransfer(ctx context.Context, trans transfers.Transfer) error {
	return m.OnSave(trans)
}
