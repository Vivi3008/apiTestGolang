package account

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

type UsecaseMock struct {
	OnCreate       func(acc account.Account) (account.Account, error)
	OnStoreAccount func(account account.Account) error
	OnListAll      func() ([]account.Account, error)
	OnListById     func(accountId string) (account.Account, error)
	OnListByCpf    func(cpf string) (account.Account, error)
	OnUpdade       func(accountId string, value int, method MethodPayment) (account.Account, error)
	OnLogin        func(u account.Login) (string, error)
}

func (m UsecaseMock) CreateAccount(ctx context.Context, acc account.Account) (account.Account, error) {
	return m.OnCreate(acc)
}

func (m UsecaseMock) StoreAccount(ctx context.Context, account account.Account) error {
	return m.OnStoreAccount(account)
}

func (m UsecaseMock) ListAllAccounts(ctx context.Context) ([]account.Account, error) {
	return m.OnListAll()
}

func (m UsecaseMock) ListAccountById(ctx context.Context, accountId string) (account.Account, error) {
	return m.OnListById(accountId)
}

func (m UsecaseMock) ListAccountByCpf(ctx context.Context, cpf string) (account.Account, error) {
	return m.OnListByCpf(cpf)
}

func (m UsecaseMock) NewLogin(ctx context.Context, login account.Login) (string, error) {
	return m.OnLogin(login)
}

func (m UsecaseMock) UpdateAccountBalance(ctx context.Context, accountId string, value int, method MethodPayment) (account.Account, error) {
	return m.OnUpdade(accountId, value, method)
}
