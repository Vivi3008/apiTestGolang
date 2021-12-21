package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

var (
	ErrInsufficientLimit = errors.New("insufficient Limit")
)

func (s Tranfers) ListTransfer(originId string) ([]transfers.Transfer, error) {
	list, err := s.storeTransfer.ListTransfers(originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
