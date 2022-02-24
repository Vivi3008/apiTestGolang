package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestListAccountById(t *testing.T) {
	type TestCase struct {
		Name string
		args string
		want account.Account
		err  error
	}

	testCases := []TestCase{
		{
			Name: "Should list an account by id",
			args: AccountsTest[0].Id,
			want: AccountsTest[0],
		},
		{
			Name: "Fail if id doesnt exist",
			args: uuid.NewString(),
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

			err := CreateAccountsInFile()
			if err != nil {
				t.Errorf("Error in create accounts test file %s", err)
			}

			str := NewAccountStore()
			str.src = "account_test.json"

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
