package account

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/commom"
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
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{person}, nil
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
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{person}, nil
				},
			},
			args: account.Login{
				Cpf:    person.Cpf,
				Secret: "dafd255",
			},
			want: "",
			err:  ErrInvalidPassword,
		},
		{
			name: "Fail if cpf doesn't exists",
			repository: account.AccountMock{
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{}, nil
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

			got, err := uc.NewLogin(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

	/* t.Run("Shoul sign up successfuly", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "13323332555",
			Secret:  "dafd33255",
			Balance: 250000,
		}

		credentials := account.Login{
			Cpf:    "13323332555",
			Secret: "dafd33255",
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Fatal("Account should have been created successfully")
		}

		acc, err := accounts.NewLogin(credentials)

		if err != nil {
			t.Fatal("Login error")
		}

		if acc != account.Id {
			t.Fatal("invalid Credentials")
		}
	}) */
}
