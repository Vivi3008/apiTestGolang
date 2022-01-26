package account

import (
	entities "github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

func (r Repository) ListAllAccounts() ([]entities.Account, error) {
	/* 	statement := `SELECT id,
	   		name,
	   		cpf,
	   		balance,
	   		created_at FROM accounts
	   		`

	   	var accounts []entities.Account

	   	rows := r.DB.QueryRowContext(ctx, statement)

	   	for rows {
	   	} */
	return []entities.Account{}, nil
}

func (r Repository) ListAccountById(id string) (entities.Account, error) {
	return entities.Account{}, nil
}
