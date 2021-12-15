package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrValueEmpty = errors.New("value is empty")
)

//cria o pagamento e atualiza a conta
func (a Accounts) CreateBill(bill domain.Bill) (domain.Bill, error) {
	pay, err := domain.NewBill(bill)

	if err != nil {
		return domain.Bill{}, err
	}

	_, err = a.VerifyAccount(bill.AccountId, pay.Value, Debit)

	if err != nil {
		return domain.Bill{}, err
	}

	return pay, nil
}
