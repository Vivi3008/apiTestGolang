package accountdb

import (
	"context"
	"errors"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (r Repository) UpdateAccount(ctx context.Context, balance int, id string) (entities.Account, error) {
	const statement = `UPDATE accounts
		SET balance=$1 WHERE id=$2
		RETURNING
			id,
			name,
			cpf,
			balance,
			created_at
		;`

	var account entities.Account

	err := r.DB.QueryRow(ctx, statement, balance, id).Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)

	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.SQLState() == "23514" {
				return entities.Account{}, ErrBalanceInvalid
			}

		}
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Account{}, ErrIdNotExists
		}
		return entities.Account{}, err
	}

	return account, nil
}
