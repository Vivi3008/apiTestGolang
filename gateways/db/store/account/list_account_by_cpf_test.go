package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func TestListAccountByCpf(t *testing.T) {
	type TestCase struct {
		Name      string
		args      string
		runBefore func() error
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should list account by cpf successfull",
			runBefore: func() error {
				return CreateAccountsInFile()
			},
			args: AccountsTest[0].Cpf,
			want: AccountsTest[0],
		},
		{
			Name: "Fail if cpf doesnt exist",
			runBefore: func() error {
				return CreateAccountsInFile()
			},
			args: "00313945153",
			want: account.Account{},
			err:  ErrCpfNotExists,
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

			got, err := str.ListAccountByCpf(context.Background(), tt.args)

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
