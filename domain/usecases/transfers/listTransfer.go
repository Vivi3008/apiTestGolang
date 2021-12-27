package transfers

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (s TranfersUsecase) ListTransfer(originId string) ([]transfers.Transfer, error) {
	list, err := s.repo.ListTransfer(originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
