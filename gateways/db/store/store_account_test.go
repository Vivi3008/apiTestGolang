package store

import (
	"context"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func TestStoreAccount(t *testing.T) {
	store := NewAccountStore()

	t.Run("Shoud store an account successfully", func(t *testing.T) {
		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 25000,
		}
		account, _ := account.NewAccount(person)                 //cria a conta
		err := store.StoreAccount(context.Background(), account) // guarda a conta num map

		if err != nil {
			t.Errorf("expected nil; got '%v'", err)
		}

		if store.accStore[account.Id].Name != person.Name {
			t.Errorf("Expected %s, got %s", person.Name, store.accStore[person.Id].Name)
		}
	})

	t.Run("Should return error if account id is empty", func(t *testing.T) {
		acc := account.Account{
			Id:     "",
			Name:   "Viviane",
			Cpf:    "00314522352",
			Secret: "dadfdasf",
		}
		err := store.StoreAccount(context.Background(), acc)
		if err != ErrEmptyID {
			t.Errorf("expected %s, got %s", ErrEmptyID, err.Error())
		}
	})

}
