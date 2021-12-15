package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
	ErrInvalidPassword    = errors.New("password invalid")
)

func (a Accounts) NewLogin(u domain.Login) (string, error) {
	listAccounts, _ := a.store.ListAll()

	var result string
	var err error

	for _, account := range listAccounts {
		if account.Cpf == u.Cpf {
			err = domain.VerifyPasswordHash(account.Secret, u.Secret)
			if err != nil {
				err = ErrInvalidPassword
			}
			result = account.Id
		}
	}
	return result, err
}
