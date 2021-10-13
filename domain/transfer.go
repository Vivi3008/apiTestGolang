package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	Id                   string
	AccountOriginId      string
	AccountDestinationId string
	Amount               float64
	CreatedAt            time.Time
}

func NewTransfer(tr Transfer) (Transfer, error) {

	return Transfer{
		Id:                   uuid.New().String(),
		AccountOriginId:      tr.AccountOriginId,
		AccountDestinationId: tr.AccountDestinationId,
		Amount:               tr.Amount,
		CreatedAt:            time.Now(),
	}, nil

}
