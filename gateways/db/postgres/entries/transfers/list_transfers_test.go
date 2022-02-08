package transfers

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListTransfers(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name         string
		runBeforeAcc func(pgx *pgxpool.Pool) error
		runBeforeTr  func(pgx *pgxpool.Pool) error
		args         string
		want         []transfers.Transfer
		err          error
	}

	testCases := []TestCase{
		{
			Name: "Should list all transfers",
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			runBeforeTr: func(pgx *pgxpool.Pool) error {
				return CreateTransfersTest(pgx)
			},
			args: accountdb.AccountsTest[1].Id,
			want: []transfers.Transfer{TransfersTest[0], TransfersTest[1]},
		},
		{
			Name: "List transfer empty if account origin id doesn't exist",
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			runBeforeTr: func(pgx *pgxpool.Pool) error {
				return CreateTransfersTest(pgx)
			},
			args: uuid.NewString(),
			want: []transfers.Transfer{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			testDb, tearDown := postgres.GetTestPool()
			repo := NewRepository(testDb)
			t.Cleanup(tearDown)

			if tt.runBeforeAcc != nil {
				err := tt.runBeforeAcc(testDb)
				if err != nil {
					t.Fatalf("Err in runBeforeAcc: %s", err)
				}
			}

			if tt.runBeforeTr != nil {
				err := tt.runBeforeTr(testDb)
				if err != nil {
					t.Fatalf("Err in runBeforeTr: %s", err)
				}
			}

			got, err := repo.ListTransfer(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for i := 0; i < len(tt.want); i++ {
				tt.want[i].CreatedAt = got[i].CreatedAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
