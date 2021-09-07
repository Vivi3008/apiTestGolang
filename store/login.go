package store

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
)

type Login struct {
	Cpf    int64
	Secret string
}

func(a AccountStore) NewLogin(u Login) (string, error) {
	listAll, _ := a.ListAll()

	var result string

 for _, account := range listAll {
		if account.Cpf == u.Cpf && account.Secret == u.Secret {
			result = account.Id
		} else {
			return "", ErrInvalidCredentials
		}
	}
	return result, nil
}
