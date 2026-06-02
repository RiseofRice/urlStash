package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func TestCopy_NotFound(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "copy", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent label, got nil")
	}
}

func TestCopy_NoArgs(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "copy")
	if err == nil {
		t.Fatal("expected error when no label given, got nil")
	}
}

func TestCopy_Success(t *testing.T) {
	withTempConfig(t)

	var copied string
	orig := copyToClipboardFn
	copyToClipboardFn = func(text string) error { copied = text; return nil }
	t.Cleanup(func() { copyToClipboardFn = orig })

	run(t, "add", "test", "https://example.com")
	out, err := run(t, "copy", "test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if copied != "https://example.com" {
		t.Errorf("expected clipboard to contain https://example.com, got %q", copied)
	}
	if !strings.Contains(out, "Copied:") {
		t.Errorf("expected 'Copied:' in output, got: %q", out)
	}
}

func TestCopy_ClipboardError(t *testing.T) {
	withTempConfig(t)

	orig := copyToClipboardFn
	copyToClipboardFn = func(text string) error { return fmt.Errorf("clipboard unavailable") }
	t.Cleanup(func() { copyToClipboardFn = orig })

	run(t, "add", "test", "https://example.com")
	_, err := run(t, "copy", "test")
	if err == nil {
		t.Fatal("expected error when clipboard fails, got nil")
	}
}

func TestCopy_AliasCp(t *testing.T) {
	withTempConfig(t)

	orig := copyToClipboardFn
	copyToClipboardFn = func(text string) error { return nil }
	t.Cleanup(func() { copyToClipboardFn = orig })

	run(t, "add", "test", "https://example.com")
	if _, err := run(t, "cp", "test"); err != nil {
		t.Fatalf("'cp' alias failed: %v", err)
	}
}

func TestCopy_AliasYank(t *testing.T) {
	withTempConfig(t)

	orig := copyToClipboardFn
	copyToClipboardFn = func(text string) error { return nil }
	t.Cleanup(func() { copyToClipboardFn = orig })

	run(t, "add", "test", "https://example.com")
	if _, err := run(t, "yank", "test"); err != nil {
		t.Fatalf("'yank' alias failed: %v", err)
	}
}
