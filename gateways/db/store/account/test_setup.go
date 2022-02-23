package account

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

var AccountsTest = []account.Account{
	{
		Id:        uuid.NewString(),
		Name:      "Teste 1",
		Cpf:       "77845100032",
		Secret:    "dafd33255",
		Balance:   250000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste 2",
		Cpf:       "55985633301",
		Secret:    "4f5ds4af54",
		Balance:   260000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste 3",
		Cpf:       "85665232145",
		Secret:    "fadsfdsaf",
		Balance:   360000,
		CreatedAt: time.Now(),
	},
}

var sourceTest = "account_test.json"

func CreateAccountsInFile() error {
	var err error

	var accounts = make([]account.Account, 0)

	accounts = append(accounts, AccountsTest...)

	err = store.StoreFile(accounts, sourceTest)

	if err != nil {
		return err
	}
	return nil
}

func DeleteDataTests() error {
	err := ioutil.WriteFile(sourceTest, []byte{}, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
