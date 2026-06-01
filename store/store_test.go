package store

import (
	"os"
	"testing"
)

func withTempHome(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("HOME", dir)
}

func TestLoad_NoFile(t *testing.T) {
	withTempHome(t)

	s, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(s.Entries) != 0 {
		t.Fatalf("expected empty entries, got %d", len(s.Entries))
	}
}

func TestSaveAndLoad(t *testing.T) {
	withTempHome(t)

	original := &Store{
		Entries: []Entry{
			{Label: "go", URL: "https://go.dev"},
			{Label: "gh", URL: "https://github.com"},
		},
	}

	if err := Save(original); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if len(loaded.Entries) != len(original.Entries) {
		t.Fatalf("expected %d entries, got %d", len(original.Entries), len(loaded.Entries))
	}
	for i, e := range original.Entries {
		if loaded.Entries[i] != e {
			t.Errorf("entry %d: expected %+v, got %+v", i, e, loaded.Entries[i])
		}
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	withTempHome(t)

	path, err := storePath()
	if err != nil {
		t.Fatalf("storePath failed: %v", err)
	}
	if err := os.WriteFile(path, []byte("not json"), 0644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	_, err = Load()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestSave_Overwrite(t *testing.T) {
	withTempHome(t)

	first := &Store{Entries: []Entry{{Label: "old", URL: "https://old.example"}}}
	if err := Save(first); err != nil {
		t.Fatalf("first Save failed: %v", err)
	}

	second := &Store{Entries: []Entry{{Label: "new", URL: "https://new.example"}}}
	if err := Save(second); err != nil {
		t.Fatalf("second Save failed: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if len(loaded.Entries) != 1 || loaded.Entries[0].Label != "new" {
		t.Errorf("expected only new entry, got %+v", loaded.Entries)
	}
}

func TestAdd(t *testing.T) {
	s := &Store{}

	if err := s.Add("go", "https://go.dev"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(s.Entries) != 1 || s.Entries[0].Label != "go" {
		t.Fatalf("expected entry 'go', got %+v", s.Entries)
	}
}

func TestAdd_DuplicateLabel(t *testing.T) {
	s := &Store{}
	_ = s.Add("go", "https://go.dev")

	err := s.Add("go", "https://other.dev")
	if err == nil {
		t.Fatal("expected error for duplicate label, got nil")
	}
	if len(s.Entries) != 1 {
		t.Errorf("expected 1 entry, got %d", len(s.Entries))
	}
}

func TestGet(t *testing.T) {
	s := &Store{Entries: []Entry{
		{Label: "go", URL: "https://go.dev"},
		{Label: "gh", URL: "https://github.com"},
	}}

	e, err := s.Get("gh")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if e.URL != "https://github.com" {
		t.Errorf("expected https://github.com, got %s", e.URL)
	}
}

func TestGet_NotFound(t *testing.T) {
	s := &Store{}

	_, err := s.Get("missing")
	if err == nil {
		t.Fatal("expected error for missing label, got nil")
	}
}

func TestDelete(t *testing.T) {
	s := &Store{Entries: []Entry{
		{Label: "a", URL: "https://a.dev"},
		{Label: "b", URL: "https://b.dev"},
		{Label: "c", URL: "https://c.dev"},
	}}

	if err := s.Delete("b"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(s.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(s.Entries))
	}
	if s.Entries[0].Label != "a" || s.Entries[1].Label != "c" {
		t.Errorf("unexpected entries after delete: %+v", s.Entries)
	}
}

func TestDelete_NotFound(t *testing.T) {
	s := &Store{Entries: []Entry{{Label: "a", URL: "https://a.dev"}}}

	err := s.Delete("missing")
	if err == nil {
		t.Fatal("expected error for missing label, got nil")
	}
	if len(s.Entries) != 1 {
		t.Errorf("expected 1 entry unchanged, got %d", len(s.Entries))
	}
}

func TestDelete_LastEntry(t *testing.T) {
	s := &Store{Entries: []Entry{{Label: "only", URL: "https://only.dev"}}}

	if err := s.Delete("only"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(s.Entries) != 0 {
		t.Errorf("expected empty entries, got %+v", s.Entries)
	}
}
