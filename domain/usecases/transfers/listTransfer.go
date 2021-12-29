package transfers

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

func (s TranfersUsecase) ListTransfer(originId string) ([]transfers.Transfer, error) {
	_, err := s.accUsecase.ListAccountById(originId)

	if err != nil {
		return []transfers.Transfer{}, account.ErrIdNotExists
	}

	list, err := s.repo.ListTransfer(originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
