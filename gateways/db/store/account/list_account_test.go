package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func TestListAccount(t *testing.T) {
	type TestCase struct {
		Name   string
		want   []account.Account
		source string
		err    error
	}

	var jsonWrong = "wrong.json"

	testCases := []TestCase{
		{
			Name: "Should list all accounts in file",
			want: []account.Account{
				{
					Name:    "Teste 1",
					Cpf:     "77845100032",
					Secret:  "dafd33255",
					Balance: 250000,
				},
				{
					Name:    "Teste 2",
					Cpf:     "55985633301",
					Secret:  "4f5ds4af54",
					Balance: 260000,
				},
				{
					Name:    "Teste 3",
					Cpf:     "85665232145",
					Secret:  "fadsfdsaf",
					Balance: 360000,
				},
			},
			source: "account_test.json",
		},
		{
			Name:   "Fail if account file source is wrong",
			want:   []account.Account{},
			source: jsonWrong,
			err:    ErrOpenFile,
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
				t.Errorf("error in create accounts file test: %s", err)
			}

			str := NewAccountStore()
			str.src = tt.source

			got, err := str.ListAllAccounts(context.Background())

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for k, acc := range got {
				tt.want[k].Id = acc.Id
				tt.want[k].CreatedAt = acc.CreatedAt
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
