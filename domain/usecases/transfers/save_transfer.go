package transfers

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (tr TranfersUsecase) SaveTransfer(trans transfers.Transfer) error {
	err := tr.repo.SaveTransfer(trans)

	if err != nil {
		return err
	}

	return nil
}
