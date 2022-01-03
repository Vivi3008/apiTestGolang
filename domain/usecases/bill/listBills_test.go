package bill

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

var (
	ErrListBills = errors.New("error to list bills")
)

func TestBills(t *testing.T) {
	t.Parallel()

	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Dfadfsa",
		Cpf:     "55566689545",
		Secret:  "dafd33255",
		Balance: 2500,
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
		accRepo    account.AccountMock
		args       string
		want       []bills.Bill
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should list payments from account id",
			repository: bills.BillMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return []bills.Bill{payment}, nil
				},
			},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{person}, nil
				},
			},
			args: person.Id,
			want: []bills.Bill{payment},
			err:  nil,
		},
		{
			name: "Fail if id account doens't exists",
			repository: bills.BillMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return []bills.Bill{payment}, nil
				},
			},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return account.Account{}, accUse.ErrListAccountEmpty
				},
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{}, nil
				},
			},
			args: person.Id,
			want: []bills.Bill{},
			err:  accUse.ErrListAccountEmpty,
		},
		{
			name: "Fail if error to list bills",
			repository: bills.BillMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return []bills.Bill{}, ErrListBills
				},
			},
			accRepo: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return person, nil
				},
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{person}, nil
				},
			},
			args: person.Id,
			want: []bills.Bill{},
			err:  ErrListBills,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ac := accUse.NewAccountUsecase(tt.accRepo)
			uc := NewBillUseCase(tt.repository, ac)

			got, err := uc.ListAllBills(tt.args)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

	/* 	t.Run("Should create a bill successfully", func(t *testing.T) {
		billStore := store.NewBillStore()
		bls := CreateNewBill(billStore)

		accountStore := store.NewAccountStore()
		accounts := CreateNewAccount(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "55566689545",
			Secret:  "dafd33255",
			Balance: 2500,
		}

		account, err := accounts.CreateAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		dueDate := time.Now().AddDate(0, 0, 2)

		bill := bills.Bill{
			AccountId:   account.Id,
			Description: "Conta internet",
			Value:       150,
			DueDate:     dueDate,
		}

		newBill, err := bills.NewBill(bill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		billOk, err := CreateNewBill(newBill)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		_, err = bls.SaveBill(billOk)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		// verificando se debitou o valor na conta
		acc, err := accountStore.ListAccountById(account.Id)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		if acc.Balance != 2350 {
			t.Errorf("Expected %v, got %v", 2350, acc.Balance)
		}
	}) */
}
