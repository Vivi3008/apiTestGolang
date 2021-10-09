package usecases

import (
	"errors"
	"log"

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
		if account.Cpf == u.Cpf {
			log.Printf("cpf: ", u.Cpf)
			log.Printf("cpf: ", account)
			err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(u.Secret))
			result = domain.AccountId(account.Id)
		} else {
			return "", ErrInvalidCredentials
		}
	}
	return result, err
}
