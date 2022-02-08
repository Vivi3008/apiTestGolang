package transfers

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

func TestTransfers(t *testing.T) {
	t.Parallel()
	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Vanny",
		Cpf:     "55566689545",
		Secret:  "dafd33255",
		Balance: 2500,
	}
	secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
	listAccounts := []account.Account{
		{
			Id:        "fads1fdsa3",
			Name:      "David",
			Cpf:       "556565656555",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
		{
			Id:        "5df4s5df45",
			Name:      "Vale",
			Cpf:       "656565656565",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
		{
			Id:        "a6fd56sad5f3",
			Name:      "Biscui",
			Cpf:       "21545454545",
			Secret:    secretHash,
			Balance:   260000,
			CreatedAt: time.Now(),
		},
	}

	trans := transfers.Transfer{
		Id:                   uuid.New().String(),
		AccountOriginId:      person.Id,
		AccountDestinationId: "asdf65asd6f5sa",
		Amount:               100,
		CreatedAt:            time.Now(),
	}

	type TestCase struct {
		name       string
		repository transfers.TransferRepository
		accRepo    account.AccountRepository
		args       string
		want       []transfers.Transfer
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should list a transfer successfull",
			repository: transfers.TransferMock{
				OnListAll: func(id string) ([]transfers.Transfer, error) {
					return []transfers.Transfer{trans}, nil
				},
			},
			args: listAccounts[0].Id,
			want: []transfers.Transfer{trans},
			err:  nil,
		},
		{
			name: "Should list empty transfers",
			repository: transfers.TransferMock{
				OnListAll: func(id string) ([]transfers.Transfer, error) {
					return []transfers.Transfer{}, nil
				},
			},
			args: listAccounts[1].Id,
			want: []transfers.Transfer{},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ac := accUse.NewAccountUsecase(tt.accRepo)

			transUse := NewTransferUsecase(tt.repository, ac)

			got, err := transUse.ListTransfer(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
