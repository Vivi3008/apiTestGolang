package activity

import (
	accountdb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	billDb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
	transferDb "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDbTest(pgx *pgxpool.Pool) error {
	err := accountdb.CreateAccountTest(pgx)

	if err != nil {
		return err
	}

	err = transferDb.CreateTransfersTest(pgx)

	if err != nil {
		return err
	}

	err = billDb.CreateBillsTest(pgx)

	if err != nil {
		return err
	}
	return nil
}
