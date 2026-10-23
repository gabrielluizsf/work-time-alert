package desktop

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestColumnsStackOnANarrowWindow(t *testing.T) {
	timer, alerts, wide := columns(600, 20, 900)
	assert.False(t, wide)
	assert.Equal(t, timer.Width, 560)
	assert.Equal(t, alerts.X, 20)
}

func TestColumnsSitSideBySideOnAWideWindow(t *testing.T) {
	timer, alerts, wide := columns(1200, 20, 900)
	assert.True(t, wide)
	assert.Equal(t, timer.X+timer.Width+20, alerts.X)
	assert.Equal(t, alerts.X+alerts.Width, 1180)
}

func TestClampScroll(t *testing.T) {
	assert.Equal(t, clampScroll(-10, 1000, 600), 0)
	assert.Equal(t, clampScroll(900, 1000, 600), 400)
	assert.Equal(t, clampScroll(50, 300, 600), 0)
}
