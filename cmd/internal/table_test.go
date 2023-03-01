package internal

import (
	"testing"

	"github.com/eleven26/grun/core"
	"github.com/stretchr/testify/assert"
)

func TestOutputTable(t *testing.T) {
	commands := []core.Command{
		{
			Id:          1,
			Name:        "test",
			Command:     "ls -lh",
			Description: "List files in current directory.",
		},
	}

	s := OutputTable(commands)

	assert.Equal(t, ` ID   名称    命令                  描述               
---- ------ -------- ----------------------------------
 1    test   ls -lh   List files in current directory. `, s)
}
