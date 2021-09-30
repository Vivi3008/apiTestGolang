package store

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
)

type Login struct {
	Cpf    int64
	Secret string
}

func (a AccountStore) NewLogin(u Login) (string, error) {
	listAll, _ := a.ListAll()

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
