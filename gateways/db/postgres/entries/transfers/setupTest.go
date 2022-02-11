package transfers

import (
	"context"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var TransfersTest = []transfers.Transfer{
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      accountdb.AccountsTest[1].Id,
		AccountDestinationId: accountdb.AccountsTest[0].Id,
		Amount:               100000,
		CreatedAt:            time.Now(),
	},
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      accountdb.AccountsTest[0].Id,
		AccountDestinationId: accountdb.AccountsTest[2].Id,
		Amount:               200000,
		CreatedAt:            time.Now(),
	},
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      accountdb.AccountsTest[0].Id,
		AccountDestinationId: accountdb.AccountsTest[1].Id,
		Amount:               300000,
		CreatedAt:            time.Now(),
	},
}

//insert list accounts in bacth
func CreateTransfersTest(pool *pgxpool.Pool) error {
	const statement = `INSERT INTO 
	transfers (
		id,
		account_origin_id,
		account_destination_id,
		amount,
		created_at
	)
		VALUES ($1, $2, $3, $4, $5)`

	batch := &pgx.Batch{}

	for _, t := range TransfersTest {
		batch.Queue(statement, t.Id, t.AccountOriginId, t.AccountDestinationId, t.Amount, t.CreatedAt)
	}

	br := pool.SendBatch(context.Background(), batch)
	defer br.Close()

	_, err := br.Exec()

	if err != nil {
		return err
	}

	return nil
}
