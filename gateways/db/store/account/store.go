package account

import "github.com/Vivi3008/apiTestGolang/domain/entities/account"

var (
	source = "./gateways/db/store/account/account.json"
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
