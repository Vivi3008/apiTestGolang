package account

import (
	"errors"
	"time"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/google/uuid"
)

var (
	ErrInvalidValue = errors.New("name, cpf and secret not be empty")
	ErrCpfCaracters = errors.New("cpf must have 11 caracters")
)

type AccountId string

type Account struct {
	Id        string
	Name      string
	Cpf       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(person Account) (Account, error) {
	if person.Name == "" || person.Cpf == "" || person.Secret == "" {
		return Account{}, ErrInvalidValue
	}

	if len(person.Cpf) != 11 {
		return Account{}, ErrCpfCaracters
	}

	hashSecret, err := commom.GenerateHashPassword(person.Secret)

	if err != nil {
		return Account{}, err
	}

	return Account{
		Id:        uuid.New().String(),
		Name:      person.Name,
		Cpf:       person.Cpf,
		Secret:    hashSecret,
		Balance:   person.Balance,
		CreatedAt: time.Now(),
	}, nil
}
