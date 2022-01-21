package account

import (
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

	listAccounts := []account.Account{
		{
			Id:        "fads1fdsa3",
			Name:      "David",
			Cpf:       "556565656555",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
		{
			Id:        "5df4s5df45",
			Name:      "Vale",
			Cpf:       "656565656565",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
		{
			Id:        "a6fd56sad5f3",
			Name:      "Biscui",
			Cpf:       "21545454545",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
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
				OnListAll: func() ([]account.Account, error) {
					return listAccounts, nil
				},
			},
			want: listAccounts[0],
			args: listAccounts[0].Id,
			err:  nil,
		},
		{
			name: "Fail if account id doesnt exists in empty list account",
			repository: account.AccountMock{
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{}, nil
				},
			},
			args: person.Id,
			want: account.Account{},
			err:  ErrListAccountEmpty,
		},
		{
			name: "Fail if account id doesnt exists in list not empty",
			repository: account.AccountMock{
				OnListAll: func() ([]account.Account, error) {
					return listAccounts, nil
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

			got, err := uc.ListAccountById(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
