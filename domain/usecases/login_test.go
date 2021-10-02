package usecases

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/Vivi3008/apiTestGolang/store"
)

func TestLogin(t *testing.T) {
	t.Run("Shoul sign up successfuly", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		credentials := domain.Login{
			Cpf:    13323332555,
			Secret: "dafd33255",
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		id, err := accounts.NewLogin(credentials)

		if err != nil {
			t.Fatal("Login error")
		}

		if id != account.Id {
			t.Fatal("invalid Credentials")
		}

	})
}
