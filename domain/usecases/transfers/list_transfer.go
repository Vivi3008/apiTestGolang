package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (s TranfersUsecase) ListTransfer(ctx context.Context, originId string) ([]transfers.Transfer, error) {
	list, err := s.repo.ListTransfer(ctx, originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
