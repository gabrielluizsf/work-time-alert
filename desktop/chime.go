package desktop

import (
	"math"
	"sync"

	"github.com/gabrielluizsf/antui"
)

// Chime plays the alert sound: the same three soft beeps the web version
// falls back to. The audio stream is opened once and plays silence between
// chimes, so ringing never waits on the sound system.
type Chime struct {
	mu      sync.Mutex
	samples []float32
	pos     int
	audio   *antui.Audio
}

const chimeRate = 44100

// NewChime opens the audio device. A machine with no sound still gets a
// working Chime; Play is then silent.
func NewChime() *Chime {
	c := &Chime{samples: Beeps(chimeRate)}
	c.pos = len(c.samples)
	audio, err := antui.OpenAudio("Work Time Alert", chimeRate, c.fill)
	if err == nil {
		c.audio = audio
	}
	return c
}

// Play starts the chime from the beginning.
func (c *Chime) Play() {
	c.mu.Lock()
	c.pos = 0
	c.mu.Unlock()
}

// Close releases the audio device.
func (c *Chime) Close() {
	if c.audio != nil {
		c.audio.Close()
	}
}

// fill writes interleaved stereo frames, called from the audio thread.
func (c *Chime) fill(buf []float32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i+1 < len(buf); i += 2 {
		var v float32
		if c.pos < len(c.samples) {
			v = c.samples[c.pos]
			c.pos++
		}
		buf[i], buf[i+1] = v, v
	}
}

// Beeps renders three 880 Hz beeps half a second apart, each with a quick
// attack and an exponential fade so they do not click.
func Beeps(rate int) []float32 {
	const (
		count  = 3
		gap    = 0.5
		length = 0.42
		freq   = 880.0
		peak   = 0.6
		attack = 0.03
	)
	out := make([]float32, int(float64(rate)*(gap*(count-1)+length)))
	for b := range count {
		start := int(float64(rate) * gap * float64(b))
		n := int(float64(rate) * length)
		for i := range n {
			t := float64(i) / float64(rate)
			env := peak * math.Exp(-(t-attack)*12)
			if t < attack {
				env = peak * t / attack
			}
			out[start+i] = float32(env * math.Sin(2*math.Pi*freq*t))
		}
	}
	return out
}
