package usecases

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var listAccount = make([]account.Account, 0)

func MockListAccount(acc account.Account) {
	listAcc1 := map[string]account.Account{
		acc.Id: acc,
	}
	listAccount = append(listAccount, listAcc1[acc.Id])
}

func TestAccounts_ListAll(t *testing.T) {
	t.Run("Should return list of accounts succesfully", func(t *testing.T) {
		accountStore := account.AccountMock{
			OnStoreAccount: func(acc account.Account) error {
				return nil
			},
		}
		accounts := CreateNewAccount(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "77845100032",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		person2 := account.Account{
			Name:    "Viviane",
			Cpf:     "55985633301",
			Secret:  "4f5ds4af54",
			Balance: 260000,
		}

		person3 := account.Account{
			Name:    "Giovanna",
			Cpf:     "85665232145",
			Secret:  "fadsfdsaf",
			Balance: 360000,
		}

		acc1, err := accounts.CreateAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		acc2, err := accounts.CreateAccount(person2)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		acc3, err := accounts.CreateAccount(person3)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		_ = account.AccountMock{
			OnListAll: func() ([]account.Account, error) {
				MockListAccount(acc1)
				MockListAccount(acc2)
				MockListAccount(acc3)

				return listAccount, nil
			},
		}

		if len(listAccount) != 3 {
			t.Errorf("expected %d; got %d", 3, len(listAccount))
		}

	})

	/* t.Run("Should list one account by Id", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		var _ = CreateAccountStore(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		account, err := account.NewAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		account, err = accountStore.ListOne(account.Id)

		if err != nil {
			t.Fatal("expected nil; got ")
		}

		if account.Name != "Vanny" {
			t.Fatal("Account should have been listed")
		}
	}) */
}
