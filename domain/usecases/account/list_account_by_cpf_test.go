package account

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestListAccountByCpf(t *testing.T) {
	t.Parallel()

	secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
	person := account.Account{
		Id:        uuid.New().String(),
		Name:      "Simon",
		Cpf:       "55985633301",
		Secret:    secretHash,
		Balance:   260000,
		CreatedAt: time.Now(),
	}

	type TestCase struct {
		Name       string
		repository account.AccountMock
		args       string
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should list account by cpf successful",
			repository: account.AccountMock{
				OnListByCpf: func(cpf string) (account.Account, error) {
					return person, nil
				},
			},
			args: person.Cpf,
			want: person,
			err:  nil,
		},
		{
			Name: "Fail if cpf doesn't exist",
			repository: account.AccountMock{
				OnListByCpf: func(cpf string) (account.Account, error) {
					return account.Account{}, ErrCpfNotExists
				},
			},
			args: person.Cpf,
			want: account.Account{},
			err:  ErrCpfNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			uc := NewAccountUsecase(tt.repository)

			got, err := uc.ListAccountByCpf(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, gto %v", tt.want, got)
			}
		})
	}
}
