package account

import (
	"context"
	"errors"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/jackc/pgconn"
)

var (
	ErrCpfExists = errors.New("this cpf already exists")
)

func (r Repository) StoreAccount(ctx context.Context, account entities.Account) error {
	statement := `INSERT INTO accounts
		(id,
		 name,
		 cpf,
		 secret,
		 balance)
		VALUES ($1, $2, $3, $4, $5)
		returning created_at`

	err := r.DB.QueryRow(ctx, statement,
		account.Id,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
	).Scan(&account.CreatedAt)

	var pgError *pgconn.PgError

	if errors.As(err, &pgError) {
		if pgError.SQLState() == "23505" {
			return ErrCpfExists
		}
	}

	return err
}
