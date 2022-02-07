package transfers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/accountdb"

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
			Name: "Should create a transfer successfull",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      accountdb.accountsTest[0].Id,
				AccountDestinationId: accountdb.accountsTest[1].Id,
				Amount:               5000,
				CreatedAt:            time.Now(),
			},
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
				tt.runBefore(testDb)
			}

			err := repo.SaveTransfer(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
