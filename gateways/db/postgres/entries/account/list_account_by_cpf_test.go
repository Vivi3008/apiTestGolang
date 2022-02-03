package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
)

func TestListAccountByCpf(t *testing.T) {
	t.Parallel()

	testRepo, _ := postgres.GetTestPool()
	repo := NewRepository(testRepo)

	type TestCase struct {
		Name      string
		runBefore bool
		args      string
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name:      "Should list account by cpf successfull",
			runBefore: true,
			args:      accountsTest[0].Cpf,
			want:      accountsTest[0],
			err:       nil,
		},
		{
			Name:      "Fail if cpf doesn't exist",
			runBefore: false,
			args:      "1111111",
			want:      account.Account{},
			err:       ErrCpfNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			if tt.runBefore {
				_ = createAccountTest(testRepo)
			}

			got, err := repo.ListAccountByCpf(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
