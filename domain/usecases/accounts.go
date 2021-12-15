package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/Vivi3008/apiTestGolang/store"
)

type Accounts struct {
	store store.AccountStore
}

func CreateNewAccount(store store.AccountStore) Accounts {
	return Accounts{
		store: store,
	}
}

func (a Accounts) VerifyAccount(accountId string, value int, method MethodPayment) (domain.Account, error) {
	acc, err := a.ListAccountById(accountId)
	var actualBalance int

	if err != nil {
		return domain.Account{}, err
	}

	actualBalance, err = modifyBalanceAccount(acc.Balance, value, method)

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

func modifyBalanceAccount(balance int, value int, method MethodPayment) (int, error) {
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
