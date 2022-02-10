package bills

import "context"

type BillRepository interface {
	ListBills(ctx context.Context, accountId string) ([]Bill, error)
	StoreBill(ctx context.Context, bill Bill) error
}
