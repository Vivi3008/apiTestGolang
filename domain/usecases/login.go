package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
)

func (a Accounts) NewLogin(u domain.Login) (string, error) {
	listAll, _ := a.store.ListAll()

	var result string
	var err error

	for _, account := range listAll {
		if account.Cpf == u.Cpf {
			err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(u.Secret))
			result = account.Id
		} else {
			return "", ErrInvalidCredentials
		}
	}
	return result, err
}
