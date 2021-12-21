package usecases

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/store"
)

type useCase struct {
	accRepository account.AccountRepository
}

type IUsecase interface {
	CreateAccount(account.Account) (account.Account, error)
	ListAllAccounts() ([]account.Account, error)
	ListAccountById(string) (account.Account, error)
}

func NewUseCase(ac account.AccountRepository) *useCase {
	return &useCase{ac}
}

type AccountStore struct {
	store store.AccountStore
}

func CreateAccountStore(store store.AccountStore) AccountStore {
	return AccountStore{
		store: store,
	}
}
