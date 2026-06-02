package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/RiseofRice/urlStash/store"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var copyToClipboardFn = defaultCopyToClipboard

func defaultCopyToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		if os.Getenv("WAYLAND_DISPLAY") != "" {
			cmd = exec.Command("wl-copy")
		} else {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		}
	}
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

var copyCmd = &cobra.Command{
	Use:     "copy <label>",
	Short:   "copies the url of the given Label",
	Aliases: []string{"cp", "yank"},
	Args:    cobra.ExactArgs(1),
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

		if err := copyToClipboardFn(entry.URL); err != nil {
			return fmt.Errorf("copy failed: %w", err)
		}

		success := color.New(color.Attribute(148), color.Bold).SprintFunc()
		urlColor := color.New(color.FgHiCyan).SprintFunc()
		fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", success("Copied:"), urlColor(entry.URL))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
