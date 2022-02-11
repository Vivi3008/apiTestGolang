package bills

import (
	"context"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Bls = []bills.Bill{
	{
		Id:            uuid.NewString(),
		AccountId:     accountdb.AccountsTest[0].Id,
		Description:   "Academia",
		Value:         11000,
		DueDate:       time.Now().AddDate(0, 0, 5),
		StatusBill:    bills.Pago,
		ScheduledDate: time.Now(),
	},
	{
		Id:            uuid.NewString(),
		AccountId:     accountdb.AccountsTest[0].Id,
		Description:   "Fatura Cartão",
		Value:         50000,
		DueDate:       time.Now().AddDate(0, 0, 4),
		StatusBill:    bills.Negado,
		ScheduledDate: time.Now(),
	},
	{
		Id:            uuid.NewString(),
		AccountId:     accountdb.AccountsTest[0].Id,
		Description:   "IPTU",
		Value:         100000,
		DueDate:       time.Now().AddDate(0, 0, 3),
		StatusBill:    bills.Pago,
		ScheduledDate: time.Now(),
	},
	{
		Id:            uuid.NewString(),
		AccountId:     accountdb.AccountsTest[1].Id,
		Description:   "Material escolar",
		Value:         200000,
		DueDate:       time.Now().AddDate(0, 0, 3),
		StatusBill:    bills.Pago,
		ScheduledDate: time.Now(),
	},
}

func CreateBillsTest(pool *pgxpool.Pool) error {
	const statement = `INSERT INTO bills (
		id,
		account_id,
		description,
		value,
		due_date,
		scheduled_date,
		status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	batch := &pgx.Batch{}

	for _, bl := range Bls {
		batch.Queue(statement, bl.Id, bl.AccountId, bl.Description, bl.Value, bl.DueDate, bl.ScheduledDate, bl.StatusBill)
	}

	br := pool.SendBatch(context.Background(), batch)
	defer br.Close()

	_, err := br.Exec()

	if err != nil {
		return err
	}

	return nil
}
