package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

//cria o pagamento e atualiza a conta
func (a Accounts) CreateBill(bill domain.Bill) (domain.Bill, error) {
	pay, err := domain.NewBill(bill)

	if err != nil {
		return domain.Bill{}, err
	}

	_, err = a.VerifyAccount(domain.AccountId(bill.AccountId), pay.Value)

	if err != nil {
		return domain.Bill{}, err
	}

	return pay, nil
}

// verifica se a conta tem saldo, e atualiza a conta
func (a Accounts) VerifyAccount(accountId domain.AccountId, value float64) (domain.Account, error) {
	acc, err := a.ListAccountById(accountId)

	if err != nil {
		return domain.Account{}, err
	}

	if acc.Balance < value {
		return domain.Account{}, ErrInsufficientLimit
	}

	actualBalance := acc.Balance - value

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
