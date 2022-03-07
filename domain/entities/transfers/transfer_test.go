package transfers

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

var transaction = Transfer{
	AccountOriginId:      "fasf313",
	AccountDestinationId: "1fads1",
	Amount:               13321,
}

func TestCreateTransfer(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		args Transfer
		want Transfer
		err  error
	}

	testCases := []testCase{
		{
			name: "Should create a new transfer successfully",
			args: transaction,
			want: Transfer{
				AccountOriginId:      transaction.AccountOriginId,
				AccountDestinationId: transaction.AccountDestinationId,
				Amount:               transaction.Amount,
				CreatedAt:            time.Now().UTC().Truncate(24 * time.Hour),
			},
			err: nil,
		},
		{
			name: "Fail if transfer with 0 amount",
			args: Transfer{
				AccountOriginId:      "165465f65",
				AccountDestinationId: "65d6asf5d",
			},
			want: Transfer{},
			err:  ErrInvalidAmount,
		},
		{
			name: "Fail if transfer with account id empty",
			args: Transfer{
				AccountDestinationId: "65d6asf5d",
				Amount:               1500,
			},
			want: Transfer{},
			err:  ErrEmptyValues,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewTransfer(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Error expected: %s, got: %s", tt.err, err)
			}

			tt.want.CreatedAt = got.CreatedAt
			tt.want.Id = got.Id

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
