package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

func TestUpdateAccount(t *testing.T) {

	type args struct {
		balance int
		id      string
	}

	type TestCase struct {
		Name       string
		args       args
		runBefore  func(string) error
		sourceTest string
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should update account successfull",
			args: args{balance: 500000, id: store.AccountsTest[0].Id},
			runBefore: func(s string) error {
				return store.CreateDataFile(s)
			},
			sourceTest: SourceTest,
			want: account.Account{
				Id:        store.AccountsTest[0].Id,
				Name:      store.AccountsTest[0].Name,
				Cpf:       store.AccountsTest[0].Cpf,
				Secret:    store.AccountsTest[0].Secret,
				Balance:   500000,
				CreatedAt: store.AccountsTest[0].CreatedAt,
			},
		},
		{
			Name: "Fail if id doens't exist",
			args: args{balance: 500000, id: uuid.NewString()},
			runBefore: func(s string) error {
				return store.CreateDataFile(s)
			},
			sourceTest: SourceTest,
			want:       account.Account{},
			err:        ErrIdNotExists,
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
				err := tt.runBefore(tt.sourceTest)
				if err != nil {
					t.Errorf("error run before %s", err)
				}
			}

			str := NewAccountStore()
			str.src = tt.sourceTest

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
