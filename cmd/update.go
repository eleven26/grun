package cmd

import (
	"github.com/eleven26/grun/console"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新命令",
	Run:   runUpdate,
	Args:  cobra.ExactArgs(1),
}

func runUpdate(cmd *cobra.Command, args []string) {
	id := cast.ToInt(args[0])

	old, err := Get(cast.ToInt(id))
	if err != nil {
		panic(err)
	}

	p := prompter{}
	c := p.askForInput(map[string]string{"Name": old.Name, "Command": old.Command, "Description": old.Description})

	err = Update(cast.ToInt(id), c)
	if err != nil {
		panic(err)
	}

	console.Success("命令更新成功")
}
