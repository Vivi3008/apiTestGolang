package domain

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewBill(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		args Bill
		want Bill
		err  error
	}

	testCases := []testCase{
		{
			name: "Should create a new bill without scheduled date successfully",
			args: Bill{
				Description: "Conta de Luz",
				AccountId:   "16sfd5465fd6s",
				Value:       26765,
				DueDate:     time.Now().AddDate(0, 0, 3),
			},
			want: Bill{
				Description:   "Conta de Luz",
				AccountId:     "16sfd5465fd6s",
				Value:         26765,
				DueDate:       time.Now().AddDate(0, 0, 3).UTC().Truncate(24 * time.Hour),
				ScheduledDate: time.Now().UTC().Truncate(24 * time.Hour),
				StatusBill:    Agendado,
			},
			err: nil,
		},
		{
			name: "Should create a new Bill with future scheduled date",
			args: Bill{
				Description:   "Academia",
				Value:         15000,
				AccountId:     "16sfd5465fd6s",
				DueDate:       time.Now().AddDate(0, 0, 5),
				ScheduledDate: time.Now().AddDate(0, 0, 6),
			},
			want: Bill{
				AccountId:     "16sfd5465fd6s",
				Description:   "Academia",
				Value:         15000,
				DueDate:       time.Now().AddDate(0, 0, 5).UTC().Truncate(24 * time.Hour),
				ScheduledDate: time.Now().AddDate(0, 0, 6).UTC().Truncate(24 * time.Hour),
				StatusBill:    Agendado,
			},
			err: nil,
		},
		{
			name: "Fail if scheduled date is before today",
			args: Bill{
				Description:   "TIM",
				Value:         5900,
				AccountId:     "16sfd5465fd6s",
				DueDate:       time.Now().AddDate(0, 0, 5),
				ScheduledDate: time.Now().AddDate(0, 0, -1),
			},
			want: Bill{},
			err:  ErrDateInvalid,
		},
		{
			name: "Fail if due date is empty",
			args: Bill{
				Description:   "TIM",
				Value:         5900,
				AccountId:     "16sfd5465fd6s",
				ScheduledDate: time.Now().AddDate(0, 0, -1),
			},
			want: Bill{},
			err:  ErrEmpty,
		},
		{
			name: "Fail if description is empty",
			args: Bill{
				Value:         5900,
				DueDate:       time.Now().AddDate(0, 0, 5),
				AccountId:     "16sfd5465fd6s",
				ScheduledDate: time.Now().AddDate(0, 0, -1),
			},
			want: Bill{},
			err:  ErrEmpty,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewBill(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("got error %v, expected error %v", err, tt.err)
			}

			tt.want.Id = got.Id
			got.ScheduledDate = got.ScheduledDate.UTC().Truncate(24 * time.Hour)
			got.DueDate = got.DueDate.UTC().Truncate(24 * time.Hour)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v expected %v", got, tt.want)
			}
		})
	}
}
