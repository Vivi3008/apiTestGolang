package usecases

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/store"
)

func TestLogin(t *testing.T) {
	t.Run("Shoul sign up successfuly", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		credentials := account.Login{
			Cpf:    "13323332555",
			Secret: "dafd33255",
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		acc, err := accounts.NewLogin(credentials)

		if err != nil {
			t.Fatal("Login error")
		}

		if acc != account.Id {
			t.Fatal("invalid Credentials")
		}
	})
}
