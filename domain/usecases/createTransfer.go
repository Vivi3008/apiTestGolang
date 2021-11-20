package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrIdOriginNotExists  = errors.New("Account origin id doesn't exists")
	ErrIdDestinyNotExists = errors.New("Account destiny id doesn't exists")
)

func (a Accounts) CreateTransfer(trans domain.Transfer) (domain.Transfer, error) {
	accountOrigin, err := a.VerifyAccount(domain.AccountId(trans.AccountOriginId), trans.Amount, true)

	if err != nil {
		return domain.Transfer{}, ErrIdOriginNotExists
	}

	accountDestination, err := a.VerifyAccount(domain.AccountId(trans.AccountDestinationId), trans.Amount, false)

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
