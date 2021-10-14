package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrIdOriginNotExists  = errors.New("Account origin id not exists")
	ErrIdDestinyNotExists = errors.New("Account destiny id not exists")
)

func (a Accounts) CreateTransfer(trans domain.Transfer) (domain.Transfer, error) {
	accountOrigin, err := a.ListAccountById(domain.AccountId(trans.AccountOriginId))

	if err != nil {
		return domain.Transfer{}, ErrIdOriginNotExists
	}

	accountDestination, err := a.ListAccountById(domain.AccountId(trans.AccountDestinationId))

	if err != nil {
		return domain.Transfer{}, ErrIdDestinyNotExists
	}

	if accountOrigin.Balance < trans.Amount {
		return domain.Transfer{}, ErrInsufficientLimit
	}

	balanceOrigin := accountOrigin.Balance - trans.Amount
	balanceDestiny := accountDestination.Balance + trans.Amount

	updateAccOrigin := domain.Account{
		Id:      string(accountOrigin.Id),
		Name:    accountOrigin.Name,
		Cpf:     accountOrigin.Cpf,
		Secret:  accountOrigin.Secret,
		Balance: balanceOrigin,
	}

	err = a.store.StoreAccount(updateAccOrigin)

	if err != nil {
		return domain.Transfer{}, err
	}

	updateAccDestiny := domain.Account{
		Id:      string(accountDestination.Id),
		Name:    accountDestination.Name,
		Cpf:     accountDestination.Cpf,
		Secret:  accountDestination.Secret,
		Balance: balanceDestiny,
	}

	err = a.store.StoreAccount(updateAccDestiny)

	if err != nil {
		return domain.Transfer{}, err
	}

	transferOk := domain.Transfer{
		AccountOriginId:      string(trans.AccountOriginId),
		AccountDestinationId: string(trans.AccountDestinationId),
		Amount:               trans.Amount,
	}

	newtransfer, err := domain.NewTransfer(transferOk)

	if err != nil {
		return domain.Transfer{}, err
	}

	return newtransfer, nil
}
