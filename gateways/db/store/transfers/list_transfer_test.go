package transfers

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

func TestListTransfers(t *testing.T) {
	type TestCase struct {
		name      string
		args      string
		runBefore func() error
		want      []transfers.Transfer
		err       error
	}

	testCases := []TestCase{
		{
			name: "Should list transfers order by date desc successfull",
			args: store.AccountsTest[0].Id,
			runBefore: func() error {
				return CreateTransfersFileTest()
			},
			want: []transfers.Transfer{store.TransfersTest[2], store.TransfersTest[1]},
		},
		{
			name: "List empty if id doens't have transfer",
			args: store.AccountsTest[2].Id,
			runBefore: func() error {
				return CreateTransfersFileTest()
			},
			want: []transfers.Transfer{},
		},
		{
			name: "List specific transfer",
			args: store.AccountsTest[1].Id,
			runBefore: func() error {
				return CreateTransfersFileTest()
			},
			want: []transfers.Transfer{store.TransfersTest[0]},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				err := DeleteDataTransfersTest()
				if err != nil {
					t.Errorf("error in delete data tests %s", err)
				}
			})

			if tt.runBefore != nil {
				tt.runBefore()
			}

			str := NewTransferStore()
			str.Src = "transfers_test.json"

			got, err := str.ListTransfer(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
