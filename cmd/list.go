package cmd

import (
	"fmt"
	"strings"

	"github.com/RiseofRice/urlStash/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "lists all Labels with their Urls",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := store.Load()
		if err != nil {
			return err
		}

		if len(s.Entries) == 0 {
			fmt.Println("No saved Urls. Use url add <label> <url>")
			return nil
		}

		maxLen := 0
		for _, e := range s.Entries {
			if len(e.Label) > maxLen {
				maxLen = len(e.Label)
			}
		}

		header := color.New(color.FgHiYellow, color.Bold).SprintFunc()
		separator := color.New(color.FgHiBlack).SprintFunc()
		label := color.New(color.FgHiMagenta).SprintFunc()
		url := color.New(color.FgHiCyan).SprintFunc()

		fmt.Printf("\n %s\n", header(fmt.Sprintf("%-*s URL", maxLen, "LABEL")))
		fmt.Printf(" %s\n", separator(strings.Repeat("-", maxLen+50)))
		for _, e := range s.Entries {
			padding := strings.Repeat(" ", maxLen-len(e.Label))
			fmt.Printf(" %s%s %s\n", label(e.Label), padding, url(e.URL))
		}
		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
