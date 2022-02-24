package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestUpdateAccount(t *testing.T) {

	type args struct {
		balance int
		id      string
	}

	type TestCase struct {
		Name      string
		args      args
		runBefore func() error
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should update account successfull",
			args: args{balance: 500000, id: AccountsTest[0].Id},
			runBefore: func() error {
				return CreateAccountsInFile()
			},
			want: account.Account{
				Id:        AccountsTest[0].Id,
				Name:      AccountsTest[0].Name,
				Cpf:       AccountsTest[0].Cpf,
				Secret:    AccountsTest[0].Secret,
				Balance:   500000,
				CreatedAt: AccountsTest[0].CreatedAt,
			},
		},
		{
			Name: "Fail if id doens't exist",
			args: args{balance: 500000, id: uuid.NewString()},
			runBefore: func() error {
				return CreateAccountsInFile()
			},
			want: account.Account{},
			err:  ErrIdNotExists,
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
				err := tt.runBefore()
				if err != nil {
					t.Errorf("error run before %s", err)
				}
			}

			str := NewAccountStore()
			str.src = "account_test.json"

			got, err := str.UpdateAccount(context.Background(), tt.args.balance, tt.args.id)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
