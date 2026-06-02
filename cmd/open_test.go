package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func TestOpen_NotFound(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "open", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent label, got nil")
	}
}

func TestOpen_NoArgs(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "open")
	if err == nil {
		t.Fatal("expected error when no label given, got nil")
	}
}

func TestOpen_Success(t *testing.T) {
	withTempConfig(t)

	var opened string
	orig := openBrowserFn
	openBrowserFn = func(url string) error { opened = url; return nil }
	t.Cleanup(func() { openBrowserFn = orig })

	run(t, "add", "test", "https://example.com")
	out, err := run(t, "open", "test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if opened != "https://example.com" {
		t.Errorf("expected browser to open https://example.com, got %q", opened)
	}
	if !strings.Contains(out, "Opening:") {
		t.Errorf("expected 'Opening:' in output, got: %q", out)
	}
}

func TestOpen_BrowserError(t *testing.T) {
	withTempConfig(t)

	orig := openBrowserFn
	openBrowserFn = func(url string) error { return fmt.Errorf("no browser available") }
	t.Cleanup(func() { openBrowserFn = orig })

	run(t, "add", "test", "https://example.com")
	_, err := run(t, "open", "test")
	if err == nil {
		t.Fatal("expected error when browser fails, got nil")
	}
}
