package store

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func TestAccountStore_ListAll(t *testing.T) {
	store := NewAccountStore()
	storeTr := NewTransferStore()

	t.Run("Should return all accounts successfully", func(t *testing.T) {
		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		person2 := domain.Account{
			Name:    "Viviane",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		acc1, err := domain.NewAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err = store.StoreAccount(acc1)

		if err != nil {
			t.Fatal("Account should have been stored successfully")
		}

		acc2, err2 := domain.NewAccount(person2)

		if err2 != nil {
			t.Fatal("Account should have been created successfully")
		}

		err2 = store.StoreAccount(acc2)

		if err2 != nil {
			t.Fatal("Account should have been stored successfully")
		}

		accounts, err := store.ListAll()

		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		size := 2

		if len(accounts) != size {
			t.Errorf("expected %d; got %d", size, len(accounts))
		}

		for _, account := range accounts {
			if account.Id == acc1.Id {
				if account != acc1 {
					t.Errorf("expected %+v; got %+v", acc1, account)
				}
			}
			if account.Id == acc2.Id {
				if account != acc2 {
					t.Errorf("expected %+v; got %+v", acc2, account)
				}
			}
		}

	})

	t.Run("Should return all transfers", func(t *testing.T) {
		transaction := domain.Transfer{
			AccountOriginId:      "332f3af2",
			AccountDestinationId: "21daf3ds",
			Amount:               665.41,
		}

		transaction2 := domain.Transfer{
			AccountOriginId:      "33fdas2f3af2",
			AccountDestinationId: "21daffsda3ds",
			Amount:               675.41,
		}

		tr1, err := domain.NewTransfer(transaction)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err = storeTr.StoreTransfer(tr1)

		if err != nil {
			t.Fatal("Account should have been stored successfully")
		}

		tr2, err := domain.NewTransfer(transaction2)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err2 := storeTr.StoreTransfer(tr2)

		if err2 != nil {
			t.Fatal("Account should have been stored successfully")
		}

		transfers, err3 := storeTr.ListTransfers()

		if err3 != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		if len(transfers) != 2 {
			t.Errorf("expected %d; got %d", 2, len(transfers))
		}

		for _, transfer := range transfers {
			if transfer.Id == tr1.Id {
				if transfer != tr1 {
					t.Errorf("expected %+v; got %+v", tr1, transfer)
				}
			}
			if transfer.Id == tr2.Id {
				if transfer != tr2 {
					t.Errorf("expected %+v; got %+v", tr2, transfer)
				}
			}
		}
	})
}
