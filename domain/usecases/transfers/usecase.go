package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type Usecase interface {
	ListTransfer(ctx context.Context, accountId string) (transfers.Transfer, error)
	SaveTransfer(ctx context.Context, transfer transfers.Transfer) (transfers.Transfer, error)
	CreateTransfer(ctx context.Context, transfer transfers.Transfer) (transfers.Transfer, error)
}
