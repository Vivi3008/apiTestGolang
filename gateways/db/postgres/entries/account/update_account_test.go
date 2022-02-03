package account

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
)

func TestUpdateAccount(t *testing.T) {
	t.Parallel()

	testeDb, tearDown := postgres.GetTestPool()
	repo := NewRepository(testeDb)

	type args struct {
		balance int
		id      string
	}

	type TestCase struct {
		Name      string
		runBefore bool
		args      args
		want      account.Account
		err       error
	}

	testCases := []TestCase{
		{
			Name:      "Should update account successfull",
			runBefore: true,
			args:      args{1000, accountsTest[0].Id},
			want:      accountsTest[0],
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			t.Cleanup(tearDown)

			if tt.runBefore {
				_ = createAccountTest(testeDb)
			}

			list, _ := repo.ListAllAccounts(context.Background())

			fmt.Println(list)
			got, err := repo.UpdateAccount(context.Background(), tt.args.balance, tt.args.id)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
