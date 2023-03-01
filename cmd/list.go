package cmd

import (
	"fmt"

	"github.com/eleven26/grun/cmd/internal"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有命令",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	cmds, err := List()
	if err != nil {
		panic(err)
	}

	fmt.Println(internal.OutputTable(cmds))
}
