package account

import (
	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (r Repository) ListAllAccounts() ([]entities.Account, error) {
	statement := `SELECT id,
	   		name,
	   		cpf,
	   		balance,
	   		created_at FROM accounts
	   		`

	var accounts []entities.Account

	rows, err := r.DB.Query(statement)

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
