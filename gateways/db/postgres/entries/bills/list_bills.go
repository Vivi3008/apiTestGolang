package bills

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (r Repository) ListBills(ctx context.Context, accountId string) ([]bills.Bill, error) {
	const statement = `SELECT
		id, 
		account_id,
		description,
		value,
		due_date,
		scheduled_date,
		status,
		created_at FROM bills WHERE account_id=$1
		ORDER BY created_at asc`

	var listBills = make([]bills.Bill, 0)

	rows, err := r.Db.Query(ctx, statement, accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	var bill bills.Bill

	for rows.Next() {
		err := rows.Scan(
			&bill.Id,
			&bill.AccountId,
			&bill.Description,
			&bill.Value,
			&bill.DueDate,
			&bill.ScheduledDate,
			&bill.StatusBill,
			&bill.CreatedAt,
		)
		if err != nil {
			return []bills.Bill{}, err
		}
		listBills = append(listBills, bill)
	}
	return listBills, err
}
