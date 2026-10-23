package desktop

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRingColorFollowsTheTimeLeft(t *testing.T) {
	p := LightPalette()
	assert.Equal(t, p.RingColor(1), p.Primary)
	assert.Equal(t, p.RingColor(0.3), p.Warning)
	assert.Equal(t, p.RingColor(0.05), p.Error)
}
