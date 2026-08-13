package player

import (
	"context"
	"fmt"
	"io"
	"math"
	"path/filepath"
	"strings"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/flac"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/gopxl/beep/v2/wav"
)

func Play(ctx context.Context, source string, volume float64) error {
	resolved, err := resolveSource(source)
	if err != nil {
		return err
	}

	rc, err := openStream(resolved)
	if err != nil {
		return err
	}

	streamer, format, err := decode(resolved, rc)
	if err != nil {
		return err
	}
	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		streamer.Close()
		return fmt.Errorf("init speaker: %w", err)
	}

	stop := func() {
		streamer.Close()
		speaker.Close()
	}
	defer stop()

	done := make(chan struct{})
	speaker.Play(beep.Seq(withVolume(streamer, volume), beep.Callback(func() {
		close(done)
	})))

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func decode(name string, rc io.ReadCloser) (beep.StreamSeekCloser, beep.Format, error) {
	ext := strings.ToLower(filepath.Ext(strings.Split(name, "?")[0]))
	switch ext {
	case ".wav":
		return wav.Decode(rc)
	case ".flac":
		return flac.Decode(rc)
	case ".ogg":
		return vorbis.Decode(rc)
	case ".mp3":
		return mp3.Decode(rc)
	default:
		streamer, format, err := mp3.Decode(rc)
		if err == nil {
			return streamer, format, nil
		}
		rc.Close()
		return nil, beep.Format{}, fmt.Errorf("decode audio (tried mp3): %w", err)
	}
}

func withVolume(streamer beep.Streamer, volume float64) beep.Streamer {
	if volume <= 0 {
		return &effects.Volume{Streamer: streamer, Silent: true}
	}
	return &effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   math.Log2(volume),
	}
}
