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
	id         string
	name       string
	cpf        int64
	secret     string
	balance   float64
	createdAt time.Time
}

func NewAccount(person Account) (Account, error) {
	if person.name == "" || person.cpf==0 || person.secret =="" {
		return Account{}, ErrInvalidValue
	}

	return Account{
		id:        uuid.New().String(),
		name:      person.name,
		cpf:       person.cpf,
		secret:    person.secret,
		balance:   person.balance,
		createdAt: time.Now(),
	}, nil
}
