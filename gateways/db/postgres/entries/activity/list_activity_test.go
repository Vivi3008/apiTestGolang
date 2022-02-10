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

	testCases := []TestCase{
		{
			Name: "Should list activities successfull",
			args: accountdb.AccountsTest[0].Id,
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
				t.Errorf("Expected %s, got %s", err, tt.err)
			}
			fmt.Println(got)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
