package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url",
	Short: "urlstash - Your personal URL Stash",
	Long: `url stash saves Urls with a short Label.
	Examples:
		url add rryt https://www.youtube.com/watch?v=dQw4w9WgXcQ
		url open rryt
		url list`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
