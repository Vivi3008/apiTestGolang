package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyValues   = errors.New("origin id, destiny id can't be empty")
	ErrInvalidAmount = errors.New("invalid value")
)

type Transfer struct {
	Id                   string
	AccountOriginId      string
	AccountDestinationId string
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(tr Transfer) (Transfer, error) {
	if tr.AccountOriginId == "" || tr.AccountDestinationId == "" {
		return Transfer{}, ErrEmptyValues
	}

	if tr.Amount <= 0 {
		return Transfer{}, ErrInvalidAmount
	}

	return Transfer{
		Id:                   uuid.New().String(),
		AccountOriginId:      tr.AccountOriginId,
		AccountDestinationId: tr.AccountDestinationId,
		Amount:               tr.Amount,
		CreatedAt:            time.Now(),
	}, nil

}
