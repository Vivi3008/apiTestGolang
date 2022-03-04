package accountdb

import (
	"context"
	"errors"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		args      account.Account
		runBefore func(pgx *pgxpool.Pool) error
		err       error
	}

	person := account.Account{
		Name:   "Teste",
		Cpf:    "113219412",
		Secret: "fasfdsa",
	}

	acc, _ := account.NewAccount(person)

	testCases := []TestCase{
		{
			Name: "Should create account successful",
			args: acc,
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			err: nil,
		},
		{
			Name: "Fail to create account with same cpf",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: account.Account{
				Id:     uuid.NewString(),
				Name:   "teste3",
				Cpf:    AccountsTest[0].Cpf,
				Secret: "16656",
			},
			err: ErrCpfExists,
		},
		{
			Name: "Fail to create account with same id",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: account.Account{
				Id:     AccountsTest[0].Id,
				Name:   "teste3",
				Cpf:    "654656",
				Secret: "16656",
			},
			err: ErrIdExists,
		},
		{
			Name: "Fail to create account with negative balance",
			args: account.Account{
				Id:      uuid.NewString(),
				Name:    "teste3",
				Cpf:     "654656",
				Balance: -56,
				Secret:  "16656",
			},
			err: ErrBalanceInvalid,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			testPool, tearDown := postgres.GetTestPool()
			repository := NewRepository(testPool)
			t.Cleanup(tearDown)

			if tt.runBefore != nil {
				tt.runBefore(testPool)
			}

			err := repository.StoreAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}

}
