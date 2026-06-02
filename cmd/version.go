package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the actual version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "urlStash", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
