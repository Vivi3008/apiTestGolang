package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (r Repository) SaveTransfer(ctx context.Context, transfer transfers.Transfer) error {
	const statement = `INSERT INTO 
	transfers (
		id,
		account_origin_id,
		account_destination_id,
		amount,
		created_at
	)
		VALUES ($1, $2, $3, $4, $5)`

	err := r.Db.QueryRow(ctx,
		statement,
		transfer.Id,
		transfer.AccountOriginId,
		transfer.AccountDestinationId,
		transfer.Amount,
		transfer.CreatedAt,
	).Scan()

	if err != nil {
		return err
	}
	/* 	var pgError *pgconn.PgError

	   	if errors.As(err, &pgError) {
	   		switch pgError.ConstraintName {
	   		case AccountsKeyUnique:
	   			return ErrCpfExists
	   		case AccountsBalanceCheck:
	   			return ErrBalanceInvalid
	   		}
	   	} */

	return nil
}
