package cmd

import (
	"github.com/eleven26/grun/core"
	store2 "github.com/eleven26/grun/store"
)

var store core.Store

func Init(file string) {
	store = store2.NewFileStore(file)
}

func Store(command core.Command) error {
	return store.Store(command)
}

func List() ([]core.Command, error) {
	return store.List()
}

func Delete(id int) error {
	return store.Remove(id)
}

func Update(id int, command core.Command) error {
	return store.Update(id, command)
}
