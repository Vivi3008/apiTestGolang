package account

import (
	"context"
	"errors"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/jackc/pgx/v4"
)

var ErrCpfNotExists = errors.New("this cpf does not exist")

func (r Repository) ListAccountByCpf(ctx context.Context, cpf string) (entities.Account, error) {
	const statement = `SELECT id,
	name,
	cpf,
	balance,
	created_at FROM accounts
	WHERE cpf=$1`

	var account entities.Account

	err := r.DB.QueryRow(ctx, statement, cpf).Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)

	switch {
	case err == pgx.ErrNoRows:
		return entities.Account{}, ErrCpfNotExists
	case err != nil:
		return entities.Account{}, err
	}

	return account, nil
}
