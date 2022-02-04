package transfers

import "context"

type TransferRepository interface {
	ListTransfer(ctx context.Context, id string) ([]Transfer, error)
	SaveTransfer(ctx context.Context, tr Transfer) error
}
