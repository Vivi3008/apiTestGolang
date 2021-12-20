package domain

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

var person = Account{
	Name:    "Vanny",
	Cpf:     "13323332555",
	Secret:  "dafd33255",
	Balance: 2500,
}

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		args Account
		want Account
		err  error
	}

	testCases := []testCase{
		{
			name: "Should create an account succesfully",
			args: Account{
				Name:   person.Name,
				Cpf:    person.Cpf,
				Secret: person.Secret,
			},
			want: Account{
				Name:      person.Name,
				Cpf:       person.Cpf,
				Balance:   0,
				CreatedAt: time.Now(),
			},
			err: nil,
		},
		{
			name: "Fail to create account with empty name",
			args: Account{
				Cpf:    person.Cpf,
				Secret: person.Secret,
			},
			want: Account{},
			err:  ErrInvalidValue,
		},
		{
			name: "Fail to create account with empty secret",
			args: Account{
				Name: person.Name,
				Cpf:  person.Cpf,
			},
			want: Account{},
			err:  ErrInvalidValue,
		},
		{
			name: "Fail to create account with empty cpf",
			args: Account{
				Name:   person.Name,
				Secret: person.Secret,
			},
			want: Account{},
			err:  ErrInvalidValue,
		},
		{
			name: "Fail to create account with cpf less 11 caracters",
			args: Account{
				Name:   "Davy",
				Cpf:    "146565",
				Secret: "teset123",
			},
			want: Account{},
			err:  ErrCpfCaracters,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewAccount(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("got error %v expected error %v", err, tt.err)
			}

			if got.Id != "" {
				got.Id = tt.want.Id
			}

			tt.want.Secret = got.Secret
			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v expected %v", got, tt.want)
			}
		})
	}
}
