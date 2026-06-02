package cmd

import (
	"strings"
	"testing"
)

func TestAdd_Success(t *testing.T) {
	withTempConfig(t)
	out, err := run(t, "add", "gh", "https://github.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "Saved:") {
		t.Errorf("expected 'Saved:' in output, got: %q", out)
	}
	if !strings.Contains(out, "gh") {
		t.Errorf("expected label in output, got: %q", out)
	}
	if !strings.Contains(out, "https://github.com") {
		t.Errorf("expected URL in output, got: %q", out)
	}
}

func TestAdd_DuplicateLabel(t *testing.T) {
	withTempConfig(t)
	if _, err := run(t, "add", "gh", "https://github.com"); err != nil {
		t.Fatalf("first add failed: %v", err)
	}
	_, err := run(t, "add", "gh", "https://github.com/other")
	if err == nil {
		t.Fatal("expected error for duplicate label, got nil")
	}
}

func TestAdd_MissingURL(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "add", "only-label")
	if err == nil {
		t.Fatal("expected error when URL is missing, got nil")
	}
}

func TestAdd_NoArgs(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "add")
	if err == nil {
		t.Fatal("expected error when no args given, got nil")
	}
}

func TestAdd_PersistsToStore(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "go", "https://go.dev")

	out, err := run(t, "list")
	if err != nil {
		t.Fatalf("list after add failed: %v", err)
	}
	if !strings.Contains(out, "go") || !strings.Contains(out, "https://go.dev") {
		t.Errorf("expected saved entry in list output, got: %q", out)
	}
}
