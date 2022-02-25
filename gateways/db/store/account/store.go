package account

import "github.com/Vivi3008/apiTestGolang/domain/entities/account"

var (
	source = "account.json"
)

type AccountStore struct {
	accStore []account.Account
	Src      string
}

func NewAccountStore() AccountStore {
	as := make([]account.Account, 0)

	return AccountStore{
		accStore: as,
		Src:      source,
	}
}
