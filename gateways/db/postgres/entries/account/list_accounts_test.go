package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
)

func TestListAllAccounts(t *testing.T) {
	t.Parallel()

	testDb, tearDown := postgres.GetTestPool()
	repo := NewRepository(testDb)

	type TestCase struct {
		Name      string
		runBefore bool
		want      []account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name:      "Should list all accounts successfull",
			runBefore: true,
			want:      accountsTest,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			t.Cleanup(tearDown)

			if tt.runBefore {
				_ = createAccountTest(testDb)
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
