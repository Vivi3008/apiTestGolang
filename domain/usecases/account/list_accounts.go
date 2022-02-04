package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrIdNotExists      = errors.New("id doesn't exists")
	ErrListAccountEmpty = errors.New("list account is empty")
)

func (a AccountUsecase) ListAllAccounts(ctx context.Context) ([]account.Account, error) {
	list, err := a.repo.ListAllAccounts(ctx)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a AccountUsecase) ListAccountById(ctx context.Context, id string) (account.Account, error) {
	acc, err := a.repo.ListAccountById(ctx, id)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}

func (a AccountUsecase) ListAccountByCpf(ctx context.Context, cpf string) (account.Account, error) {
	acc, err := a.repo.ListAccountByCpf(ctx, cpf)

	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
