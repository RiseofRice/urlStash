package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Entry struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Store struct {
	Entries []Entry `json:"entries"`
}

func storePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".urlstash.json"), nil
}

func Load() (*Store, error) {
	path, err := storePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return &Store{Entries: []Entry{}}, nil
	} else if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s Store
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func Save(s *Store) error {
	path, err := storePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)

}

func (s *Store) Add(label, url string) error {
	// Check if label exists
	for _, e := range s.Entries {
		if e.Label == label {
			return fmt.Errorf("label '%s' exists", label)
		}
	}
	s.Entries = append(s.Entries, Entry{Label: label, URL: url})
	return nil
}

func (s *Store) Get(label string) (Entry, error) {
	for _, e := range s.Entries {
		if e.Label == label {
			return e, nil
		}
	}
	return Entry{}, fmt.Errorf("label '%s' not found", label)
}

func (s *Store) Delete(label string) error {
	for i, e := range s.Entries {

		if e.Label == label {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("label '%s' not found", label)
}
