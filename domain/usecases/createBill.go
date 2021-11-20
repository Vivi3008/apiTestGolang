package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrValueEmpty = errors.New("Value is empty")
)

//cria o pagamento e atualiza a conta
func (a Accounts) CreateBill(bill domain.Bill) (domain.Bill, error) {
	pay, err := domain.NewBill(bill)

	if err != nil {
		return domain.Bill{}, err
	}

	_, err = a.VerifyAccount(domain.AccountId(bill.AccountId), pay.Value, true)

	if err != nil {
		return domain.Bill{}, err
	}

	return pay, nil
}

// verifica se a conta tem saldo, e atualiza a conta
func (a Accounts) VerifyAccount(accountId domain.AccountId, value float64, debit bool) (domain.Account, error) {
	acc, err := a.ListAccountById(accountId)
	var actualBalance float64

	if err != nil {
		return domain.Account{}, err
	}

	if debit {
		actualBalance, err = debitFromAccount(acc.Balance, value)
	} else {
		actualBalance, err = creditToAccount(acc.Balance, value)
	}

	if err != nil {
		return domain.Account{}, err
	}

	updateAcc := domain.Account{
		Id:        acc.Id,
		Name:      acc.Name,
		Cpf:       acc.Cpf,
		Balance:   actualBalance,
		Secret:    acc.Secret,
		CreatedAt: acc.CreatedAt,
	}
	err = a.store.StoreAccount(updateAcc)

	if err != nil {
		return domain.Account{}, err
	}

	return updateAcc, nil
}

func debitFromAccount(balance float64, value float64) (float64, error) {
	if balance < value {
		return 0, ErrInsufficientLimit
	}
	actualBalance := balance - value

	return actualBalance, nil
}

func creditToAccount(balance float64, value float64) (float64, error) {
	if value <= 0 {
		return 0, ErrValueEmpty
	}

	actualBalance := balance + value

	return actualBalance, nil
}
