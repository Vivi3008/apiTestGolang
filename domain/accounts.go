package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id         string
	name       string
	cpf        int64
	secret     string
	balance    float64
	created_at time.Time
}

func NewAccount(person Account) (Account, error) {
	return Account{
		id:         uuid.New().String(),
		name:       person.name,
		cpf:        person.cpf,
		secret:     person.secret,
		balance:    person.balance,
		created_at: time.Now(),
	}, nil
}
