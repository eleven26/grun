package cmd

import (
	"github.com/eleven26/grun/cmd/internal"
	"github.com/eleven26/grun/console"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "新建命令",
	Run:   runStore,
}

func runStore(cmd *cobra.Command, args []string) {
	p := prompter{}
	command := p.askForInput(map[string]string{"Name": "", "Command": "", "Description": ""})

	c, err := Add(command)
	if err != nil {
		panic(err)
	}

	console.Success("命令添加成功，你可以通过 grun run " + cast.ToString(c.Id) + " 来运行它")
	console.Success(internal.OutputCmd(*c))
}
