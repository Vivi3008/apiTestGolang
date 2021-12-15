package domain

import (
	"testing"
	"time"
)

func TestNewBill(t *testing.T) {
	t.Run("Should create a new bill without scheduled date successfully", func(t *testing.T) {
		dueDate := time.Now().AddDate(0, 0, 3)

		data := Bill{
			Description: "Conta de Luz",
			AccountId:   "16sfd5465fd6s",
			Value:       26765,
			DueDate:     dueDate,
		}

		newBill, err := NewBill(data)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		scheduledDate := newBill.ScheduledDate.UTC().Truncate(24 * time.Hour)
		actualDay := time.Now().UTC().Truncate(24 * time.Hour)

		if scheduledDate != actualDay {
			t.Errorf("Expected scheduled Date is %v, got %v", actualDay, scheduledDate)
		}

		if newBill.StatusBill != Agendado {
			t.Errorf("Expected status bill is %s, got %s", Agendado, newBill.StatusBill)
		}
	})

	t.Run("Should create a new Bill with future scheduled date", func(t *testing.T) {
		dueDate := time.Now().AddDate(0, 0, 5)
		scheduledDate := time.Now().AddDate(0, 0, 6)

		billScheduled := Bill{
			Description:   "Conta de Internet",
			Value:         15000,
			AccountId:     "16sfd5465fd6s",
			DueDate:       dueDate,
			ScheduledDate: scheduledDate,
		}

		newBillScheduled, err := NewBill(billScheduled)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		dateNewBillScheduled := newBillScheduled.ScheduledDate

		if dateNewBillScheduled != scheduledDate {
			t.Errorf("Expected %v, got %v", dateNewBillScheduled, scheduledDate)
		}
	})
}
