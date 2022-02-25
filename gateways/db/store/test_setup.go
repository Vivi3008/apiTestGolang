package store

import (
	"io/ioutil"
	"os"
)

func CreateDataFile(src string) error {
	err := StoreFile(AccountsTest, src)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDataFile(src string) error {
	err := ioutil.WriteFile(src, []byte{}, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
