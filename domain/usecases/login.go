package usecases

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
)

func (a Accounts) NewLogin(u domain.Login) (domain.AccountId, error) {
	listAll, _ := a.store.ListAll()

	var result domain.AccountId
	var err error

	for _, account := range listAll {
		if account.Cpf == int64(u.Cpf) {
			err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(u.Secret))
			result = domain.AccountId(account.Id)
		}
	}
	return result, err
}
