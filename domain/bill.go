package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Bill struct {
	Id            string
	AccountId     string
	Description   string
	Value         int64
	DueDate       time.Time
	ScheduledDate time.Time
	StatusBill    Status
}

type Status string

const (
	Pago      Status = "Pago"
	Agendado  Status = "Agendado"
	Negado    Status = "Negado"
	layoutIso        = "2006-01-02"
)

var (
	ErrEmpty       = errors.New("missing is data")
	ErrDateInvalid = errors.New("scheduled date cannot be before today")
)

func NewBill(bill Bill) (Bill, error) {
	if bill.Description == "" || bill.Value == 0 || bill.DueDate.IsZero() {
		return Bill{}, ErrEmpty
	}

	scheduled, err := verifyDate(bill.ScheduledDate)

	if err != nil {
		return Bill{}, err
	}

	return Bill{
		Id:            uuid.New().String(),
		AccountId:     bill.AccountId,
		Description:   bill.Description,
		Value:         bill.Value,
		DueDate:       bill.DueDate,
		ScheduledDate: scheduled,
		StatusBill:    Agendado,
	}, nil
}

func verifyDate(date time.Time) (time.Time, error) {
	if date.IsZero() {
		return time.Now(), nil
	}

	if date.Before(time.Now()) {
		return time.Now(), ErrDateInvalid
	}

	return date, nil
}
