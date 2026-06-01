package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the actual version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("urlStash", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
