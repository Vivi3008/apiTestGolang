package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
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
	CreatedAt time.Time
}

type AccountId string

func NewAccount(person Account) (Account, error) {
	if person.Name == "" || person.Cpf == 0 || person.Secret == "" {
		return Account{}, ErrInvalidValue
	}

	hashSecret, err := bcrypt.GenerateFromPassword([]byte(person.Secret), 14)

	return Account{
		Id:        uuid.New().String(),
		Name:      person.Name,
		Cpf:       person.Cpf,
		Secret:    string(hashSecret),
		Balance:   person.Balance,
		CreatedAt: time.Now(),
	}, err
}
