package activity

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

var sourceTestBill = "../bills/bills_test.json"
var sourceTestTransfer = "../transfers/transfers_test.json"
var sourceTestAccount = "../account/account_test.json"

func TestListActivities(t *testing.T) {
	type TestCase struct {
		Name              string
		args              string
		runBeforeAccount  func(string, interface{}) error
		runBeforeBill     func(string, interface{}) error
		runBeforeTransfer func(string, interface{}) error
		sourceBill        string
		sourceTransfer    string
		sourceAccount     string
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
			Type:      activities.Transfer,
			Amount:    store.TransfersTest[2].Amount,
			CreatedAt: store.TransfersTest[2].CreatedAt,
			Details: activities.DestinyAccount{
				Name:                 store.AccountsTest[1].Name,
				AccountDestinationId: store.TransfersTest[2].AccountDestinationId,
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
	}

	testCases := []TestCase{
		{
			Name: "Should list activities succesffull",
			args: store.AccountsTest[0].Id,
			runBeforeAccount: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			runBeforeBill: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			runBeforeTransfer: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceAccount:  sourceTestAccount,
			sourceBill:     sourceTestBill,
			sourceTransfer: sourceTestTransfer,
			want:           wantActitivies,
		},
		{
			Name: "List empty if account id doesnt exist",
			args: uuid.NewString(),
			runBeforeAccount: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			runBeforeBill: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			runBeforeTransfer: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceAccount:  sourceTestAccount,
			sourceBill:     sourceTestBill,
			sourceTransfer: sourceTestTransfer,
			want:           []activities.AccountActivity{},
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
				err = store.DeleteDataFile(tt.sourceAccount)
				if err != nil {
					t.Errorf("err delete account test %s", err)
				}
				err = store.DeleteDataFile(tt.sourceTransfer)
				if err != nil {
					t.Errorf("err delete transfer test %s", err)
				}
			})

			if tt.runBeforeBill != nil {
				tt.runBeforeBill(tt.sourceBill, store.BillsTest)
			}

			if tt.runBeforeTransfer != nil {
				tt.runBeforeTransfer(tt.sourceTransfer, store.TransfersTest)
			}

			if tt.runBeforeAccount != nil {
				tt.runBeforeAccount(tt.sourceAccount, store.AccountsTest)
			}

			str := NewAccountActivity()
			str.billStore.Src = tt.sourceBill
			str.transferStore.Src = tt.sourceTransfer
			str.accountStore.Src = tt.sourceAccount

			got, err := str.ListActivity(context.Background(), tt.args)

			for i := 0; i < len(got); i++ {
				tt.want[i].CreatedAt = got[i].CreatedAt
			}

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
