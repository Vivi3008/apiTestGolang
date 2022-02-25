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

func TestListAccountById(t *testing.T) {
	type TestCase struct {
		Name       string
		args       string
		runBefore  func(string, interface{}) error
		sourceTest string
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should list an account by id",
			args: store.AccountsTest[0].Id,
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			want:       store.AccountsTest[0],
		},
		{
			Name: "Fail if id doesnt exist",
			args: uuid.NewString(),
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
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
				tt.runBefore(tt.sourceTest, store.AccountsTest)
			}

			str := NewAccountStore()
			str.src = tt.sourceTest

			got, err := str.ListAccountById(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expeceted %s, got %s", tt.err, err)
			}

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
