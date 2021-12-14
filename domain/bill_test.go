package domain

import (
	"testing"
)

func TestNewBill(t *testing.T) {
	t.Run("Should create a new bill without scheduled date successfully", func(t *testing.T) {
		layoutIso := "2006-01-02"
		dueDate := actualDate.AddDate(0, 0, 3)

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

		scheduledDate := newBill.ScheduledDate.Format(layoutIso)

		if scheduledDate != actualDate.Format(layoutIso) {
			t.Errorf("Expected scheduled Date is actual day, got %v", newBill.ScheduledDate)
		}

		if newBill.StatusBill != Agendado {
			t.Errorf("Expected status bill is %s, got %s", Agendado, newBill.StatusBill)
		}
	})

	t.Run("Should create a new Bill with future scheduled date", func(t *testing.T) {
		layoutIso := "2006-01-02"
		dueDate := actualDate.AddDate(0, 0, 5)
		scheduledDate := actualDate.AddDate(0, 0, 6)

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

		dateNewBillScheduled := newBillScheduled.ScheduledDate.Format(layoutIso)

		if dateNewBillScheduled != scheduledDate.Format(layoutIso) {
			t.Errorf("Expected %v, got %v", dateNewBillScheduled, scheduledDate.Format(layoutIso))
		}
	})
}
