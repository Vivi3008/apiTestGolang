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
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should update account successful",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: args{1000, AccountsTest[0].Id},
			want: account.Account{
				Id:        AccountsTest[0].Id,
				Name:      AccountsTest[0].Name,
				Cpf:       AccountsTest[0].Cpf,
				Balance:   1000,
				CreatedAt: AccountsTest[0].CreatedAt,
			},
		},
		{
			Name: "Fail if amount is invalid",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: args{-50, AccountsTest[0].Id},
			want: account.Account{},
			err:  ErrBalanceInvalid,
		},
		{
			Name: "Fail update balace if id doesn't exists",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			args: args{2000, uuid.NewString()},
			want: account.Account{},
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

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
