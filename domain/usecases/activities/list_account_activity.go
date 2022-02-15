package activities

import "context"

func (a ActivityUsecase) ListActivity(ctx context.Context, accountId string) ([]AccountActivity, error) {
	listActivities, err := a.actRepo.ListActivity(ctx, accountId)

	if err != nil {
		return []AccountActivity{}, err
	}

	return listActivities, nil
}
