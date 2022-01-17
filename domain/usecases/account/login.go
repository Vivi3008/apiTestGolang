package account

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrCpfNotExists    = errors.New("this cpf doesn't exists")
	ErrInvalidPassword = errors.New("password invalid")
)

func (a AccountUsecase) NewLogin(u account.Login) (string, error) {
	listAccounts, _ := a.ListAllAccounts()

	var err error
	var result string

	for _, acc := range listAccounts {
		if acc.Cpf == u.Cpf {
			err = commom.VerifyPasswordHash(acc.Secret, u.Secret)
			if err != nil {
				return "", ErrInvalidPassword
			}
			result = acc.Id
		}
	}

	if result == "" {
		return result, ErrCpfNotExists
	}

	return result, nil
}
