package transfers

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

var ErrIdDestiny = errors.New("account destiny id can't be the same account origin id")

func (a TranfersUsecase) CreateTransfer(ctx context.Context, trans transfers.Transfer) (transfers.Transfer, error) {
	if trans.AccountOriginId == trans.AccountDestinationId {
		return transfers.Transfer{}, ErrIdDestiny
	}

	accountOrigin, err := a.accUsecase.UpdateAccountBalance(ctx, trans.AccountOriginId, trans.Amount, account.Debit)

	if err != nil {
		return transfers.Transfer{}, err
	}

	accountDestination, err := a.accUsecase.UpdateAccountBalance(ctx, trans.AccountDestinationId, trans.Amount, account.Credit)

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
