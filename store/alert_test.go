package store

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestAlertTotalAndDuration(t *testing.T) {
	al := Alert{Name: "Foco", Hours: 1, Minutes: 5}
	assert.Equal(t, al.Total(), 3900)
	assert.Equal(t, al.Duration(), "01:05:00")
}

func TestAlertValid(t *testing.T) {
	assert.True(t, Alert{Name: "Foco", Seconds: 1}.Valid())
	assert.False(t, Alert{Name: "   ", Seconds: 1}.Valid())
	assert.False(t, Alert{Name: "Foco"}.Valid())
}

func TestClockTimeNeverGoesNegative(t *testing.T) {
	assert.Equal(t, ClockTime(-5), "00:00:00")
	assert.Equal(t, ClockTime(59), "00:00:59")
	assert.Equal(t, ClockTime(36000), "10:00:00")
}
