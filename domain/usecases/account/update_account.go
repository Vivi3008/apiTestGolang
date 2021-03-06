package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrInsufficientLimit = errors.New("insufficient Limit")
	ErrValueEmpty        = errors.New("value is empty")
)

type MethodPayment string

const (
	Debit  MethodPayment = "Débito"
	Credit MethodPayment = "Crédito"
)

func (a AccountUsecase) UpdateAccountBalance(ctx context.Context, accountId string, value int, method MethodPayment) (account.Account, error) {
	acc, err := a.repo.ListAccountById(ctx, accountId)
	var actualBalance int

	if err != nil {
		return account.Account{}, err
	}

	actualBalance, err = ModifyBalanceAccount(acc.Balance, value, method)

	if err != nil {
		return account.Account{}, err
	}

	updatedAcc, err := a.repo.UpdateAccount(ctx, actualBalance, acc.Id)
	if err != nil {
		return account.Account{}, err
	}

	return updatedAcc, nil
}

func ModifyBalanceAccount(balance int, value int, method MethodPayment) (int, error) {
	if value <= 0 {
		return 0, ErrValueEmpty
	}

	if method == Debit {
		if balance < value {
			return 0, ErrInsufficientLimit
		}
		return balance - value, nil
	}

	return balance + value, nil
}
