package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/RiseofRice/urlStash/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <label>",
	Short: "opens the url in your Standard browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		label := args[0]

		s, err := store.Load()
		if err != nil {
			return err
		}

		entry, err := s.Get(label)
		if err != nil {
			return err
		}

		if err := openBrowser(entry.URL); err != nil {
			return fmt.Errorf("opening browser failed %w", err)
		}

		success := color.New(color.Attribute(148), color.Bold).SprintFunc()
		urlColor := color.New(color.FgHiCyan).SprintFunc()
		fmt.Printf("%s %s\n", success("Opening:"), urlColor(entry.URL))
		return nil
	},
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}

func init() {
	rootCmd.AddCommand(openCmd)
}
