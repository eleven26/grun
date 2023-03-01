package cmd

import (
	"github.com/eleven26/grun/core"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新命令",
	Run:   runUpdate,
}

func init() {
	_ = UpdateCmd.MarkFlagRequired("id")
	_ = UpdateCmd.MarkFlagRequired("command")
}

func runUpdate(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	name, _ := cmd.Flags().GetString("name")
	command, _ := cmd.Flags().GetString("command")
	description, _ := cmd.Flags().GetString("description")

	err := Update(cast.ToInt(id), core.Command{
		Name:        name,
		Command:     command,
		Description: description,
	})
	if err != nil {
		panic(err)
	}
}
