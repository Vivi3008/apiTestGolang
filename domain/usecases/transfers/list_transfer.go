package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

func (s TranfersUsecase) ListTransfer(ctx context.Context, originId string) ([]transfers.Transfer, error) {
	_, err := s.accUsecase.ListAccountById(ctx, originId)

	if err != nil {
		return []transfers.Transfer{}, account.ErrIdNotExists
	}

	list, err := s.repo.ListTransfer(context.Background(), originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
