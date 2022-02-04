package account

import "context"

type AccountRepository interface {
	StoreAccount(ctx context.Context, account Account) error
	ListAllAccounts(ctx context.Context) ([]Account, error)
	ListAccountById(ctx context.Context, accountId string) (Account, error)
	ListAccountByCpf(ctx context.Context, cpf string) (Account, error)
	UpdateAccount(ctx context.Context, balance int, id string) (Account, error)
}
