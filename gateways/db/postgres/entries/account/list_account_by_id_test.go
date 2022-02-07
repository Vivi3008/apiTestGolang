package accountdb

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListAccountById(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		runBefore func(pgx *pgxpool.Pool) error
		args      string
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should list acccount by id succesfull",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: AccountsTest[0].Id,
			want: AccountsTest[0],
		},
		{
			Name: "Fail if id doesn't exists",
			args: uuid.NewString(),
			want: account.Account{},
			err:  ErrIdNotExists,
		},
		{
			Name: "Fail if list account is empty",
			args: uuid.NewString(),
			want: account.Account{},
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			testDb, tearDown := postgres.GetTestPool()
			repo := NewRepository(testDb)
			t.Cleanup(tearDown)

			if tt.runBefore != nil {
				err := tt.runBefore(testDb)

				if err != nil {
					t.Fatalf("Error in run before %s", err)
				}
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
