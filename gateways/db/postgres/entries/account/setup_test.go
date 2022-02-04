package account

import (
	"context"
	"time"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var accountsTest = []entities.Account{
	{
		Id:        uuid.NewString(),
		Name:      "Teste1",
		Cpf:       "13256589412",
		Balance:   16565,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste2",
		Cpf:       "132565889412",
		Balance:   16565,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste3",
		Cpf:       "13256589712",
		Balance:   16565,
		CreatedAt: time.Now(),
	},
}

//insert list accounts in bacth
func createAccountTest(pool *pgxpool.Pool) error {
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

	batch := &pgx.Batch{}

	for _, acc := range accountsTest {
		batch.Queue(statement, acc.Id, acc.Name, acc.Cpf, acc.Secret, acc.Balance, acc.CreatedAt)
	}

	br := pool.SendBatch(context.Background(), batch)
	defer br.Close()

	_, err := br.Exec()

	if err != nil {
		return err
	}

	return nil
}
