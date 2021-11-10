package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Bill struct {
	Id            string
	Description   string
	Value         float64
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
	actualDate = time.Now()
	ErrEmpty   = errors.New("Missing is data")
)

func NewBill(bill Bill) (Bill, error) {
	if bill.Description == "" || bill.Value == 0 || bill.DueDate.IsZero() {
		return Bill{}, ErrEmpty
	}

	scheduled := verifyDate(bill.ScheduledDate)

	return Bill{
		Id:            uuid.New().String(),
		Description:   bill.Description,
		Value:         bill.Value,
		DueDate:       bill.DueDate,
		ScheduledDate: scheduled,
		StatusBill:    Agendado,
	}, nil
}

func verifyDate(date time.Time) time.Time {
	if date.IsZero() || date.Before(actualDate) {
		return actualDate
	}

	return date
}
