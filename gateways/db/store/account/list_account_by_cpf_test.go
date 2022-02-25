package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var SourceTest = "account_test.json"

func TestListAccountByCpf(t *testing.T) {
	type TestCase struct {
		Name       string
		args       string
		sourceTest string
		runBefore  func(string, interface{}) error
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should list account by cpf successfull",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args:       store.AccountsTest[0].Cpf,
			want:       store.AccountsTest[0],
		},
		{
			Name: "Fail if cpf doesnt exist",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args:       "00313945153",
			want:       account.Account{},
			err:        ErrCpfNotExists,
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
				err := tt.runBefore(tt.sourceTest, store.AccountsTest)
				if err != nil {
					t.Errorf("error run before %s", err)
				}
			}

			str := NewAccountStore()
			str.Src = tt.sourceTest

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
