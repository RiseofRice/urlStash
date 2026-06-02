package cmd

import (
	"strings"
	"testing"
)

func TestList_Empty(t *testing.T) {
	withTempConfig(t)
	out, err := run(t, "list")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "No saved Urls") {
		t.Errorf("expected empty-store message, got: %q", out)
	}
}

func TestList_ShowsEntries(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "go", "https://go.dev")
	run(t, "add", "gh", "https://github.com")

	out, err := run(t, "list")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	for _, want := range []string{"go", "gh", "https://go.dev", "https://github.com", "LABEL", "URL"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in list output, got: %q", want, out)
		}
	}
}

func TestList_AliasLs(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "ls")
	if err != nil {
		t.Fatalf("'ls' alias failed: %v", err)
	}
}

func TestList_ColumnAlignment(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "a", "https://short.io")
	run(t, "add", "longerlabel", "https://longer-url.example.com")

	out, err := run(t, "list")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Both URLs should appear on their own lines
	lines := strings.Split(out, "\n")
	urlLineCount := 0
	for _, line := range lines {
		if strings.Contains(line, "https://") {
			urlLineCount++
		}
	}
	if urlLineCount != 2 {
		t.Errorf("expected 2 URL lines, got %d in:\n%s", urlLineCount, out)
	}
}
