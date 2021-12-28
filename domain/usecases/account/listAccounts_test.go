package account

import (
	"errors"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestListAccounts(t *testing.T) {
	t.Run("Should return list of accounts succesfully", func(t *testing.T) {
		arrayAccounts := make([]account.Account, 0)
		arrayAccounts = append(arrayAccounts, account.Account{
			Name:    "Vanny",
			Cpf:     "77845100032",
			Secret:  "dafd33255",
			Balance: 250000,
		},
			account.Account{
				Name:    "Viviane",
				Cpf:     "55985633301",
				Secret:  "4f5ds4af54",
				Balance: 260000,
			},
			account.Account{
				Name:    "Giovanna",
				Cpf:     "85665232145",
				Secret:  "fadsfdsaf",
				Balance: 360000,
			},
		)

		accountStore := account.AccountMock{
			OnListAll: func() ([]account.Account, error) {
				return arrayAccounts, nil
			},
		}

		accountUsecase := NewAccountUsecase(accountStore)

		got, err := accountUsecase.ListAllAccounts()

		if err != nil {
			t.Errorf("expected nil got %s", err)
		}

		if len(got) != len(arrayAccounts) {
			t.Errorf("Expected %v, got %v", len(arrayAccounts), len(got))
		}
	})

	t.Run("Should return error list of accounts", func(t *testing.T) {
		accountStore := account.AccountMock{
			OnListAll: func() ([]account.Account, error) {
				return nil, errors.New("error to list")
			},
		}

		accountUsecase := NewAccountUsecase(accountStore)

		_, err := accountUsecase.ListAllAccounts()

		if err == nil {
			t.Errorf("expected nil got %s", err)
		}
	})

	t.Run("Should list account by id", func(t *testing.T) {
		person := account.Account{
			Id:      uuid.New().String(),
			Name:    "Viviane",
			Cpf:     "77845100032",
			Secret:  "dafd33255",
			Balance: 250000,
		}
		accStore := account.AccountMock{
			OnListById: func(accountId string) (account.Account, error) {
				return person, nil
			},
		}

		accountUsecase := NewAccountUsecase(accStore)

		acount, err := accountUsecase.ListAccountById(person.Id)

		if err != nil {
			t.Errorf("expected nil got %s", err)
		}

		if acount.Name != "Viviane" {
			t.Errorf("Expected %s, got %s", "Viviane", acount.Name)
		}
	})

	t.Run("Fail to list account by Id", func(t *testing.T) {
		accMock := account.AccountMock{
			OnListById: func(accountId string) (account.Account, error) {
				return account.Account{}, errors.New("This id doesn't exists")
			},
		}

		accUsecase := NewAccountUsecase(accMock)

		_, err := accUsecase.ListAccountById(uuid.NewString())

		if err == nil {
			t.Errorf("Expected error be %s", errors.New("This id doesn't exists"))
		}
	})
}
