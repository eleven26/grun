package store

import (
	"encoding/json"
	"testing"

	fs "github.com/eleven26/go-filesystem"
	"github.com/eleven26/grun/core"
	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	command := core.Command{
		Name:        "test",
		Command:     "ls -lh",
		Description: "List files in current directory.",
	}

	path := "store.json"
	defer fs.Delete("store.json") // nolint

	store := NewFileStore(path)
	err := store.Store(command)
	assert.Nil(t, err)

	// 读取文件内容
	data := fs.MustGet(path)
	var cmd []core.Command
	err = json.Unmarshal(data, &cmd)
	assert.Nil(t, err)

	command.Id = 1
	assert.Equal(t, command, cmd[0])
}

func TestRemove(t *testing.T) {
	path := "store.json"
	defer fs.Delete("store.json") // nolint

	store := NewFileStore(path)
	err := store.Store(core.Command{
		Name:        "foo",
		Command:     "ls",
		Description: "List files in current directory.",
	})
	assert.Nil(t, err)

	err = store.Store(core.Command{
		Name:        "bar",
		Command:     "ls",
		Description: "List files in current directory.",
	})
	assert.Nil(t, err)

	err = store.Remove(1)
	assert.Nil(t, err)

	// 读取文件内容
	data := fs.MustGet(path)
	var cmd []core.Command
	err = json.Unmarshal(data, &cmd)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(cmd))

	assert.Equal(t, 2, cmd[0].Id)
	assert.Equal(t, "bar", cmd[0].Name)
	assert.Equal(t, "ls", cmd[0].Command)
	assert.Equal(t, "List files in current directory.", cmd[0].Description)

	err = store.Remove(2)
	assert.Nil(t, err)

	// 读取文件内容
	data = fs.MustGet(path)
	err = json.Unmarshal(data, &cmd)
	assert.Nil(t, err)

	assert.Equal(t, 0, len(cmd))
}

func TestUpdate(t *testing.T) {
	path := "store.json"
	defer fs.Delete("store.json") // nolint

	store := NewFileStore(path)
	err := store.Store(core.Command{
		Name:        "foo",
		Command:     "ls",
		Description: "List files in current directory.",
	})
	assert.Nil(t, err)

	err = store.Update(2, core.Command{
		Name:        "bar",
		Command:     "ls -lh",
		Description: "List files in current directory.",
	})
	assert.NotNil(t, err)
	assert.Equal(t, "command not found", err.Error())

	command := core.Command{
		Name:        "bar(updated)",
		Command:     "ls -lh(updated)",
		Description: "",
	}
	err = store.Update(1, command)
	assert.Nil(t, err)

	// 读取文件内容
	data := fs.MustGet(path)
	var cmd []core.Command
	err = json.Unmarshal(data, &cmd)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(cmd))
	command.Id = 1
	assert.NotEqual(t, cmd[0], command)
	assert.NotEmpty(t, cmd[0].Description)
}

func TestList(t *testing.T) {
	path := "store.json"
	defer fs.Delete("store.json") // nolint

	cmds := []core.Command{
		{
			Name:        "foo",
			Command:     "ls",
			Description: "List files in current directory.",
		},
		{
			Name:        "foo1",
			Command:     "ls1",
			Description: "List files in current directory.1",
		},
	}

	store := NewFileStore(path)
	err := store.Store(cmds[0])
	assert.Nil(t, err)

	err = store.Store(cmds[1])
	assert.Nil(t, err)

	commands, err := store.List()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(commands))

	for _, command := range commands {
		assert.True(t, command.Id == 1 || command.Id == 2)
		assert.True(t, command.Name == "foo" || command.Name == "foo1")
		assert.True(t, command.Command == "ls" || command.Command == "ls1")
		assert.True(t, command.Description == "List files in current directory." || command.Description == "List files in current directory.1")
	}
}
