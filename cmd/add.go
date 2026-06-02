package cmd

import (
	"fmt"

	"github.com/RiseofRice/urlStash/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <label> <url>",
	Short: "Saves the Url with a label",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		label := args[0]
		url := args[1]

		s, err := store.Load()
		if err != nil {
			return fmt.Errorf("Storageloading failed: %w", err)
		}

		if err := s.Add(label, url); err != nil {
			return err
		}

		if err := store.Save(s); err != nil {
			return fmt.Errorf("failed to save the storage: %w", err)
		}

		success := color.New(color.Attribute(148), color.Bold).SprintFunc()
		labelColor := color.New(color.FgHiMagenta).SprintFunc()
		urlColor := color.New(color.FgHiCyan).SprintFunc()
		fmt.Fprintf(cmd.OutOrStdout(), "%s %s -> %s\n", success("Saved:"), labelColor(label), urlColor(url))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
