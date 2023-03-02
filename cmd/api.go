package cmd

import (
	"github.com/eleven26/grun/core"
	store2 "github.com/eleven26/grun/store"
)

var store core.Store

func Init(file string) {
	store = store2.NewFileStore(file)
}

func Add(command core.Command) (*core.Command, error) {
	return store.Add(command)
}

func List() ([]core.Command, error) {
	return store.List()
}

func Get(id int) (*core.Command, error) {
	return store.Get(id)
}

func Delete(id int) error {
	return store.Remove(id)
}

func Update(id int, command core.Command) error {
	return store.Update(id, command)
}
