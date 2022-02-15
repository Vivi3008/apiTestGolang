package activities

import "context"

type AccountActivityRepository interface {
	ListActivity(ctx context.Context, accountId string) ([]AccountActivity, error)
}
