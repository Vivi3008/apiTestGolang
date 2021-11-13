package store

import (
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func TestAccountStore_ListAll(t *testing.T) {
	store := NewAccountStore()
	storeTr := NewTransferStore()
	storeBl := NewBillStore()
	layoutIso := "2006-01-02"
	dueDate, _ := time.Parse(layoutIso, "2021-12-31")

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

	t.Run("Should return all transfers from autenticated user", func(t *testing.T) {
		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		acc1, _ := domain.NewAccount(person)

		transaction := domain.Transfer{
			AccountOriginId:      acc1.Id,
			AccountDestinationId: "21daf3ds",
			Amount:               665.41,
		}

		transaction2 := domain.Transfer{
			AccountOriginId:      acc1.Id,
			AccountDestinationId: "21daffsda3ds",
			Amount:               675.41,
		}

		tr1, err := domain.NewTransfer(transaction)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err = store.StoreAccount(acc1)

		if err != nil {
			t.Fatal("Account should have been stored successfully")
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

		transfers, err3 := storeTr.ListTransfers(domain.AccountId(acc1.Id))

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

	t.Run("Should list all bills", func(t *testing.T) {
		bill := domain.Bill{
			AccountId:   "54545453232",
			Description: "Unimed",
			Value:       450.00,
			DueDate:     dueDate,
		}

		newBill, _ := domain.NewBill(bill)
		err := storeBl.StoreBill(newBill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		bill2 := domain.Bill{
			AccountId:   "54545453232",
			Description: "Academia",
			Value:       100,
			DueDate:     dueDate,
		}

		newBill2, _ := domain.NewBill(bill2)
		err = storeBl.StoreBill(newBill2)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		bills, err := storeBl.ListBills(domain.AccountId(bill.AccountId))

		if len(bills) != 2 {
			t.Errorf("Expected list bills 2, got %v", len(bills))
		}
	})
}
