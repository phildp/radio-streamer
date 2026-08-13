package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultPath(t *testing.T) {
	path, err := DefaultPath()
	if err != nil {
		t.Fatalf("DefaultPath() error = %v", err)
	}

	wantSuffix := filepath.Join(".config", "radio", "stations.yml")
	if !strings.HasSuffix(path, wantSuffix) {
		t.Fatalf("DefaultPath() = %q, want suffix %q", path, wantSuffix)
	}
}

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "stations.yml")
	content := `stations:
  test:
    title: Test FM
    filename: http://example.com/stream
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	station, ok := cfg.Stations["test"]
	if !ok {
		t.Fatal("expected station \"test\" in config")
	}
	if station.Title != "Test FM" {
		t.Fatalf("station title = %q, want %q", station.Title, "Test FM")
	}
	if station.Filename != "http://example.com/stream" {
		t.Fatalf("station filename = %q, want %q", station.Filename, "http://example.com/stream")
	}
}

func TestLoadMissingFile(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "missing.yml"))
	if err == nil {
		t.Fatal("Load() expected error for missing file")
	}
}
