package account

import (
	"context"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (r Repository) ListAllAccounts(ctx context.Context) ([]entities.Account, error) {
	const statement = `SELECT id,
	   		name,
	   		cpf,
	   		balance,
	   		created_at FROM accounts
	   		`

	var accounts []entities.Account

	rows, err := r.DB.Query(ctx, statement)

	if err != nil {
		return []entities.Account{}, err
	}

	var account entities.Account

	for rows.Next() {
		err = rows.Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)
		if err != nil {
			return []entities.Account{}, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
