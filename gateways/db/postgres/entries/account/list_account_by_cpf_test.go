package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListAccountByCpf(t *testing.T) {
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
			Name: "Should list account by cpf successfull",
			runBefore: func(pgx *pgxpool.Pool) error {
				return createAccountTest(pgx)
			},
			args: accountsTest[0].Cpf,
			want: accountsTest[0],
			err:  nil,
		},
		{
			Name: "Fail if cpf doesn't exist",
			runBefore: func(pgx *pgxpool.Pool) error {
				return createAccountTest(pgx)
			},
			args: "1111111",
			want: account.Account{},
			err:  ErrCpfNotExists,
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
