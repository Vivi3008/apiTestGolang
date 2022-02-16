package activities

import (
	"context"
)

type Usecase interface {
	ListActivity(ctx context.Context, accountId string) ([]AccountActivity, error)
}
type ActivityUsecase struct {
	actRepo AccountActivityRepository
}

func NewAccountActivityUsecase(act AccountActivityRepository) ActivityUsecase {
	return ActivityUsecase{actRepo: act}
}
