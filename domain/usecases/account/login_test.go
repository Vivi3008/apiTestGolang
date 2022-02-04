package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	secretHash, _ := commom.GenerateHashPassword("dafd33255")
	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Dfadfsa",
		Cpf:     "55566689545",
		Secret:  secretHash,
		Balance: 2500,
	}

	type TestCase struct {
		name       string
		repository account.AccountMock
		args       account.Login
		want       string
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should log in successfull",
			repository: account.AccountMock{
				OnListByCpf: func(cpf string) (account.Account, error) {
					return person, nil
				},
			},
			args: account.Login{
				Cpf:    person.Cpf,
				Secret: "dafd33255",
			},
			want: person.Id,
			err:  nil,
		},
		{
			name: "Fail if password is wrong",
			repository: account.AccountMock{
				OnListByCpf: func(cpf string) (account.Account, error) {
					return person, nil
				},
			},
			args: account.Login{
				Cpf:    person.Cpf,
				Secret: "wrong",
			},
			want: "",
			err:  ErrInvalidPassword,
		},
		{
			name: "Fail if cpf doesn't exists",
			repository: account.AccountMock{
				OnListByCpf: func(cpf string) (account.Account, error) {
					return account.Account{}, ErrCpfNotExists
				},
			},
			args: account.Login{
				Cpf:    person.Cpf,
				Secret: "dafd33255",
			},
			want: "",
			err:  ErrCpfNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewAccountUsecase(tt.repository)

			got, err := uc.NewLogin(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
