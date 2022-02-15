package activities

import (
	"context"
)

type Usecase interface {
	ListActivity(ctx context.Context, accountId string) ([]AccountActivity, error)
}
