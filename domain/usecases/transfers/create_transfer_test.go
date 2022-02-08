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

func TestCreateTransfer(t *testing.T) {
	t.Parallel()
	secretHash, _ := commom.GenerateHashPassword("16d5fs6a5f6")
	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Vanny",
		Cpf:     "55566689545",
		Secret:  secretHash,
		Balance: 2500,
	}
	personTarget := account.Account{
		Id:      uuid.New().String(),
		Name:    "Vivi",
		Cpf:     "544545454545",
		Secret:  secretHash,
		Balance: 3000,
	}

	type TestCase struct {
		name    string
		repo    transfers.TransferRepository
		accRepo account.AccountRepository
		args    transfers.Transfer
		want    transfers.Transfer
		err     error
	}

	testCases := []TestCase{
		{
			name: "Should create a transfer succesffull",
			repo: transfers.TransferMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					if accountId == person.Id {
						return person, nil
					}
					return personTarget, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					if id == person.Id {
						return account.Account{
							Id:        person.Id,
							Name:      person.Name,
							Cpf:       person.Cpf,
							Secret:    person.Secret,
							Balance:   1500,
							CreatedAt: person.CreatedAt,
						}, nil
					}
					return account.Account{
						Id:        personTarget.Id,
						Name:      personTarget.Name,
						Cpf:       personTarget.Cpf,
						Secret:    personTarget.Secret,
						Balance:   4000,
						CreatedAt: personTarget.CreatedAt,
					}, nil
				},
			},
			args: transfers.Transfer{
				AccountOriginId:      person.Id,
				AccountDestinationId: personTarget.Id,
				Amount:               1000,
			},
			want: transfers.Transfer{
				AccountOriginId:      person.Id,
				AccountDestinationId: personTarget.Id,
				Amount:               1000,
			},
			err: nil,
		},
		{
			name: "Fail if account origin haven't sufficient limit",
			repo: transfers.TransferMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					if accountId == "12345" {
						return account.Account{
							Id:        "12345",
							Name:      "test",
							Balance:   0,
							Secret:    "fjafdsai",
							CreatedAt: time.Now(),
						}, nil
					}
					return account.Account{
						Id:        "6789",
						Name:      "test2",
						Balance:   0,
						Secret:    "fjafdsai",
						CreatedAt: time.Now(),
					}, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, nil
				},
			},
			args: transfers.Transfer{
				AccountOriginId:      "12345",
				AccountDestinationId: "6789",
				Amount:               1000,
			},
			want: transfers.Transfer{},
			err:  accUse.ErrInsufficientLimit,
		},
		{
			name: "Fail if amount is empty or zero",
			repo: transfers.TransferMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					if accountId == "12345" {
						return account.Account{
							Id:        "12345",
							Name:      "test",
							Balance:   0,
							Secret:    "fjafdsai",
							CreatedAt: time.Now(),
						}, nil
					}
					return account.Account{
						Id:        "6789",
						Name:      "test2",
						Balance:   0,
						Secret:    "fjafdsai",
						CreatedAt: time.Now(),
					}, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, nil
				},
			},
			args: transfers.Transfer{
				AccountOriginId:      "12345",
				AccountDestinationId: "6789",
				Amount:               0,
			},
			want: transfers.Transfer{},
			err:  accUse.ErrValueEmpty,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ac := accUse.NewAccountUsecase(tt.accRepo)
			uc := NewTransferUsecase(tt.repo, ac)

			got, err := uc.CreateTransfer(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.Id = got.Id
			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
