package domain

import (
	"errors"
	"github.com/google/uuid"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

var (
	ErrInvalidValue = errors.New("name, cpf and secret not be empty")
)

type Account struct {
	Id   string
	Name   string
	Cpf     int64
	Secret    string
	Balance   float64
	createdAt time.Time
}

func NewAccount(person Account) (Account, error) {
	if person.Name == "" || person.Cpf ==0 || person.Secret =="" {
		return Account{}, ErrInvalidValue
	}

	passwordHasher := sha1.New()
	passwordHasher.Write([]byte(person.Secret))
	sha := passwordHasher.Sum(nil)

	shaStr := hex.EncodeToString(sha)

	return Account{
		Id:        uuid.New().String(),
		Name:      person.Name,
		Cpf:       person.Cpf,
		Secret:    shaStr,
		Balance:   person.Balance,
		createdAt: time.Now(),
	}, nil
}
