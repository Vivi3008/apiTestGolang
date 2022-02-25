package activity

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var sourceTestBill = "../bills/bills_test.json"
var sourceTestTransfer = "../transfers/transfers_test.json"

func TestListActivities(t *testing.T) {
	type TestCase struct {
		Name              string
		args              string
		runBeforeBill     func(string, interface{}) error
		runBeforeTransfer func(string, interface{}) error
		sourceBill        string
		sourceTransfer    string
		want              []activities.AccountActivity
		err               error
	}

	wantActitivies := []activities.AccountActivity{
		{
			Type:      activities.Bill,
			Amount:    store.BillsTest[1].Value,
			CreatedAt: store.BillsTest[1].ScheduledDate,
			Details: activities.DescriptionPayment{
				Description: store.BillsTest[1].Description,
				Status:      store.BillsTest[1].StatusBill,
			},
		},
		{
			Type:      activities.Bill,
			Amount:    store.BillsTest[2].Value,
			CreatedAt: store.BillsTest[2].ScheduledDate,
			Details: activities.DescriptionPayment{
				Description: store.BillsTest[2].Description,
				Status:      store.BillsTest[2].StatusBill,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    store.TransfersTest[1].Amount,
			CreatedAt: store.TransfersTest[1].CreatedAt,
			Details: activities.DestinyAccount{
				Name:                 store.AccountsTest[2].Name,
				AccountDestinationId: store.TransfersTest[1].AccountDestinationId,
			},
		},
		{
			Type:      activities.Transfer,
			Amount:    store.TransfersTest[2].Amount,
			CreatedAt: store.TransfersTest[2].CreatedAt,
			Details: activities.DestinyAccount{
				Name:                 store.AccountsTest[1].Name,
				AccountDestinationId: store.TransfersTest[2].AccountDestinationId,
			},
		},
	}

	testCases := []TestCase{
		{
			Name: "Should list activities succesffull",
			args: store.BillsTest[0].AccountId,
			runBeforeBill: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			runBeforeTransfer: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceBill:     sourceTestBill,
			sourceTransfer: sourceTestTransfer,
			want:           wantActitivies,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := store.DeleteDataFile(tt.sourceBill)
				if err != nil {
					t.Errorf("err delete bill test %s", err)
				}
			})

			if tt.runBeforeBill != nil {
				tt.runBeforeBill(tt.sourceBill, store.BillsTest)
			}

			if tt.runBeforeTransfer != nil {
				tt.runBeforeTransfer(tt.sourceTransfer, store.TransfersTest)
			}

			str := NewAccountActivity()
			str.billStore.Src = tt.sourceBill

			got, err := str.ListActivity(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
