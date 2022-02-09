package bills

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/jackc/pgx/v4"
)

func (r Repository) StoreBill(ctx context.Context, bill bills.Bill) error {
	const statement = `INSERT INTO bills (
		id,
		account_id,
		description,
		value,
		due_date,
		scheduled_date,
		status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := r.Db.QueryRow(ctx,
		statement,
		bill.Id,
		bill.AccountId,
		bill.Description,
		bill.Value,
		bill.DueDate,
		bill.ScheduledDate,
		bill.StatusBill).Scan()

	if err != pgx.ErrNoRows {
		return err
	}
	return nil
}
