package usecases

import (
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/store"
)

func TestBills(t *testing.T) {
	t.Run("Should create a bill successfully", func(t *testing.T) {
		billStore := store.NewBillStore()
		bls := CreateNewBill(billStore)

		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "55566689545",
			Secret:  "dafd33255",
			Balance: 2500,
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		dueDate := time.Now().AddDate(0, 0, 2)

		bill := bills.Bill{
			AccountId:   account.Id,
			Description: "Conta internet",
			Value:       150,
			DueDate:     dueDate,
		}

		newBill, err := bills.NewBill(bill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		billOk, err := CreateNewBill(newBill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		_, err = bls.SaveBill(billOk)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		// verificando se debitou o valor na conta
		acc, err := accountStore.ListAccountById(account.Id)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		if acc.Balance != 2350 {
			t.Errorf("Expected %v, got %v", 2350, acc.Balance)
		}
	})
}
