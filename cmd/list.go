package cmd

import (
	"fmt"

	"github.com/eleven26/grun/console"

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

	console.Success("共有 " + fmt.Sprintf("%d", len(cmds)) + " 条命令，你可以通过 grun run [id] 来运行它")
	fmt.Println(internal.OutputTable(cmds))
}
