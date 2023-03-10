package store

import (
	"encoding/json"
	"os"
	"reflect"

	fs "github.com/eleven26/go-filesystem"
	"github.com/eleven26/grun/core"
	"github.com/pkg/errors"
)

var _ core.Store = &file{}

type file struct {
	filepath string
}

func (f *file) Get(id int) (*core.Command, error) {
	commands, err := f.commands()
	if err != nil {
		return nil, err
	}

	for _, command := range commands {
		if command.Id == id {
			return &command, nil
		}
	}

	return nil, errors.New("command not found")
}

func NewFileStore(filepath string) core.Store {
	return &file{
		filepath: filepath,
	}
}

func (f *file) Add(command core.Command) (*core.Command, error) {
	commands, err := f.commands()
	if err != nil {
		return nil, err
	}

	command.Id = f.nextId(commands)
	commands = append(commands, command)

	err = f.save(commands)
	if err != nil {
		return nil, err
	}

	return &command, nil
}

func (f *file) nextId(commands []core.Command) int {
	var maxId int

	for _, command := range commands {
		if command.Id > maxId {
			maxId = command.Id
		}
	}

	return maxId + 1
}

func (f *file) commands() ([]core.Command, error) {
	var commands []core.Command

	exists, err := fs.Exists(f.filepath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return commands, errors.Wrap(err, "check file exists failed")
	}
	if exists {
		content := fs.MustGet(f.filepath)
		err := json.Unmarshal(content, &commands)
		if err != nil {
			return commands, err
		}
	}

	return commands, nil
}

func (f *file) Remove(id int) error {
	commands, err := f.commands()
	if err != nil {
		return err
	}

	var result []core.Command
	for _, command := range commands {
		if command.Id == id {
			continue
		}
		result = append(result, command)
	}

	return f.save(result)
}

func (f *file) save(commands []core.Command) error {
	data, err := json.Marshal(commands)
	if err != nil {
		return err
	}

	return fs.Put(f.filepath, data)
}

func (f *file) Update(id int, command core.Command) error {
	commands, err := f.commands()
	if err != nil {
		return err
	}

	var result []core.Command
	var exists bool
	for _, c := range commands {
		if c.Id == id {
			c = f.fillNonEmpty(c, command)
			c.Id = id
			exists = true
		}
		result = append(result, c)
	}

	if !exists {
		return errors.New("command not found")
	}

	return f.save(result)
}

// ???????????????????????????
func (f *file) fillNonEmpty(old, new core.Command) core.Command {
	fieldCount := reflect.TypeOf(old).NumField()
	for i := 0; i < fieldCount; i++ {
		oldField := reflect.ValueOf(&old).Elem().Field(i)
		newField := reflect.ValueOf(&new).Elem().Field(i)

		if !newField.IsZero() {
			oldField.Set(newField)
		}
	}

	return old
}

func (f *file) List() ([]core.Command, error) {
	return f.commands()
}
