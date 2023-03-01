package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/eleven26/grun/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "grun",
		Short: "grun",
		Long:  `使用 "-h" 参数查看所有子命令`,
	}

	rootCmd.PersistentFlags().StringP("id", "i", "", "id")
	rootCmd.PersistentFlags().StringP("name", "n", "", "名称")
	rootCmd.PersistentFlags().StringP("command", "c", "", "命令")
	rootCmd.PersistentFlags().StringP("description", "d", "", "描述")

	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.UpdateCmd)
	rootCmd.AddCommand(cmd.DeleteCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	rootCmd.AddCommand(cmd.RunCmd)

	// 初始化
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Init(filepath.Join(dirname, ".grun.json"))

	if err = rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
