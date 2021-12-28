package account

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var person = entities.Account{
	Name:    "Giovanna",
	Cpf:     "85665232145",
	Secret:  "fadsfdsaf",
	Balance: 360000,
}

func TestCreateAccount(t *testing.T) {
	t.Run("Should create an account successfull", func(t *testing.T) {
		accountMock := entities.AccountMock{
			OnCreate: func(account entities.Account) (entities.Account, error) {
				return person, nil
			},
			OnListAll: func() ([]entities.Account, error) {
				return []entities.Account{}, nil
			},
			OnStoreAccount: func(account entities.Account) error {
				return nil
			},
		}

		accUsecase := NewAccountUsecase(accountMock)

		newAccount, err := accUsecase.CreateAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		if newAccount.Id == "" {
			t.Errorf("Expected id not be nil")
		}

		if newAccount.CreatedAt.IsZero() {
			t.Errorf("Expected Created At not be nil")
		}
	})

	t.Run("Should not create an account if cpf already exists", func(t *testing.T) {
		accountMock := entities.AccountMock{
			OnCreate: func(account entities.Account) (entities.Account, error) {
				return person, nil
			},
			OnListAll: func() ([]entities.Account, error) {
				return []entities.Account{person}, nil
			},
			OnStoreAccount: func(account entities.Account) error {
				return nil
			},
		}

		accUsecase := NewAccountUsecase(accountMock)

		_, err := accUsecase.CreateAccount(person)

		if err == nil {
			t.Errorf("Expected error %s", ErrCpfExists)
		}
	})

	t.Run("Fail if name, cpf or secret is empty or missing caracters", func(t *testing.T) {
		personFail := account.Account{
			Cpf:    "00123568974",
			Secret: "a63f5ds6a5f",
		}
		accMock := account.AccountMock{
			OnListAll: func() ([]entities.Account, error) {
				return []entities.Account{person}, nil
			},
			OnStoreAccount: func(account entities.Account) error {
				return nil
			},
		}

		accUsecase := NewAccountUsecase(accMock)

		_, err := accUsecase.CreateAccount(personFail)

		if err != account.ErrInvalidValue {
			t.Errorf("Expected error %s, got %s", account.ErrInvalidValue, err)
		}

		personFail2 := account.Account{
			Name:   "fjadsiuijn",
			Cpf:    "1363565",
			Secret: "a63f5ds6a5f",
		}

		_, err = accUsecase.CreateAccount(personFail2)

		if err != account.ErrCpfCaracters {
			t.Errorf("Expected error %s, got %s", account.ErrCpfCaracters, err)
		}
	})
}
