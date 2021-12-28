package account

import (
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestAccountUsecase(t *testing.T) {
	t.Run("Should modify balance account with credit and debit", func(t *testing.T) {
		secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
		person := account.Account{
			Id:        uuid.New().String(),
			Name:      "Viviane",
			Cpf:       "55985633301",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		}

		accMock := account.AccountMock{
			OnListById: func(accountId string) (account.Account, error) {
				return person, nil
			},
			OnStoreAccount: func(account account.Account) error {
				return nil
			},
		}

		accUsecase := NewAccountUsecase(accMock)

		accUpdated, err := accUsecase.VerifyAccount(person.Id, 1000, Credit)

		if err != nil {
			t.Errorf("Expected nil got %s", err)
		}

		if accUpdated.Balance != 261000 {
			t.Errorf("Expected balance %v, got %v", 261000, accUpdated.Balance)
		}

		accUpdated, err = accUsecase.VerifyAccount(accUpdated.Id, 1000, Debit)

		if err != nil {
			t.Errorf("Expected nil got %s", err)
		}

		if accUpdated.Balance != 259000 {
			t.Errorf("Expected balance %v, got %v", 259000, accUpdated.Balance)
		}
	})

	t.Run("Should fail if balance has insufficient limit", func(t *testing.T) {
		secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
		person := account.Account{
			Id:        uuid.New().String(),
			Name:      "Viviane",
			Cpf:       "55985633301",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		}

		accMock := account.AccountMock{
			OnListById: func(accountId string) (account.Account, error) {
				return person, nil
			},
			OnStoreAccount: func(account account.Account) error {
				return nil
			},
		}

		accUsecase := NewAccountUsecase(accMock)

		_, err := accUsecase.VerifyAccount(person.Id, 300000, Debit)

		if err != ErrInsufficientLimit {
			t.Errorf("Expected %s, got nil", ErrInsufficientLimit)
		}
	})
}
