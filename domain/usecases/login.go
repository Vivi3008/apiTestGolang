package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
	ErrInvalidPassword    = errors.New("password invalid")
)

func (a Accounts) NewLogin(u account.Login) (string, error) {
	listAccounts, _ := a.store.ListAll()

	var result string
	var err error

	for _, acc := range listAccounts {
		if acc.Cpf == u.Cpf {
			err = account.VerifyPasswordHash(acc.Secret, u.Secret)
			if err != nil {
				err = ErrInvalidPassword
			}
			result = acc.Id
		}
	}
	return result, err
}
