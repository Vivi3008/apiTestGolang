package domain

import "errors"

var (
	ErrEmptyData = errors.New("Cpf and secret should not be empty")
)

type Login struct {
	Cpf    int64
	Secret string
}

type JwtLogin string

func SignUp(jwt string) (JwtLogin, error) {
	return JwtLogin(jwt), nil
}
