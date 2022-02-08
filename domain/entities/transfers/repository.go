package transfers

import "context"

type TransferRepository interface {
	ListTransfer(ctx context.Context, originAccId string) ([]Transfer, error)
	SaveTransfer(ctx context.Context, tr Transfer) error
}
