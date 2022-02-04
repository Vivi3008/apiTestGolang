package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestUpdateAccount(t *testing.T) {
	t.Parallel()

	type args struct {
		amount int
		id     string
	}

	type TestCase struct {
		Name      string
		runBefore func(pgx *pgxpool.Pool) error
		args      args
		want      int
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should update account successfull",
			runBefore: func(pgx *pgxpool.Pool) error {
				return createAccountTest(pgx)
			},
			args: args{1000, accountsTest[0].Id},
			want: 1000,
		},
		{
			Name: "Fail if amount is invalid",
			runBefore: func(pgx *pgxpool.Pool) error {
				return createAccountTest(pgx)
			},
			args: args{-50, accountsTest[0].Id},
			want: 0,
			err:  ErrBalanceInvalid,
		},
		{
			Name: "Fail update balace if id doesn't exists",
			runBefore: func(pgx *pgxpool.Pool) error {
				return createAccountTest(pgx)
			},
			args: args{2000, uuid.NewString()},
			want: 0,
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			testeDb, tearDown := postgres.GetTestPool()
			repo := NewRepository(testeDb)
			t.Cleanup(tearDown)

			if tt.runBefore != nil {
				err := tt.runBefore(testeDb)

				if err != nil {
					t.Fatalf("error in run before %s", err)
				}
			}

			got, err := repo.UpdateAccount(context.Background(), tt.args.amount, tt.args.id)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got.Balance, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
