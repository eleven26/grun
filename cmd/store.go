package cmd

import (
	"github.com/eleven26/grun/core"
	"github.com/spf13/cobra"
)

var StoreCmd = &cobra.Command{
	Use:   "store",
	Short: "新建新的命令",
	Run:   runStore,
}

func init() {
	_ = StoreCmd.MarkFlagRequired("name")
	_ = StoreCmd.MarkFlagRequired("command")
}

func runStore(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	c, _ := cmd.Flags().GetString("command")
	description, _ := cmd.Flags().GetString("description")

	command := core.Command{
		Name:        name,
		Command:     c,
		Description: description,
	}

	if err := Store(command); err != nil {
		panic(err)
	}
}
