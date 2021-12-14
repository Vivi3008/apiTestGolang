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
			Cpf:     77845100032,
			Secret:  "dafd33255",
			Balance: 250000,
		}

		person2 := domain.Account{
			Name:    "Viviane",
			Cpf:     55985633301,
			Secret:  "4f5ds4af54",
			Balance: 260000,
		}

		person3 := domain.Account{
			Name:    "Giovanna",
			Cpf:     85665232145,
			Secret:  "fadsfdsaf",
			Balance: 360000,
		}

		_, err := accounts.CreateAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		_, err = accounts.CreateAccount(person2)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		_, err = accounts.CreateAccount(person3)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
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
			Balance: 250000,
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
