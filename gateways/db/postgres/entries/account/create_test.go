package account

import (
	"context"
	"errors"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/google/uuid"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	testPool, tearDown := postgres.GetTestPool()
	repository := NewRepository(testPool)

	type TestCase struct {
		Name string
		args account.Account
		err  error
	}

	person := account.Account{
		Name:   "Teste",
		Cpf:    "13256589412",
		Secret: "fasfdsa",
	}

	acc, _ := account.NewAccount(person)

	testCases := []TestCase{
		{
			Name: "Should create account successfull",
			args: acc,
			err:  nil,
		},
		{
			Name: "Fail to create account with same cpf",
			args: acc,
			err:  ErrCpfExists,
		},
		{
			Name: "Fail to create account with same id",
			args: account.Account{
				Id:     acc.Id,
				Name:   "teste3",
				Cpf:    "654656",
				Secret: "16656",
			},
			err: ErrCpfExists,
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

			t.Cleanup(tearDown)

			err := repository.StoreAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}

}
