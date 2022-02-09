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
		want      bills.Bill
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
			Name: "Should create a bills successfull",
			runBefore: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			args: billTest,
			want: billTest,
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
