package transfers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSaveTransfer(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		runBefore func(pgx *pgxpool.Pool) error
		args      transfers.Transfer
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should create a transfer successful",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      accountdb.AccountsTest[0].Id,
				AccountDestinationId: accountdb.AccountsTest[1].Id,
				Amount:               5000,
				CreatedAt:            time.Now(),
			},
		},
		{
			Name: "Fail if account id is equal destiny id",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      accountdb.AccountsTest[0].Id,
				AccountDestinationId: accountdb.AccountsTest[0].Id,
				Amount:               5000,
				CreatedAt:            time.Now(),
			},
			err: ErrIdEquals,
		},
		{
			Name: "Fail if account origin id dont exists",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      uuid.NewString(),
				AccountDestinationId: accountdb.AccountsTest[0].Id,
				Amount:               400000,
				CreatedAt:            time.Now(),
			},
			err: ErrIdOriginNotExist,
		},
		{
			Name: "Fail if account destiny id don't exists",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      accountdb.AccountsTest[0].Id,
				AccountDestinationId: uuid.NewString(),
				Amount:               400000,
				CreatedAt:            time.Now(),
			},
			err: ErrIdDestinyNotExist,
		},
		{
			Name: "Fail if amount is less than zero",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      accountdb.AccountsTest[0].Id,
				AccountDestinationId: accountdb.AccountsTest[1].Id,
				Amount:               -23,
				CreatedAt:            time.Now(),
			},
			err: ErrAmountInvalid,
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
					t.Errorf("error in run before %s", err)
				}
			}

			err := repo.SaveTransfer(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
