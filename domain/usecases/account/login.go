package account

import (
	"context"
	"errors"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrCpfNotExists    = errors.New("this cpf doesn't exists")
	ErrInvalidPassword = errors.New("password invalid")
)

func (a AccountUsecase) NewLogin(ctx context.Context, u account.Login) (string, error) {
	account, err := a.repo.ListAccountByCpf(ctx, u.Cpf)

	if err != nil {
		return "", err
	}

	err = commom.VerifyPasswordHash(account.Secret, u.Secret)

	if err != nil {
		return "", ErrInvalidPassword
	}

	return account.Id, nil
}
