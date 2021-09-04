package domain

import (
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	secretHash, _ := HashPassword(person.Secret)

	return Account{
		Id:        uuid.New().String(),
		Name:      person.Name,
		Cpf:       person.Cpf,
		Secret:    string(secretHash),
		Balance:   person.Balance,
		createdAt: time.Now(),
	}, nil
}

func HashPassword(password string) ([]byte, error) {
	cost := bcrypt.DefaultCost
	secretHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return []byte(""), err
	}

	return []byte(hex.EncodeToString(secretHash)), nil
}
