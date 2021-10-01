package usecases

import "github.com/Vivi3008/apiTestGolang/store"

type Accounts struct {
	store store.AccountStore
}

func CreateNewAccount(store store.AccountStore) Accounts {
	return Accounts{
		store: store,
	}
}
