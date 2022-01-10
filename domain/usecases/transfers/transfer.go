package transfers

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

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
