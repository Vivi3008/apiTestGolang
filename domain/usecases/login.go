package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
	ErrInvalidPassword    = errors.New("Password invalid")
)

func (a Accounts) NewLogin(u domain.Login) (domain.AccountId, error) {
	listAll, _ := a.store.ListAll()

	var result domain.AccountId
	var err error

	for _, account := range listAll {
		if account.Cpf == u.Cpf {
			err = domain.VerifyPasswordHash(account.Secret, u.Secret)
			if err != nil {
				err = ErrInvalidPassword
			}
			result = domain.AccountId(account.Id)
		}
	}
	return result, err
}
