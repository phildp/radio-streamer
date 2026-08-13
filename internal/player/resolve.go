package player

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func resolveSource(source string) (string, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		if isHLS(source) {
			return "", fmt.Errorf("HLS (.m3u8) streams are not supported")
		}
		return source, nil
	}

	ext := strings.ToLower(filepath.Ext(source))
	switch ext {
	case ".pls":
		return parsePlaylistURL(source, parsePLS)
	case ".m3u":
		return parsePlaylistURL(source, parseM3U)
	case ".m3u8":
		return "", fmt.Errorf("HLS (.m3u8) streams are not supported")
	default:
		return source, nil
	}
}

func isHLS(source string) bool {
	path := strings.Split(source, "?")[0]
	return strings.HasSuffix(strings.ToLower(path), ".m3u8")
}

func parsePlaylistURL(path string, parse func(io.Reader) (string, error)) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open playlist %q: %w", path, err)
	}
	defer f.Close()

	url, err := parse(f)
	if err != nil {
		return "", fmt.Errorf("parse playlist %q: %w", path, err)
	}
	return url, nil
}

func parsePLS(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(strings.ToUpper(line), "FILE1=") {
			return strings.TrimSpace(line[6:]), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no stream URL found in playlist")
}

func parseM3U(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			return line, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no stream URL found in playlist")
}

func openStream(source string) (io.ReadCloser, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		resp, err := http.Get(source)
		if err != nil {
			return nil, fmt.Errorf("fetch stream: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("fetch stream: HTTP %s", resp.Status)
		}
		return resp.Body, nil
	}

	f, err := os.Open(source)
	if err != nil {
		return nil, fmt.Errorf("open stream: %w", err)
	}
	return f, nil
}
