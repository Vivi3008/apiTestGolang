package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) CreateTransfer(trans domain.Transfer) (domain.Transfer, error) {
	accountOrigin, err := a.VerifyAccount(trans.AccountOriginId, trans.Amount, Debit)

	if err != nil {
		return domain.Transfer{}, err
	}
	accountDestination, err := a.VerifyAccount(trans.AccountDestinationId, trans.Amount, Credit)

	if err != nil {
		return domain.Transfer{}, err
	}

	transferOk := domain.Transfer{
		AccountOriginId:      accountOrigin.Id,
		AccountDestinationId: accountDestination.Id,
		Amount:               trans.Amount,
	}

	newtransfer, err := domain.NewTransfer(transferOk)

	if err != nil {
		return domain.Transfer{}, err
	}

	return newtransfer, nil
}
