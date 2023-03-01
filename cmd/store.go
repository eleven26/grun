package cmd

import (
	"github.com/eleven26/grun/core"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "新建命令",
	Run:   runStore,
}

func init() {
	_ = AddCmd.MarkFlagRequired("name")
	_ = AddCmd.MarkFlagRequired("command")
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

	if err := Add(command); err != nil {
		panic(err)
	}
}
