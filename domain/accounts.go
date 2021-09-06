package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidValue = errors.New("name, cpf and secret not be empty")
)

type Account struct {
	Id        string
	Name      string
	Cpf       int64
	Secret    string
	Balance   float64
	createdAt time.Time
}

func NewAccount(person Account) (Account, error) {
	if person.Name == "" || person.Cpf == 0 || person.Secret == "" {
		return Account{}, ErrInvalidValue
	}

	return Account{
		Id:        uuid.New().String(),
		Name:      person.Name,
		Cpf:       person.Cpf,
		Secret:    person.Secret,
		Balance:   person.Balance,
		createdAt: time.Now(),
	}, nil
}
