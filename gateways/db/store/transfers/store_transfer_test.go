package transfers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/google/uuid"
)

func TestStoreTransfer(t *testing.T) {
	type TestCase struct {
		name string
		args transfers.Transfer
		err  error
	}

	testCases := []TestCase{
		{
			name: "Should store transfer in file succesfull",
			args: transfers.Transfer{
				Id:                   uuid.NewString(),
				AccountOriginId:      uuid.NewString(),
				AccountDestinationId: uuid.NewString(),
				Amount:               150000,
				CreatedAt:            time.Now(),
			},
		},
		{
			name: "Fail if empty transfer id",
			args: transfers.Transfer{
				AccountOriginId:      uuid.NewString(),
				AccountDestinationId: uuid.NewString(),
				Amount:               150000,
				CreatedAt:            time.Now(),
			},
			err: ErrEmptyID,
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

			str := NewTransferStore()
			str.Src = "transfers_test.json"

			err := str.SaveTransfer(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
