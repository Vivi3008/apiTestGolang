package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

type Usecase interface {
	ListTransfer(ctx context.Context, accountId string) ([]transfers.Transfer, error)
	SaveTransfer(ctx context.Context, transfer transfers.Transfer) error
	CreateTransfer(ctx context.Context, transfer transfers.Transfer) (transfers.Transfer, error)
}

type TranfersUsecase struct {
	repo       transfers.TransferRepository
	accUsecase account.AccountUsecase
}

func NewTransferUsecase(tr transfers.TransferRepository, acc account.AccountUsecase) TranfersUsecase {
	return TranfersUsecase{
		repo:       tr,
		accUsecase: acc,
	}
}
