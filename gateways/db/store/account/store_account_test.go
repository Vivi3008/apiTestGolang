package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

func TestStoreAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name       string
		runBefore  func(string, interface{}) error
		sourceTest string
		args       account.Account
		err        error
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
			sourceTest: SourceTest,
		},
		{
			Name: "Fail if cpf exists",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args:       store.AccountsTest[0],
			err:        ErrCpfExists,
		},
		{
			Name: "Fail if account id is empty",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args: account.Account{
				Name:      "Teste sem id",
				Cpf:       "13233255666",
				Secret:    "123456",
				Balance:   400000,
				CreatedAt: time.Now(),
			},
			err: ErrEmptyID,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := store.DeleteDataFile(tt.sourceTest)
				if err != nil {
					t.Errorf("error in delete data tests %s", err)
				}
			})

			if tt.runBefore != nil {
				tt.runBefore(tt.sourceTest, store.AccountsTest)
			}

			str := NewAccountStore()

			str.Src = tt.sourceTest
			err := str.StoreAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
