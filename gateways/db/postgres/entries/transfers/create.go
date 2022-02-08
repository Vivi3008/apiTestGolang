package transfers

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

const (
	TransfersIdEquals        = "transfers_check"
	AccountOriginIdFkey      = "transfers_account_origin_id_fkey"
	AccountDestinationIdFkey = "transfers_account_destination_id_fkey"
	AmountCheck              = "transfers_amount_check"
)

var (
	ErrIdEquals          = errors.New("account destiny id can't be the same account origin id")
	ErrIdOriginNotExist  = errors.New("account origin id don't exist")
	ErrIdDestinyNotExist = errors.New("account destiny id don't exist")
	ErrAmountInvalid     = errors.New("amount can't be less than zero")
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

	var pgError *pgconn.PgError

	if err != pgx.ErrNoRows {
		if errors.As(err, &pgError) {
			switch pgError.ConstraintName {
			case TransfersIdEquals:
				return ErrIdEquals
			case AccountOriginIdFkey:
				return ErrIdOriginNotExist
			case AccountDestinationIdFkey:
				return ErrIdDestinyNotExist
			case AmountCheck:
				return ErrAmountInvalid
			}
		}
		return err
	}
	return nil
}
