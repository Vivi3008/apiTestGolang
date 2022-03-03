package bills

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

func TestStoreBill(t *testing.T) {
	type TestCase struct {
		Name       string
		sourceTest string
		args       bills.Bill
		err        error
	}

	testCases := []TestCase{
		{
			Name: "Should store a bill in file succesffull",
			args: bills.Bill{
				Id:            uuid.NewString(),
				AccountId:     uuid.NewString(),
				Value:         6000,
				DueDate:       time.Now(),
				ScheduledDate: time.Now(),
				CreatedAt:     time.Now(),
			},
			sourceTest: SourceTest,
		},
		{
			Name: "Fail if id bill is empty",
			args: bills.Bill{
				AccountId:     uuid.NewString(),
				Value:         6000,
				DueDate:       time.Now(),
				ScheduledDate: time.Now(),
				CreatedAt:     time.Now(),
			},
			sourceTest: SourceTest,
			err:        ErrEmptyID,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Cleanup(func() {
				err := store.DeleteDataFile(tt.sourceTest)
				if err != nil {
					t.Errorf("err delelte bill test %s", err)
				}
			})

			str := NewBillStore()
			str.Src = tt.sourceTest

			err := str.StoreBill(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
