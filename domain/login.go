package domain

import "errors"

var (
	ErrEmptyData = errors.New("Cpf and secret should not be empty")
)

type Login struct {
	Cpf    int
	Secret string
}
