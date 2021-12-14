package domain

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidValue = errors.New("name, cpf and secret not be empty")
	ErrCpfCaracters = errors.New("Cpf must have 11 caracters")
)

type Account struct {
	Id        string
	Name      string
	Cpf       int
	Secret    string
	Balance   int64
	CreatedAt time.Time
}

type AccountId string

func NewAccount(person Account) (Account, error) {
	if person.Name == "" || person.Cpf == 0 || person.Secret == "" {
		return Account{}, ErrInvalidValue
	}

	cpfFormat := strconv.Itoa(person.Cpf)

	if len(cpfFormat) != 11 {
		return Account{}, ErrCpfCaracters
	}

	hashSecret, err := GenerateHashPassword(person.Secret)

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
