package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain"
)

func (a Accounts) CreateTransfer(trans domain.Transfer) (domain.Transfer, error) {
	accountOrigin, err := a.VerifyAccount(domain.AccountId(trans.AccountOriginId), trans.Amount, true)

	if err != nil {
		return domain.Transfer{}, err
	}

	accountDestination, err := a.VerifyAccount(domain.AccountId(trans.AccountDestinationId), trans.Amount, false)

	if err != nil {
		return domain.Transfer{}, err
	}

	transferOk := domain.Transfer{
		AccountOriginId:      string(accountOrigin.Id),
		AccountDestinationId: string(accountDestination.Id),
		Amount:               trans.Amount,
	}

	newtransfer, err := domain.NewTransfer(transferOk)

	if err != nil {
		return domain.Transfer{}, err
	}

	return newtransfer, nil
}
