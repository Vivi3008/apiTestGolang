package usecases

import (
	"errors"
	"fmt"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/google/uuid"
)

var (
	ErrInsufficientLimit = errors.New("Insufficient Limit")
)

func (s Tranfers) ListTransfer(originId domain.AccountOriginId) ([]domain.Transfer, error) {
	list, err := s.storeTransfer.ListTransfers(originId)

	if err != nil {
		return nil, fmt.Errorf("Could not list all accounts: %v\n", err)
	}

	return list, nil
}

func (a Accounts) Transfer(originId domain.AccountOriginId, destinationId domain.AccountDestinationId, amount float64) (domain.Transfer, error) {
	accountOrigin, _ := a.ListAccountById(domain.AccountId(originId))

	accountDestination, _ := a.ListAccountById(domain.AccountId(destinationId))

	if accountOrigin.Balance < amount {
		return domain.Transfer{}, ErrInsufficientLimit
	}

	balanceOrigin := accountOrigin.Balance - accountDestination.Balance
	balanceDestiny := accountOrigin.Balance + accountDestination.Balance

	updateAccOrigin := domain.Account{
		Id:      string(originId),
		Name:    accountOrigin.Name,
		Cpf:     accountOrigin.Cpf,
		Secret:  accountOrigin.Secret,
		Balance: balanceOrigin,
	}

	err := a.store.StoreAccount(updateAccOrigin)

	if err != nil {
		return domain.Transfer{}, err
	}

	updateAccDestiny := domain.Account{
		Id:      string(destinationId),
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
		Id:                   uuid.New().String(),
		AccountOriginId:      string(originId),
		AccountDestinationId: string(destinationId),
		Amount:               amount,
	}

	return transferOk, nil
}
