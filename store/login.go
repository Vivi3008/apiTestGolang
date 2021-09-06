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

func(a AccountStore) NewLogin(u Login) (bool, error) {
	listAll, _ := a.ListAll()

	var result bool

 for _, account := range listAll {
		if account.Cpf == u.Cpf && account.Secret == u.Secret {
			result = true
		} else {
			return false, ErrInvalidCredentials
		}
	}
	return result, nil
}
