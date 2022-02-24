package transfers

import (
	"io/ioutil"
	"os"

	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var sourceTest = "transfers_test.json"

func CreateTransfersFileTest() error {
	err := store.StoreFile(store.TransfersTest, sourceTest)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDataTransfersTest() error {
	err := ioutil.WriteFile(sourceTest, []byte{}, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
