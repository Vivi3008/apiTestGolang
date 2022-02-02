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

func TestListOneAccountById(t *testing.T) {
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
		name       string
		repository account.AccountRepository
		args       string
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should list account by id",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
			},
			want: person,
			args: person.Id,
			err:  nil,
		},
		{
			name: "Fail if account id doesnt exists",
			repository: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return account.Account{}, ErrIdNotExists
				},
			},
			args: person.Id,
			want: account.Account{},
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewAccountUsecase(tt.repository)

			got, err := uc.ListAccountById(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
