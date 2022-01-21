package account

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func TestListAllAccounts(t *testing.T) {
	t.Parallel()

	arrayAccounts := make([]account.Account, 0)
	arrayAccounts = append(arrayAccounts, account.Account{
		Name:    "Vanny",
		Cpf:     "77845100032",
		Secret:  "dafd33255",
		Balance: 250000,
	},
		account.Account{
			Name:    "Viviane",
			Cpf:     "55985633301",
			Secret:  "4f5ds4af54",
			Balance: 260000,
		},
		account.Account{
			Name:    "Giovanna",
			Cpf:     "85665232145",
			Secret:  "fadsfdsaf",
			Balance: 360000,
		},
	)

	type TestCase struct {
		name       string
		repository account.AccountRepository
		want       []account.Account
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should return list of accounts succesfully",
			repository: account.AccountMock{
				OnListAll: func() ([]account.Account, error) {
					return arrayAccounts, nil
				},
			},
			want: arrayAccounts,
			err:  nil,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewAccountUsecase(tt.repository)

			got, err := uc.ListAllAccounts()

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
