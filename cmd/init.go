package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: init called, do stuff")
		// go mod init?
		// go install stuff
		// create "content" with example stuff
		// create "main.go" with some custom routes
		// gen + static / ssr in main.go file
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
