package accountdb

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListAllAccounts(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name      string
		runBefore func(pgx *pgxpool.Pool) error
		want      []account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should list all accounts successful",
			runBefore: func(pgx *pgxpool.Pool) error {
				return CreateAccountTest(pgx)
			},
			want: AccountsTest,
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
					t.Fatalf("error in run before %s", err)
				}
			}

			got, err := repo.ListAllAccounts(context.Background())

			if !errors.Is(tt.err, err) {
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
