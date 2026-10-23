package desktop

import (
	"testing"
	"time"

	"github.com/i9si-sistemas/assert"
)

func TestWindowTitleShowsTheCountdownWhileRunning(t *testing.T) {
	app, c, _ := newTestApp(nil)
	assert.Equal(t, windowTitle(app), "Work Time Alert")
	app.Start()
	c.advance(time.Minute)
	assert.Equal(t, windowTitle(app), "00:24:00 · Hora de trabalhar")
}
