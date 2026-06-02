package cmd

import (
	"bytes"
	"testing"
)

// withTempConfig redirects the config dir to an isolated temp directory so
// tests never touch the real store and don't interfere with each other.
func withTempConfig(t *testing.T) {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
}

// run executes the CLI with the given args and returns combined stdout+stderr.
func run(t *testing.T, args ...string) (string, error) {
	t.Helper()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	t.Cleanup(func() { rootCmd.SetArgs(nil) })
	err := rootCmd.Execute()
	return buf.String(), err
}
