package domain

import "testing"

func TestStoreTransaction(t *testing.T) {

	t.Run("Should create a new transaction successfully", func(t *testing.T) {
		transaction := Transfer{
			AccountOriginId:      "fasf313",
			AccountDestinationId: "1fads1",
			Amount:               1332.1,
		}
		expected := "fasf313"
		result, err := NewTransfer(transaction)

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if result.AccountOriginId != expected {
			t.Errorf("Expected %v, result %v", expected, result.AccountOriginId)
		}

		if result.CreatedAt.IsZero() {
			t.Errorf("Expected createdAt at not to be zero")
		}

		if result.Id == "" {
			t.Errorf("Expected ID not to be empty")
		}

	})
}
