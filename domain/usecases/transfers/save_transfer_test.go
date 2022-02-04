package transfers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

var ErrToSaveTest = errors.New("error to save transfer")

func TestSaveTransfer(t *testing.T) {
	t.Parallel()

	trans := transfers.Transfer{
		Id:                   uuid.New().String(),
		AccountOriginId:      uuid.New().String(),
		AccountDestinationId: uuid.New().String(),
		Amount:               2000,
		CreatedAt:            time.Now(),
	}

	type TestCase struct {
		name       string
		repository transfers.TransferRepository
		accRepo    account.AccountRepository
		args       transfers.Transfer
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should save a transfer succesfull",
			repository: transfers.TransferMock{
				OnStore: func(trans transfers.Transfer) error {
					return nil
				},
			},
			accRepo: account.AccountMock{},
			args:    trans,
			err:     nil,
		},
		{
			name: "Fail to save transfer",
			repository: transfers.TransferMock{
				OnStore: func(trans transfers.Transfer) error {
					return ErrToSaveTest
				},
			},
			accRepo: account.AccountMock{},
			args:    trans,
			err:     ErrToSaveTest,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ac := accUse.NewAccountUsecase(tt.accRepo)

			uc := NewTransferUsecase(tt.repository, ac)

			err := uc.SaveTransfer(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
