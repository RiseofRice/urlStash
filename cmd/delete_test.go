package cmd

import (
	"strings"
	"testing"
)

func TestDelete_Success(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "tmp", "https://tmp.example.com")

	out, err := run(t, "delete", "tmp")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(out, "Deleted:") {
		t.Errorf("expected 'Deleted:' in output, got: %q", out)
	}
}

func TestDelete_RemovedFromStore(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "tmp", "https://tmp.example.com")
	run(t, "delete", "tmp")

	_, err := run(t, "delete", "tmp")
	if err == nil {
		t.Fatal("expected error after deleting already-deleted label")
	}
}

func TestDelete_NotFound(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "delete", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent label, got nil")
	}
}

func TestDelete_NoArgs(t *testing.T) {
	withTempConfig(t)
	_, err := run(t, "delete")
	if err == nil {
		t.Fatal("expected error when no label given, got nil")
	}
}

func TestDelete_AliasRm(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "tmp", "https://tmp.example.com")
	_, err := run(t, "rm", "tmp")
	if err != nil {
		t.Fatalf("'rm' alias failed: %v", err)
	}
}

func TestDelete_AliasDel(t *testing.T) {
	withTempConfig(t)
	run(t, "add", "tmp", "https://tmp.example.com")
	_, err := run(t, "del", "tmp")
	if err != nil {
		t.Fatalf("'del' alias failed: %v", err)
	}
}
