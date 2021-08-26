package store

import (
	"github.com/Vivi3008/apiTestGolang/domain"
	"testing"
)


func TestAccountStore_ListAll(t *testing.T) {
	store := NewAccountStore()

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

		err2 = store.StoreAccount(acc2);

		if err != nil {
			t.Fatal("Account should have been stored successfully")
		}

		accounts, err := store.ListAll()

		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		size := 1

		// verifica se o tamanho do map é diferente 2 se sim retorna erro, esse test vai falhar
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
}