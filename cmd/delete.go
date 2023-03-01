package cmd

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除命令",
	Run:   runDelete,
}

func init() {
	_ = DeleteCmd.MarkFlagRequired("id")
}

func runDelete(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")

	if err := Delete(cast.ToInt(id)); err != nil {
		panic(err)
	}
}
