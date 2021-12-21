package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

var (
	ErrValueEmpty = errors.New("value is empty")
)

//cria o pagamento e atualiza a conta
func (a Accounts) CreateBill(bill bills.Bill) (bills.Bill, error) {
	pay, err := bills.NewBill(bill)

	if err != nil {
		return bills.Bill{}, err
	}

	_, err = a.VerifyAccount(bill.AccountId, pay.Value, Debit)

	if err != nil {
		return bills.Bill{}, err
	}

	return pay, nil
}
