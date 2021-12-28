package account

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

type AccountUsecase struct {
	repo account.AccountRepository
}

var (
	ErrInsufficientLimit = errors.New("insufficient Limit")
	ErrValueEmpty        = errors.New("value is empty")
)

type MethodPayment string

const (
	Debit  MethodPayment = "Débito"
	Credit MethodPayment = "Crédito"
)

func NewAccountUsecase(acc account.AccountRepository) AccountUsecase {
	return AccountUsecase{
		repo: acc,
	}
}

func (a AccountUsecase) VerifyAccount(accountId string, value int, method MethodPayment) (account.Account, error) {
	acc, err := a.repo.ListAccountById(accountId)
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
	err = a.repo.StoreAccount(updateAcc)

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
