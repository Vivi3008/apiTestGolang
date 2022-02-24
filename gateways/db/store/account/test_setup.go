package account

import (
	"io/ioutil"
	"os"

	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var sourceTest = "account_test.json"

func CreateAccountsInFile() error {
	err := store.StoreFile(store.AccountsTest, sourceTest)

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
