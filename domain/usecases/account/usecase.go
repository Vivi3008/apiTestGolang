package account

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

type AccountUsecase struct {
	repo account.AccountRepository
}

type Usecase interface {
	CreateAccount(ctx context.Context, account account.Account) (account.Account, error)
	ListAllAccounts(ctx context.Context) ([]account.Account, error)
	ListAccountById(ctx context.Context, accountId string) (account.Account, error)
	NewLogin(ctx context.Context, login account.Login) (string, error)
}

func NewAccountUsecase(acc account.AccountRepository) AccountUsecase {
	return AccountUsecase{
		repo: acc,
	}
}
