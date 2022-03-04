package activity

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	billdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
	transferdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListActitivies(t *testing.T) {
	t.Parallel()

	bl := billdb.Bls
	tr := transferdb.TransfersTest

	wantActitivies := []activities.AccountActivity{
		{
			Type:      activities.Bill,
			Amount:    bl[0].Value,
			CreatedAt: bl[0].ScheduledDate,
			Details: activities.DescriptionPayment{
				Description: bl[0].Description,
				Status:      bl[0].StatusBill,
			},
		},
		{
			Type:      activities.Bill,
			Amount:    bl[1].Value,
			CreatedAt: bl[1].ScheduledDate,
			Details: activities.DescriptionPayment{
				Description: bl[1].Description,
				Status:      bl[1].StatusBill,
			},
		},
		{
			Type:      activities.Bill,
			Amount:    bl[2].Value,
			CreatedAt: bl[2].ScheduledDate,
			Details: activities.DescriptionPayment{
				Description: bl[2].Description,
				Status:      bl[2].StatusBill,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    tr[1].Amount,
			CreatedAt: tr[1].CreatedAt,
			Details: activities.DestinyAccount{
				Name:                 accountdb.AccountsTest[2].Name,
				AccountDestinationId: tr[1].AccountDestinationId,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    tr[2].Amount,
			CreatedAt: tr[2].CreatedAt,
			Details: activities.DestinyAccount{
				Name:                 accountdb.AccountsTest[1].Name,
				AccountDestinationId: tr[2].AccountDestinationId,
			},
		},
	}

	type TestCase struct {
		Name      string
		args      string
		runBefore func(pgx *pgxpool.Pool) error
		want      []activities.AccountActivity
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should list activities successful in order by created_at",
			args: accountdb.AccountsTest[0].Id,
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateDbTest(pgx)
			},
			want: wantActitivies,
		},
		{
			Name: "List empty if account id doesnt exist",
			args: uuid.NewString(),
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateDbTest(pgx)
			},
			want: []activities.AccountActivity{},
		},
		{
			Name: "Empty list for accounts that have no outputs",
			args: accountdb.AccountsTest[2].Id,
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateDbTest(pgx)
			},
			want: []activities.AccountActivity{},
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

			got, err := repo.ListActivity(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for i := 0; i < len(got); i++ {
				tt.want[i].CreatedAt = got[i].CreatedAt
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
