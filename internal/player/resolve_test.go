package player

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveSourceHTTP(t *testing.T) {
	url := "https://example.com/stream"
	got, err := resolveSource(url)
	if err != nil {
		t.Fatalf("resolveSource() error = %v", err)
	}
	if got != url {
		t.Fatalf("resolveSource() = %q, want %q", got, url)
	}
}

func TestResolveSourcePLS(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "station.pls")
	content := "[playlist]\nNumberOfEntries=1\nFile1=http://example.com/live\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := resolveSource(path)
	if err != nil {
		t.Fatalf("resolveSource() error = %v", err)
	}
	if got != "http://example.com/live" {
		t.Fatalf("resolveSource() = %q, want %q", got, "http://example.com/live")
	}
}

func TestResolveSourceM3U(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "station.m3u")
	content := "#EXTM3U\n#EXTINF:-1,Station\nhttps://example.com/live\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := resolveSource(path)
	if err != nil {
		t.Fatalf("resolveSource() error = %v", err)
	}
	if got != "https://example.com/live" {
		t.Fatalf("resolveSource() = %q, want %q", got, "https://example.com/live")
	}
}

func TestResolveSourceM3U8Unsupported(t *testing.T) {
	_, err := resolveSource("https://example.com/playlist.m3u8")
	if err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("resolveSource() error = %v, want HLS unsupported error", err)
	}
}
