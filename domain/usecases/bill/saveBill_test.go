package bill

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

var (
	ErrToSave = errors.New("error to save in database")
)

func TestSaveBill(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository bills.BillRepository
		accRepo    account.AccountRepository
		args       bills.Bill
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should save bill successfull",
			repository: bills.BillMock{
				OnStore: func(b bills.Bill) error {
					return nil
				},
			},
			accRepo: account.AccountMock{},
			args: bills.Bill{
				Id:            uuid.New().String(),
				Description:   "Academia",
				Value:         5300,
				DueDate:       time.Now().AddDate(0, 0, 2),
				ScheduledDate: time.Now(),
			},
			err: nil,
		},
		{
			name: "Fail to store bill in database",
			repository: bills.BillMock{
				OnStore: func(b bills.Bill) error {
					return ErrToSave
				},
			},
			accRepo: account.AccountMock{},
			args: bills.Bill{
				Id:            uuid.New().String(),
				Description:   "Academia",
				Value:         5300,
				DueDate:       time.Now().AddDate(0, 0, 2),
				ScheduledDate: time.Now(),
			},
			err: ErrToSave,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ac := accUse.NewAccountUsecase(tt.accRepo)
			uc := NewBillUseCase(tt.repository, ac)

			err := uc.SaveBill(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
