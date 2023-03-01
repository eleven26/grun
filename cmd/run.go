package cmd

import (
	"os"
	"os/exec"

	"github.com/eleven26/grun/core"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "运行命令",
	Run:   runCmd,
}

func init() {
	_ = UpdateCmd.MarkFlagRequired("id")
}

func runCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		panic("invalid arguments")
	}
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
		panic("command not found")
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
