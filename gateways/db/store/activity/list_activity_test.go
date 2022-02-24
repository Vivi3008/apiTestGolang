package activity

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
)

func TestListActivities(t *testing.T) {
	type TestCase struct {
		Name          string
		args          string
		runBeforeBill func() error
		want          []activities.AccountActivity
		err           error
	}

	testCases := []TestCase{
		{
			Name: "Should list activities succesffull",
			args: BillsTest[0].AccountId,
			runBeforeBill: func() error {
				return CreateBillsActivities()
			},
			want: []activities.AccountActivity{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := DeleteBillsActivities()
				if err != nil {
					t.Errorf("err delelte bill test %s", err)
				}
			})

			if tt.runBeforeBill != nil {
				tt.runBeforeBill()
			}

			str := NewAccountAcitivity()
			str.billStore.Src = "../bills/bills_test.json"

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
