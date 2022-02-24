package bills

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/google/uuid"
)

func TestListBills(t *testing.T) {
	type TestCase struct {
		Name      string
		args      string
		runBefore func() error
		want      []bills.Bill
		err       error
	}

	testCases := []TestCase{
		{
			Name: "Should list a bill successfull",
			runBefore: func() error {
				return CreateBillsTest()
			},
			args: BillsTest[0].AccountId,
			want: BillsTest,
		},
		{
			Name: "List empty bills if account id doesnt have bill",
			runBefore: func() error {
				return CreateBillsTest()
			},
			args: uuid.NewString(),
			want: []bills.Bill{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := DeleteDataBillTests()
				if err != nil {
					t.Errorf("err delelte bill test %s", err)
				}
			})

			if tt.runBefore != nil {
				tt.runBefore()
			}

			str := NewBillStore()
			str.Src = "bills_test.json"

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
