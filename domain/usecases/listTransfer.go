package usecases

import (
	"errors"
	"fmt"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrInsufficientLimit = errors.New("Insufficient Limit")
)

func (s Tranfers) ListTransfer(originId domain.AccountId) ([]domain.Transfer, error) {
	list, err := s.storeTransfer.ListTransfers(originId)

	if err != nil {
		return nil, fmt.Errorf("Could not list all accounts: %v\n", err)
	}

	return list, nil
}
