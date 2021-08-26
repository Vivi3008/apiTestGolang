package store

import (
	"errors"

	"github.com/Vivi3008/apiTestGolang/domain"
)

var (
	ErrEmptyID = errors.New("id not be empty")
)

type AccountStore struct {
	accStore map[string]domain.Account
}

func NewAccountStore() AccountStore {
	as := make(map[string]domain.Account)

	return AccountStore{
		accStore: as,
	}
}