package store

import (
	"io/ioutil"
	"os"
)

func CreateDataFile(src string, data interface{}) error {
	err := StoreFile(data, src)
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
