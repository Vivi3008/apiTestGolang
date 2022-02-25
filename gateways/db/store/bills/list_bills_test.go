package bills

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

var SourceTest = "bills_test.json"

func TestListBills(t *testing.T) {
	type TestCase struct {
		Name       string
		args       string
		runBefore  func(string, interface{}) error
		sourceTest string
		want       []bills.Bill
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should list a bill successfull",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args:       store.AccountsTest[0].Id,
			want:       []bills.Bill{store.BillsTest[2], store.BillsTest[1]},
		},
		{
			Name: "List empty bills if account id doesnt have bill",
			runBefore: func(s string, i interface{}) error {
				return store.CreateDataFile(s, i)
			},
			sourceTest: SourceTest,
			args:       uuid.NewString(),
			want:       []bills.Bill{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := store.DeleteDataFile(tt.sourceTest)
				if err != nil {
					t.Errorf("err delete bill test %s", err)
				}
			})

			if tt.runBefore != nil {
				tt.runBefore(tt.sourceTest, store.BillsTest)
			}

			str := NewBillStore()
			str.Src = tt.sourceTest

			got, err := str.ListBills(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			for i := 0; i < len(got); i++ {
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
