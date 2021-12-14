package domain

import "errors"

var (
	ErrEmptyData = errors.New("cpf and secret should not be empty")
)

type Login struct {
	Cpf    int
	Secret string
}
