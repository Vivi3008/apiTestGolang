package bill

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

func TestCreateBill(t *testing.T) {
	t.Parallel()

	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Dfadfsa",
		Cpf:     "55566689545",
		Secret:  "dafd33255",
		Balance: 25000,
	}

	payment := bills.Bill{
		Id:            uuid.New().String(),
		Description:   "Academia",
		Value:         5300,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
	}

	type TestCase struct {
		name       string
		repository bills.BillRepository
		accRepo    account.AccountRepository
		args       bills.Bill
		want       bills.Bill
		err        error
	}

	testCases := []TestCase{
		{
			name:       "Should create a payment successfull",
			repository: bills.BillMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{
						Id:      uuid.New().String(),
						Name:    "Dfadfsa",
						Cpf:     "55566689545",
						Secret:  "dafd33255",
						Balance: 19700,
					}, nil
				},
			},
			args: payment,
			want: payment,
			err:  nil,
		},
		{
			name:       "Fail if account has insufficient limit",
			repository: bills.BillMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, accUse.ErrInsufficientLimit
				},
			},
			args: payment,
			want: bills.Bill{},
			err:  accUse.ErrInsufficientLimit,
		},
		{
			name:       "Create bill if scheduled date is yesterday",
			repository: bills.BillMock{},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{
						Id:      person.Id,
						Name:    "Dfadfsa",
						Cpf:     "55566689545",
						Secret:  "dafd33255",
						Balance: 50000,
					}, nil
				},
			},
			args: bills.Bill{
				Id:            uuid.New().String(),
				Description:   "Academia",
				Value:         5300,
				DueDate:       time.Now().AddDate(0, 0, 2),
				ScheduledDate: time.Now().AddDate(0, 0, -1),
			},
			want: bills.Bill{
				Id:            uuid.New().String(),
				Description:   "Academia",
				Value:         5300,
				DueDate:       time.Now().AddDate(0, 0, 2),
				ScheduledDate: time.Now(),
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ac := accUse.NewAccountUsecase(tt.accRepo)
			uc := NewBillUseCase(tt.repository, ac)

			got, err := uc.CreateBill(context.Background(), tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			tt.want.Id = got.Id
			tt.want.ScheduledDate = got.ScheduledDate
			tt.want.DueDate = got.DueDate
			tt.want.StatusBill = got.StatusBill

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
