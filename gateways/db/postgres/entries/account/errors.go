package accountdb

import "errors"

var (
	ErrCpfExists      = errors.New("this cpf already exists")
	ErrBalanceInvalid = errors.New("balance can't be less than 0")
	ErrCpfNotExists   = errors.New("this cpf does not exist")
	ErrIdNotExists    = errors.New("id does not exist")
	ErrIdExists       = errors.New("this id already exists")
)
