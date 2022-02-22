package account

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func TestListAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name string
		want []account.Account
		err  error
	}

	testCases := []TestCase{
		{
			Name: "Should list all accounts in file",
			want: []account.Account{
				{
					Id:      "475f0fa0-7eb6-4e2b-9782-c937d48d4bbb",
					Name:    "Teste",
					Cpf:     "13233255666",
					Secret:  "123456",
					Balance: 400000,
				},
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			str := NewAccountStore()

			got, err := str.ListAllAccounts()

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for k, acc := range got {
				tt.want[k].CreatedAt = acc.CreatedAt
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
