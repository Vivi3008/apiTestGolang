package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) CreateTransfer(trans domain.Transfer) (domain.Transfer, error) {
	accountOrigin, _ := a.ListAccountById(domain.AccountId(trans.AccountOriginId))

	accountDestination, _ := a.ListAccountById(domain.AccountId(trans.AccountDestinationId))

	if accountOrigin.Balance < trans.Amount {
		return domain.Transfer{}, ErrInsufficientLimit
	}

	balanceOrigin := accountOrigin.Balance - accountDestination.Balance
	balanceDestiny := accountOrigin.Balance + accountDestination.Balance

	updateAccOrigin := domain.Account{
		Id:      string(accountOrigin.Id),
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

/* func (tr Tranfers) SaveTransfer(trans domain.Transfer) (domain.Transfer, error) {
	err := tr.storeTransfer.StoreTransfer(trans)

	if err != nil {
		return domain.Transfer{}, err
	}

	return trans, nil
} */
