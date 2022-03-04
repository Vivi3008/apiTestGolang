package bills

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestCreateBills(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		runBefore func(pgx *pgxpool.Pool) error
		args      bills.Bill
		err       error
	}

	billTest := bills.Bill{
		Id:            uuid.NewString(),
		AccountId:     accountdb.AccountsTest[0].Id,
		Description:   "Academia",
		Value:         11000,
		DueDate:       time.Now().AddDate(0, 0, 3),
		ScheduledDate: time.Now(),
	}

	testCases := []TestCase{
		{
			Name: "Should create a bills successful",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: billTest,
		},
		{
			Name: "Fail if account id doesn't exist",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: bills.Bill{
				Id:            uuid.NewString(),
				AccountId:     uuid.NewString(),
				Description:   "Fatura cartao",
				Value:         11000,
				DueDate:       time.Now().AddDate(0, 0, 3),
				ScheduledDate: time.Now(),
			},
			err: ErrAccountIdNotExist,
		},
		{
			Name: "Fail if bill id is empty",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: bills.Bill{
				AccountId:     accountdb.AccountsTest[0].Id,
				Description:   "IPTU",
				Value:         456,
				DueDate:       time.Now().AddDate(0, 0, 3),
				ScheduledDate: time.Now(),
			},
			err: ErrBillIdEmpty,
		},
		{
			Name: "Fail if value is less than zero",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: bills.Bill{
				Id:            uuid.NewString(),
				AccountId:     accountdb.AccountsTest[0].Id,
				Value:         -56,
				DueDate:       time.Now().AddDate(0, 0, 3),
				ScheduledDate: time.Now(),
			},
			err: ErrValueInvalid,
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

			err := repo.StoreBill(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
