package transfers

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

func (r Repository) ListTransfer(ctx context.Context, id string) ([]transfers.Transfer, error) {
	const statement = `SELECT * FROM transfers WHERE account_origin_id=$1`

	var listTransfers = make([]transfers.Transfer, 0)

	rows, err := r.Db.Query(ctx, statement, id)

	if err != nil {
		return []transfers.Transfer{}, err
	}

	var transfer transfers.Transfer

	for rows.Next() {
		err := rows.Scan(&transfer.Id, &transfer.AccountOriginId, &transfer.AccountDestinationId, &transfer.Amount, &transfer.CreatedAt)

		if err != nil {
			return []transfers.Transfer{}, err
		}
		listTransfers = append(listTransfers, transfer)
	}

	return listTransfers, nil
}
