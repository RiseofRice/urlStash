package cmd

import (
	"strings"
	"testing"
)

func TestVersion_Output(t *testing.T) {
	out, err := run(t, "version")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "urlStash") {
		t.Errorf("expected 'urlStash' in output, got: %q", out)
	}
	if !strings.Contains(out, version) {
		t.Errorf("expected version %q in output, got: %q", version, out)
	}
}
