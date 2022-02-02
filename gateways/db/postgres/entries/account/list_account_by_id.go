package account

import (
	"context"
	"database/sql"
	"errors"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var ErrIdNotExists = errors.New("id does not exist")

func (r Repository) ListAccountById(ctx context.Context, id string) (entities.Account, error) {
	const statement = `SELECT id,
		name,
		cpf,
		balance,
		created_at FROM accounts
		WHERE id=$1`

	var account entities.Account

	err := r.DB.QueryRow(ctx, statement, id).Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)

	switch {
	case err == sql.ErrNoRows:
		return entities.Account{}, ErrIdNotExists
	case err != nil:
		return entities.Account{}, err
	}

	return account, nil
}
