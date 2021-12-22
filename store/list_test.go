package store

import (
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func TestAccountStore_ListAll(t *testing.T) {
	store := NewAccountStore()
	storeTr := NewTransferStore()
	storeBl := NewBillStore()
	layoutIso := "2006-01-02"
	dueDate, _ := time.Parse(layoutIso, "2021-12-31")

	t.Run("Should return all accounts successfully", func(t *testing.T) {
		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		person2 := account.Account{
			Name:    "Viviane",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		acc1, err := account.NewAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err = store.StoreAccount(acc1)

		if err != nil {
			t.Fatal("Account should have been stored successfully")
		}

		acc2, err2 := account.NewAccount(person2)

		if err2 != nil {
			t.Fatal("Account should have been created successfully")
		}

		err2 = store.StoreAccount(acc2)

		if err2 != nil {
			t.Fatal("Account should have been stored successfully")
		}

		accounts, err := store.ListAllAccounts()

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
		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		acc1, _ := account.NewAccount(person)

		transaction := transfers.Transfer{
			AccountOriginId:      acc1.Id,
			AccountDestinationId: "21daf3ds",
			Amount:               66541,
		}

		transaction2 := transfers.Transfer{
			AccountOriginId:      acc1.Id,
			AccountDestinationId: "21daffsda3ds",
			Amount:               67541,
		}

		tr1, err := transfers.NewTransfer(transaction)

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

		tr2, err := transfers.NewTransfer(transaction2)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		err2 := storeTr.StoreTransfer(tr2)

		if err2 != nil {
			t.Fatal("Account should have been stored successfully")
		}

		transfers, err3 := storeTr.ListTransfers(acc1.Id)

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
		bill := bills.Bill{
			AccountId:   "54545453232",
			Description: "Unimed",
			Value:       450.00,
			DueDate:     dueDate,
		}

		newBill, _ := bills.NewBill(bill)
		err := storeBl.StoreBill(newBill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		bill2 := bills.Bill{
			AccountId:   "54545453232",
			Description: "Academia",
			Value:       100,
			DueDate:     dueDate,
		}

		newBill2, _ := bills.NewBill(bill2)
		err = storeBl.StoreBill(newBill2)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		bills, err := storeBl.ListBills(bill.AccountId)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		if len(bills) != 2 {
			t.Errorf("Expected list bills 2, got %v", len(bills))
		}
	})
}
