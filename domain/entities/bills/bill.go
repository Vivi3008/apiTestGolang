package bills

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Bill struct {
	Id            string
	AccountId     string
	Description   string
	Value         int
	DueDate       time.Time
	ScheduledDate time.Time
	StatusBill    Status
	CreatedAt     time.Time
}

type Status string

const (
	Pago     Status = "Pago"
	Agendado Status = "Agendado"
	Negado   Status = "Negado"
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
		CreatedAt:     time.Now(),
	}, nil
}

func verifyDate(date time.Time) (time.Time, error) {
	if date.IsZero() || date.Before(time.Now().UTC().Truncate(24*time.Hour)) {
		return time.Now(), nil
	}

	return date, nil
}
