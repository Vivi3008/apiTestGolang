package usecases

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/Vivi3008/apiTestGolang/store"
)

func TestAccounts_ListAll(t *testing.T) {
	t.Run("Should return list of accounts succesfully", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		person2 := domain.Account{
			Name:    "Viviane",
			Cpf:     554545454,
			Secret:  "4f5ds4af54",
			Balance: 2.600,
		}

		person3 := domain.Account{
			Name:    "Giovanna",
			Cpf:     54656565,
			Secret:  "fadsfdsaf",
			Balance: 3.600,
		}

		_, err := accounts.CreateAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		_, err = accounts.CreateAccount(person2)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		_, err = accounts.CreateAccount(person3)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		list, err := accounts.ListAllAccounts()

		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		if len(list) != 3 {
			t.Errorf("expected %d; got %d", 3, len(list))
		}

	})

	t.Run("Should list one account by Id", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		account, err = accounts.ListAccountById(domain.AccountId(account.Id))

		if err != nil {
			t.Fatal("expected nil; got ")
		}

		if account.Name != "Vanny" {
			t.Fatal("Account should have been listed")
		}
	})
}
