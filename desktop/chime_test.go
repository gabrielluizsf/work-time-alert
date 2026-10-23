package desktop

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestBeepsStayInRangeAndEndSilent(t *testing.T) {
	s := Beeps(8000)
	assert.Equal(t, len(s), int(8000*1.42))
	for _, v := range s {
		assert.True(t, v >= -1 && v <= 1)
	}
	assert.Equal(t, s[0], float32(0))
}

func TestChimeFillsSilenceUntilPlayed(t *testing.T) {
	c := &Chime{samples: []float32{0.5, 0.25}}
	c.pos = len(c.samples)
	buf := make([]float32, 4)
	c.fill(buf)
	assert.Equal(t, buf, []float32{0, 0, 0, 0})

	c.Play()
	c.fill(buf)
	assert.Equal(t, buf, []float32{0.5, 0.5, 0.25, 0.25})
}
