package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (tr TranfersUsecase) SaveTransfer(ctx context.Context, trans transfers.Transfer) error {
	err := tr.repo.SaveTransfer(ctx, trans)

	if err != nil {
		return err
	}

	return nil
}
