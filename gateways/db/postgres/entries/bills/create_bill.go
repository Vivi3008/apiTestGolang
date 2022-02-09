package bills

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

const (
	BillAccountIdConstraint = "bills_account_id_fkey"
	BillValueCheck          = "bills_value_check"
)

var (
	ErrAccountIdNotExist = errors.New("account id don't exist")
	ErrBillIdEmpty       = errors.New("bill id can't be empty")
	ErrValueInvalid      = errors.New("value can't be less than zero")
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
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			switch {
			case pgError.ConstraintName == BillAccountIdConstraint:
				return ErrAccountIdNotExist
			case pgError.SQLState() == "22P02":
				return ErrBillIdEmpty
			case pgError.ConstraintName == BillValueCheck:
				return ErrValueInvalid
			default:
				return err
			}
		}
	}
	return nil
}
