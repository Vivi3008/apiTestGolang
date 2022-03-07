package bills

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListBills(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name         string
		args         string
		runBeforeAcc func(pgx *pgxpool.Pool) error
		runBeforeBl  func(pgx *pgxpool.Pool) error
		want         []bills.Bill
		err          error
	}

	testCases := []TestCase{
		{
			Name: "Should list bills successful",
			args: accountdb.AccountsTest[0].Id,
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			runBeforeBl: func(pgx *pgxpool.Pool) error {
				return CreateBillsTest(pgx)
			},
			want: []bills.Bill{Bls[2], Bls[1], Bls[0]},
		},
		{
			Name: "Shoul list empty if account id doen's exist",
			args: uuid.NewString(),
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			runBeforeBl: func(pgx *pgxpool.Pool) error {
				return CreateBillsTest(pgx)
			},
			want: []bills.Bill{},
		},
		{
			Name: "Shoul list empty if db is empty",
			args: uuid.NewString(),
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			want: []bills.Bill{},
		},
		{
			Name: "Shoul list bills in ascendent order by due date",
			args: accountdb.AccountsTest[0].Id,
			runBeforeAcc: func(pgx *pgxpool.Pool) error {
				return accountdb.CreateAccountTest(pgx)
			},
			runBeforeBl: func(pgx *pgxpool.Pool) error {
				return CreateBillsTest(pgx)
			},
			want: []bills.Bill{Bls[2], Bls[1], Bls[0]},
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
					t.Errorf("error in run before %s", err)
				}
			}

			if tt.runBeforeBl != nil {
				err := tt.runBeforeBl(testDb)
				if err != nil {
					t.Errorf("error in run before %s", err)
				}
			}

			got, err := repo.ListBills(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for i := 0; i < len(tt.want); i++ {
				tt.want[i].DueDate = got[i].DueDate
				tt.want[i].ScheduledDate = got[i].ScheduledDate
				tt.want[i].CreatedAt = got[i].CreatedAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
