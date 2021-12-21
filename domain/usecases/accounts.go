package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

type Accounts struct {
	accs useCase
}

func CreateNewAccount(acc useCase) Accounts {
	return Accounts{
		accs: acc,
	}
}

func (a Accounts) VerifyAccount(accountId string, value int, method MethodPayment) (account.Account, error) {
	acc, err := a.ListAccountById(accountId)
	var actualBalance int

	if err != nil {
		return account.Account{}, err
	}

	actualBalance, err = ModifyBalanceAccount(acc.Balance, value, method)

	if err != nil {
		return account.Account{}, err
	}

	updateAcc := account.Account{
		Id:        acc.Id,
		Name:      acc.Name,
		Cpf:       acc.Cpf,
		Balance:   actualBalance,
		Secret:    acc.Secret,
		CreatedAt: acc.CreatedAt,
	}
	err = a.accs.accRepository.StoreAccount(updateAcc)

	if err != nil {
		return account.Account{}, err
	}

	return updateAcc, nil
}

func ModifyBalanceAccount(balance int, value int, method MethodPayment) (int, error) {
	if method == Debit {
		if balance < value {
			return 0, ErrInsufficientLimit
		}

		return balance - value, nil
	} else {
		if value <= 0 {
			return 0, ErrValueEmpty
		}

		return balance + value, nil
	}
}
