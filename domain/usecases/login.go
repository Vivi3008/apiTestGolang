package usecases

import (
	"errors"
	"fmt"
	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/Vivi3008/apiTestGolang/store"
)

var (
	ErrInvalidCredentials = errors.New("cpf, secret are invalid")
	ErrSearchedProblem    = errors.New("error to search account")
	ErrSaveLogin = errors.New("error to save person")
)

type Login struct {
	Cpf    int64
	Secret string
}

func NewLogin(u Login) (bool, error) {
	storeAcc := store.NewAccountStore()

	person := domain.Account{
		Name:    "Vanny",
		Cpf:     13323332555,
		Secret:  "dafd33255",
		Balance: 2.500,
	}

	err1 := storeAcc.StoreAccount(person)

	if err1 != nil {
		return false, error.Error()
	}

	accounts, err := storeAcc.ListAll()

	if err != nil {
		return false, ErrSearchedProblem
	}

	var result bool

 for _, account := range accounts {
		if account.Cpf == u.Cpf && account.Secret == u.Secret {
			fmt.Println("existe essa caçamba")
			result = true
		} else {
			fmt.Println("credenciais inexistentes")
			result = false
		}
	}
	return result, nil
}
