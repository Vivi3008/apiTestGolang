package account

import (
	"context"
	"errors"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/jackc/pgconn"
)

const (
	AccountsKeyUnique    = "accounts_pkey"
	AccountsBalanceCheck = "accounts_balance_check"
)

func (r Repository) StoreAccount(ctx context.Context, account entities.Account) error {
	const statement = `INSERT INTO 
	accounts (
		id,
		name,
		cpf,
		secret,
		balance,
		created_at
	)
		VALUES ($1, $2, $3, $4, $5, $6)`

	err := r.DB.QueryRow(ctx, statement,
		account.Id,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
		account.CreatedAt,
	).Scan()

	var pgError *pgconn.PgError

	if errors.As(err, &pgError) {
		switch pgError.ConstraintName {
		case AccountsKeyUnique:
			return ErrCpfExists
		case AccountsBalanceCheck:
			return ErrBalanceInvalid
		}
	}

	return nil
}
