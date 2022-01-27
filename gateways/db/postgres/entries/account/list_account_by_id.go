package account

import (
	"context"
	"database/sql"
	"fmt"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (r Repository) ListAccountById(ctx context.Context, id string) (entities.Account, error) {
	statement := `SELECT id,
		name,
		cpf,
		balance,
		created_at FROM accounts
		WHERE id=?`

	var account entities.Account

	err := r.DB.QueryRow(ctx, statement, id).Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)

	switch {
	case err == sql.ErrNoRows:
		return entities.Account{}, fmt.Errorf("no account with id %s", id)
	case err != nil:
		return entities.Account{}, fmt.Errorf("query error: %s", err)
	}

	return account, nil
}
