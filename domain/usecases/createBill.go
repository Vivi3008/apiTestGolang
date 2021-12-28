package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

//cria o pagamento e atualiza a conta
func (a AccountUsecase) CreateBill(bill bills.Bill) (bills.Bill, error) {
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
