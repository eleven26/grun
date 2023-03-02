package cmd

import (
	"os"
	"os/exec"

	"github.com/eleven26/grun/console"

	"github.com/eleven26/grun/core"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "运行命令",
	Run:   runCmd,
	Args:  cobra.ExactArgs(1),
}

func runCmd(cmd *cobra.Command, args []string) {
	id := cast.ToInt(args[0])

	cmds, err := List()
	if err != nil {
		panic(err)
	}

	var found bool
	for _, c := range cmds {
		if c.Id == id {
			found = true
			execute(c)
			break
		}
	}

	if !found {
		console.Error("command not found")
	}
}

func execute(command core.Command) {
	cmd := exec.Command("sh", "-c", command.Command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
