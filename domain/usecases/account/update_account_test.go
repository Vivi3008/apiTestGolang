package account

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
)

func TestAccountUsecase(t *testing.T) {
	t.Parallel()

	secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
	person := account.Account{
		Id:        uuid.New().String(),
		Name:      "Viviane",
		Cpf:       "55985633301",
		Secret:    secretHash,
		Balance:   260000,
		CreatedAt: time.Now(),
	}

	type args struct {
		id      string
		value   int
		payment MethodPayment
	}

	type TestCase struct {
		name       string
		args       args
		repository account.AccountRepository
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should modify balance account with credit",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{
						Id:        person.Id,
						Name:      person.Name,
						Cpf:       person.Cpf,
						Secret:    person.Secret,
						Balance:   261000,
						CreatedAt: person.CreatedAt,
					}, nil
				},
			},
			args: args{
				id:      person.Id,
				value:   1000,
				payment: Credit,
			},
			want: account.Account{
				Id:        person.Id,
				Name:      person.Name,
				Cpf:       person.Cpf,
				Secret:    person.Secret,
				Balance:   261000,
				CreatedAt: person.CreatedAt,
			},
			err: nil,
		},
		{
			name: "Should modify balance account with debit",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{
						Id:        person.Id,
						Name:      "Viviane",
						Cpf:       "55985633301",
						Secret:    secretHash,
						Balance:   259000,
						CreatedAt: person.CreatedAt,
					}, nil
				},
			},
			args: args{
				id:      person.Id,
				value:   1000,
				payment: Debit,
			},
			want: account.Account{
				Id:        person.Id,
				Name:      "Viviane",
				Cpf:       "55985633301",
				Secret:    secretHash,
				Balance:   259000,
				CreatedAt: person.CreatedAt,
			},
			err: nil,
		},
		{
			name: "Should fail if balance has insufficient limit",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, fmt.Errorf("error")
				},
			},
			args: args{
				id:      person.Id,
				value:   300000,
				payment: Debit,
			},
			want: account.Account{},
			err:  ErrInsufficientLimit,
		},
		{
			name: "Fail if value is zero",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, fmt.Errorf("error")
				},
			},
			args: args{
				id:      person.Id,
				payment: Debit,
			},
			want: account.Account{},
			err:  ErrValueEmpty,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewAccountUsecase(tt.repository)

			got, err := uc.UpdateAccountBalance(context.Background(), tt.args.id, tt.args.value, tt.args.payment)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
