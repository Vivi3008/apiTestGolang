package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
	store2 "github.com/Vivi3008/apiTestGolang/store"
	"testing"
)

func TestLogin(t *testing.T) {
	store := store2.NewAccountStore()

	t.Run("Should verify credentials and return true or false", func(t *testing.T) {
		credentials := Login{
			Cpf:      13323332555,
			Secret: "dafd33255",
		}

		person := domain.Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		acc, err2 := domain.NewAccount(person)

		if err2 != nil {
			t.Fatal("Account should have been created successfully")
		}

		err3 := store.StoreAccount(acc)

		if err3 != nil {
			t.Fatal("Account should have been stored successfully")
		}

		result, err := NewLogin(credentials)

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if result != true {
			t.Errorf("Resultado %v, esperado %v", result, true)
		}
	})
}
