package domain

import (
	"testing"
	"time"
)

func TestNewBill(t *testing.T) {
	t.Run("Should create a new bill successfully", func(t *testing.T) {
		layoutIso := "2006-01-02"
		dueDate, _ := time.Parse(layoutIso, "2021-12-31")

		data := Bill{
			Description: "Conta de Luz",
			Value:       267.65,
			DueDate:     dueDate,
		}

		newBill, err := NewBill(data)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		dataAgendadaAtual := newBill.ScheduledDate.Format(layoutIso)

		if dataAgendadaAtual != actualDate.Format(layoutIso) {
			t.Errorf("Expected scheduled Date is actual day, got %v", newBill.ScheduledDate)
		}

		if newBill.StatusBill != Agendado {
			t.Errorf("Expected status bill is %s, got %s", Agendado, newBill.StatusBill)
		}
	})

	t.Run("Should create a new Bill with future scheduled date", func(t *testing.T) {
		layoutIso := "2006-01-02"
		dueDate, _ := time.Parse(layoutIso, "2021-12-31")
		scheduledDate, _ := time.Parse(layoutIso, "2021-11-12")

		billScheduled := Bill{
			Description:   "Conta de Internet",
			Value:         150,
			DueDate:       dueDate,
			ScheduledDate: scheduledDate,
		}

		newBillScheduled, err := NewBill(billScheduled)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		dateNewBillScheduled := newBillScheduled.ScheduledDate.Format(layoutIso)

		if dateNewBillScheduled != scheduledDate.Format(layoutIso) {
			t.Errorf("Expected %v, got %v", dateNewBillScheduled, scheduledDate.Format(layoutIso))
		}
	})

	t.Run("Should create a bill if scheduled date is less than today with actual date", func(t *testing.T) {
		layoutIso := "2006-01-02"
		dueDate, _ := time.Parse(layoutIso, "2021-12-31")
		scheduledDate, _ := time.Parse(layoutIso, "2021-11-08")

		billScheduled := Bill{
			Description:   "Conta de Agua",
			Value:         90,
			DueDate:       dueDate,
			ScheduledDate: scheduledDate,
		}

		newBillScheduled, err := NewBill(billScheduled)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		dateNewBillScheduled := newBillScheduled.ScheduledDate.Format(layoutIso)

		if actualDate.Format(layoutIso) != dateNewBillScheduled {
			t.Errorf("Expected %v, got %v", actualDate.Format(layoutIso), dateNewBillScheduled)
		}
	})
}
