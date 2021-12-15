package store

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func TestStoreTransfer(t *testing.T) {
	store := NewTransferStore()

	t.Run("Should store any transfer", func(t *testing.T) {
		transaction := domain.Transfer{
			AccountOriginId:      "332f3af2",
			AccountDestinationId: "21daf3ds",
			Amount:               66541,
		}

		transfer, _ := domain.NewTransfer(transaction) //cria a transferencia
		err := store.StoreTransfer(transfer)           //guarda a transfer num map

		if err != nil {
			t.Errorf("expected nil; got '%v'", err)
		}

		if store.tranStore[transfer.Id].AccountOriginId != transaction.AccountOriginId {
			t.Errorf(
				"Expected %s, got %s", transaction.AccountOriginId, store.tranStore[transfer.AccountOriginId].AccountOriginId)
		}

	})

	t.Run("Should return error if transfer id is empty", func(t *testing.T) {
		transaction := domain.Transfer{
			AccountOriginId:      "332f3af2",
			AccountDestinationId: "21daf3ds",
			Amount:               66541,
		}

		err := store.StoreTransfer(transaction)

		if err != ErrEmptyID {
			t.Errorf("expected %s, got %s", ErrEmptyID, err.Error())
		}
	})

	t.Run("Should not list transfers if account origin Id doesnt exists", func(t *testing.T) {
		transaction := domain.Transfer{
			AccountOriginId:      "f63cb25b-786c-4ff2-9a67-22a065d307d3",
			AccountDestinationId: "21daf3ds",
			Amount:               66541,
		}

		_, err := store.ListTransfers(transaction.AccountOriginId)

		if err == nil {
			t.Errorf("Expected error id doesnt exists, got %s", err)
		}
	})
}
