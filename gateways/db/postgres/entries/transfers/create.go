package transfers

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/jackc/pgconn"
)

const (
	TransfersIdEquals = "transfers_check"
)

var (
	ErrIdEquals = errors.New("account destiny id can't be the same account origin id")
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

	if errors.As(err, &pgError) {
		switch pgError.ConstraintName {
		case TransfersIdEquals:
			return ErrIdEquals
		}
	}

	return nil
}
