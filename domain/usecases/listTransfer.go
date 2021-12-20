package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrInsufficientLimit = errors.New("insufficient Limit")
)

func (s Tranfers) ListTransfer(originId string) ([]account.Transfer, error) {
	list, err := s.storeTransfer.ListTransfers(originId)

	if err != nil {
		return nil, err
	}

	return list, nil
}
