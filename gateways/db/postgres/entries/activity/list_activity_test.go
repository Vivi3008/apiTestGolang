package activity

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	billdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
	transferdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListActitivies(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		args      string
		runBefore func(pgx *pgxpool.Pool) error
		want      []activities.AccountActivity
		err       error
	}

	bl := billdb.Bls
	tr := transferdb.TransfersTest

	wantActitivies := []activities.AccountActivity{
		{
			Type:      activities.Bill,
			Amount:    bl[2].Value,
			CreatedAt: bl[2].ScheduledDate,
			Details: DescriptionPayment{
				Description: bl[2].Description,
				Status:      bl[2].StatusBill,
			},
		},
		{
			Type:      activities.Bill,
			Amount:    bl[1].Value,
			CreatedAt: bl[1].ScheduledDate,
			Details: DescriptionPayment{
				Description: bl[1].Description,
				Status:      bl[1].StatusBill,
			},
		},
		{
			Type:      activities.Bill,
			Amount:    bl[0].Value,
			CreatedAt: bl[0].ScheduledDate,
			Details: DescriptionPayment{
				Description: bl[0].Description,
				Status:      bl[0].StatusBill,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    tr[1].Amount,
			CreatedAt: tr[1].CreatedAt,
			Details: DestinyAccount{
				Name:                 accountdb.AccountsTest[2].Name,
				AccountDestinationId: tr[1].AccountDestinationId,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    tr[2].Amount,
			CreatedAt: tr[2].CreatedAt,
			Details: DestinyAccount{
				Name:                 accountdb.AccountsTest[1].Name,
				AccountDestinationId: tr[2].AccountDestinationId,
			},
		},
	}

	testCases := []TestCase{
		{
			Name: "Should list activities successfull",
			args: accountdb.AccountsTest[0].Id,
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateDbTest(pgx)
			},
			want: wantActitivies,
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
				t.Errorf("Expected %s, got %s", err, tt.err)
			}

			for i := 0; i < len(got); i++ {
				tt.want[i].CreatedAt = got[i].CreatedAt
				fmt.Printf("Expected %v\n", got[i])
				fmt.Printf("Want %v\n", tt.want[i])
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}