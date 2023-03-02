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
	_, err := store.Add(command)
	assert.Nil(t, err)

	// 读取文件内容
	data := fs.MustGet(path)
	var cmd []core.Command
	err = json.Unmarshal(data, &cmd)
	assert.Nil(t, err)

	command.Id = 1
	assert.Equal(t, command, cmd[0])
}

func TestGet(t *testing.T) {
	command := core.Command{
		Name:        "test",
		Command:     "ls -lh",
		Description: "List files in current directory.",
	}

	path := "store.json"
	defer fs.Delete("store.json") // nolint

	store := NewFileStore(path)
	_, err := store.Add(command)
	assert.Nil(t, err)

	cmd, err := store.Get(1)
	assert.Nil(t, err)

	command.Id = 1
	assert.Equal(t, command, *cmd)
}

func TestRemove(t *testing.T) {
	path := "store.json"
	defer fs.Delete("store.json") // nolint

	store := NewFileStore(path)
	_, err := store.Add(core.Command{
		Name:        "foo",
		Command:     "ls",
		Description: "List files in current directory.",
	})
	assert.Nil(t, err)

	_, err = store.Add(core.Command{
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
	_, err := store.Add(core.Command{
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
	_, err := store.Add(cmds[0])
	assert.Nil(t, err)

	_, err = store.Add(cmds[1])
	assert.Nil(t, err)

	commands, err := store.List()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(commands))
	cmds[0].Id = 1
	cmds[1].Id = 2
	assert.Equal(t, cmds, commands)
}
