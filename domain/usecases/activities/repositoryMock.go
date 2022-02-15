package activities

import "context"

type AcitivitiesMock struct {
	OnListActivities func(accountId string) ([]AccountActivity, error)
}

func (a AcitivitiesMock) ListActivity(ctx context.Context, accountId string) ([]AccountActivity, error) {
	return a.OnListActivities(accountId)
}
