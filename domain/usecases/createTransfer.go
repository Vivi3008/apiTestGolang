package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (a Accounts) CreateTransfer(trans account.Transfer) (account.Transfer, error) {
	accountOrigin, err := a.VerifyAccount(trans.AccountOriginId, trans.Amount, Debit)

	if err != nil {
		return account.Transfer{}, err
	}
	accountDestination, err := a.VerifyAccount(trans.AccountDestinationId, trans.Amount, Credit)

	if err != nil {
		return account.Transfer{}, err
	}

	transferOk := account.Transfer{
		AccountOriginId:      accountOrigin.Id,
		AccountDestinationId: accountDestination.Id,
		Amount:               trans.Amount,
	}

	newtransfer, err := account.NewTransfer(transferOk)

	if err != nil {
		return account.Transfer{}, err
	}

	return newtransfer, nil
}
