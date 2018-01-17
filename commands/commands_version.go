package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		Long:  "show version",
		Run:   versionCommand,
	}
	version = "0.0.1"
)

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Println(version)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
