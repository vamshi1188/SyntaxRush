package core

import (
	"io"
	"math"
	"time"

	"github.com/hajimehoshi/oto/v2"
)

// AudioManager handles sound effects for the application
type AudioManager struct {
	context    *oto.Context
	sampleRate int
	channels   int
	bitDepth   int
}

// NewAudioManager creates a new audio manager
func NewAudioManager() (*AudioManager, error) {
	// Audio configuration
	sampleRate := 44100
	channels := 2
	bitDepth := 2 // 16-bit

	// Initialize oto context with v2 API
	ctx, ready, err := oto.NewContext(sampleRate, channels, oto.FormatSignedInt16LE)
	if err != nil {
		return nil, err
	}

	// Wait for the audio context to be ready
	<-ready

	return &AudioManager{
		context:    ctx,
		sampleRate: sampleRate,
		channels:   channels,
		bitDepth:   bitDepth,
	}, nil
}

// generateBeepWave generates a beep sound wave
func (am *AudioManager) generateBeepWave(frequency float64, duration time.Duration) []byte {
	samples := int(float64(am.sampleRate) * duration.Seconds())
	data := make([]byte, samples*am.channels*am.bitDepth)

	for i := 0; i < samples; i++ {
		// Generate sine wave
		t := float64(i) / float64(am.sampleRate)
		sample := math.Sin(2 * math.Pi * frequency * t)

		// Apply envelope to avoid clicks (fade in/out)
		envelope := 1.0
		fadeTime := 0.01 // 10ms fade
		fadeSamples := int(fadeTime * float64(am.sampleRate))

		if i < fadeSamples {
			envelope = float64(i) / float64(fadeSamples)
		} else if i > samples-fadeSamples {
			envelope = float64(samples-i) / float64(fadeSamples)
		}

		sample *= envelope * 0.3 // Reduce volume to 30%

		// Convert to 16-bit PCM
		pcm := int16(sample * 32767)

		// Write stereo samples
		for ch := 0; ch < am.channels; ch++ {
			offset := (i*am.channels + ch) * am.bitDepth
			data[offset] = byte(pcm & 0xff)
			data[offset+1] = byte((pcm >> 8) & 0xff)
		}
	}

	return data
}

// PlayErrorBeep plays a beep sound for typing mistakes
func (am *AudioManager) PlayErrorBeep() {
	if am.context == nil {
		return
	}

	// Generate a short beep (800Hz for 100ms)
	beepData := am.generateBeepWave(800, 100*time.Millisecond)

	// Create a player and play the sound
	go func() {
		player := am.context.NewPlayer(io.NopCloser(&beepReader{data: beepData}))
		defer player.Close()
		player.Play()

		// Wait for the sound to finish
		time.Sleep(100 * time.Millisecond)
	}()
}

// PlaySuccessSound plays a pleasant sound for line completion
func (am *AudioManager) PlaySuccessSound() {
	if am.context == nil {
		return
	}

	// Generate a pleasant two-tone sound
	go func() {
		// First tone (C note - 523Hz)
		tone1 := am.generateBeepWave(523, 80*time.Millisecond)

		// Second tone (E note - 659Hz)
		tone2 := am.generateBeepWave(659, 80*time.Millisecond)

		// Combine tones with a small gap
		combined := make([]byte, len(tone1)+len(tone2))
		copy(combined[:len(tone1)], tone1)
		copy(combined[len(tone1):], tone2)

		player := am.context.NewPlayer(io.NopCloser(&beepReader{data: combined}))
		defer player.Close()
		player.Play()

		// Wait for the sound to finish
		time.Sleep(200 * time.Millisecond)
	}()
}

// Close closes the audio context
func (am *AudioManager) Close() {
	if am.context != nil {
		am.context.Suspend()
	}
}

// beepReader implements io.Reader for audio data
type beepReader struct {
	data   []byte
	offset int
}

func (br *beepReader) Read(p []byte) (n int, err error) {
	if br.offset >= len(br.data) {
		return 0, io.EOF
	}

	n = copy(p, br.data[br.offset:])
	br.offset += n
	return n, nil
}

func (br *beepReader) Close() error {
	return nil
}
