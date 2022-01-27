package account

import "context"

type AccountMock struct {
	OnCreate       func(acc Account) (Account, error)
	OnStoreAccount func(account Account) error
	OnListAll      func() ([]Account, error)
	OnListById     func(accountId string) (Account, error)
}

func (m AccountMock) CreateAccount(acc Account) (Account, error) {
	return m.OnCreate(acc)
}

func (m AccountMock) StoreAccount(ctx context.Context, account Account) error {
	return m.OnStoreAccount(account)
}

func (m AccountMock) ListAllAccounts(ctx context.Context) ([]Account, error) {
	return m.OnListAll()
}

func (m AccountMock) ListAccountById(ctx context.Context, accountId string) (Account, error) {
	return m.OnListById(accountId)
}
