package account

import (
	"context"
	"errors"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	testPool, tearDown := postgres.GetTestPool()
	repository := NewRepository(testPool)

	t.Cleanup(tearDown)

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
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			err := repository.StoreAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}

}
