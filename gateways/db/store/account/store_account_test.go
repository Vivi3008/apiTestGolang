package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestStoreAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		runBefore func() error
		args      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should write account in file",
			args: account.Account{
				Id:        uuid.NewString(),
				Name:      "Teste",
				Cpf:       "13233255666",
				Secret:    "123456",
				Balance:   400000,
				CreatedAt: time.Now(),
			},
		},
		{
			Name: "Fail if cpf exists",
			runBefore: func() error {
				return CreateAccountsInFile()
			},
			args: AccountsTest[0],
			err:  ErrCpfExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := DeleteDataTests()
				if err != nil {
					t.Errorf("error in delete data tests %s", err)
				}
			})

			if tt.runBefore != nil {
				tt.runBefore()
			}

			str := NewAccountStore()

			str.src = "account_test.json"
			err := str.StoreAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
