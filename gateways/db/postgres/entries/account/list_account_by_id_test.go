package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/google/uuid"
)

func TestListAccountById(t *testing.T) {
	t.Parallel()

	testRepo, _ := postgres.GetTestPool()
	repo := NewRepository(testRepo)

	type TestCase struct {
		Name       string
		runBefore  bool
		cleanTable bool
		args       string
		want       account.Account
		err        error
	}

	testCases := []TestCase{
		{
			Name:      "Should list acccount by id succesfull",
			runBefore: true,
			args:      accountsTest[0].Id,
			want:      accountsTest[0],
		},
		{
			Name:      "Fail if id doesn't exists",
			runBefore: false,
			args:      uuid.NewString(),
			want:      account.Account{},
			err:       ErrIdNotExists,
		},
		{
			Name:       "Fail if list account is empty",
			runBefore:  false,
			cleanTable: true,
			args:       uuid.NewString(),
			want:       account.Account{},
			err:        ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			if tt.cleanTable {
				defer cleanAccountsTable(testRepo)
			}

			if tt.runBefore {
				_ = createAccountTest(testRepo)
			}

			got, err := repo.ListAccountById(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
