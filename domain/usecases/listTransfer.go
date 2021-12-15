package usecases

import (
	"errors"
	"fmt"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrInsufficientLimit = errors.New("insufficient Limit")
)

func (s Tranfers) ListTransfer(originId string) ([]domain.Transfer, error) {
	list, err := s.storeTransfer.ListTransfers(originId)

	fmt.Printf("Lista da conta %v\n", list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
