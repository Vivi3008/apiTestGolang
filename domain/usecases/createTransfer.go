package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (a Accounts) CreateTransfer(trans transfers.Transfer) (transfers.Transfer, error) {
	accountOrigin, err := a.VerifyAccount(trans.AccountOriginId, trans.Amount, Debit)

	if err != nil {
		return transfers.Transfer{}, err
	}
	accountDestination, err := a.VerifyAccount(trans.AccountDestinationId, trans.Amount, Credit)

	if err != nil {
		return transfers.Transfer{}, err
	}

	transferOk := transfers.Transfer{
		AccountOriginId:      accountOrigin.Id,
		AccountDestinationId: accountDestination.Id,
		Amount:               trans.Amount,
	}

	newtransfer, err := transfers.NewTransfer(transferOk)

	if err != nil {
		return transfers.Transfer{}, err
	}

	return newtransfer, nil
}
