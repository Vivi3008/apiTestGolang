package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var person = account.Account{
	Name:    "Giovanna",
	Cpf:     "85665232145",
	Secret:  "fadsfdsaf",
	Balance: 360000,
}

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		repository account.AccountRepository
		args       account.Account
		want       account.Account
		err        error
	}

	testCases := []testCase{
		{
			name: "Should create an account successfull",
			repository: account.AccountMock{
				OnStoreAccount: func(account account.Account) error {
					return nil
				},
			},
			args: person,
			want: account.Account{
				Name:    person.Name,
				Cpf:     person.Cpf,
				Balance: person.Balance,
			},
			err: nil,
		},
		{
			name: "Should not create an account if cpf already exists",
			repository: account.AccountMock{
				OnStoreAccount: func(account account.Account) error {
					return ErrCpfExists
				},
			},
			args: person,
			want: account.Account{},
			err:  ErrCpfExists,
		},
		{
			name: "Fail if name, cpf or secret is empty or missing caracters",
			repository: account.AccountMock{
				OnStoreAccount: func(account account.Account) error {
					return nil
				},
			},
			args: account.Account{
				Cpf:    "00123568974",
				Secret: "a63f5ds6a5f",
			},
			want: account.Account{},
			err:  account.ErrInvalidValue,
		},
		{
			name: "Fail if cpf has less than 11 caracters",
			repository: account.AccountMock{
				OnStoreAccount: func(account account.Account) error {
					return nil
				},
			},
			args: account.Account{
				Name:   "fjadsiuijn",
				Cpf:    "1363565",
				Secret: "a63f5ds6a5f",
			},
			want: account.Account{},
			err:  account.ErrCpfCaracters,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := NewAccountUsecase(tt.repository)

			got, err := uc.CreateAccount(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.Id = got.Id
			tt.want.Secret = got.Secret
			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}

		})
	}
}
