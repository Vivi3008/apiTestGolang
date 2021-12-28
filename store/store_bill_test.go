package store

import (
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func TestStoreBill(t *testing.T) {
	store := NewBillStore()
	layoutIso := "2006-01-02"
	dueDate, _ := time.Parse(layoutIso, "2021-12-31")

	t.Run("Should store a bill successfuly", func(t *testing.T) {
		bill := bills.Bill{
			AccountId:   "54545453232",
			Description: "Unimed",
			Value:       45000,
			DueDate:     dueDate,
		}

		newBill, _ := bills.NewBill(bill)
		err := store.StoreBill(newBill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		if store.blStore[bill.Id].Id != bill.Id {
			t.Errorf("Expected %s, got %s", bill.Id, store.blStore[bill.Id].Id)
		}
	})
}
