package cmd

import (
	"fmt"

	"github.com/RiseofRice/urlStash/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <label>",
	Aliases: []string{"rm", "del"},
	Short:   "Deletes the Url to the given label",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		label := args[0]

		s, err := store.Load()
		if err != nil {
			return err
		}

		if err := s.Delete(label); err != nil {
			return err
		}

		if err := store.Save(s); err != nil {
			return err
		}

		success := color.New(color.Attribute(148), color.Bold).SprintFunc()
		labelColor := color.New(color.FgHiMagenta).SprintFunc()
		fmt.Printf("%s %s\n", success("Deleted:"), labelColor(label))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
