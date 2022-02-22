package account

import "github.com/Vivi3008/apiTestGolang/domain/entities/account"

var source = "account.json"

type AccountStore struct {
	accStore map[string]account.Account
	src      string
}

func NewAccountStore() AccountStore {
	as := make(map[string]account.Account)

	return AccountStore{
		accStore: as,
		src:      source,
	}
}
