package desktop

import (
	"testing"
	"time"

	"github.com/i9si-sistemas/assert"
)

var epoch = time.Date(2026, 1, 1, 9, 0, 0, 0, time.UTC)

func TestTimerCountsDownFromTheDeadline(t *testing.T) {
	var tm Timer
	tm.Set(10 * time.Second)
	tm.Start(epoch)

	assert.Equal(t, tm.Seconds(epoch.Add(2500*time.Millisecond)), 8)
	assert.Equal(t, tm.Fraction(epoch.Add(5*time.Second)), 0.5)
	assert.False(t, tm.Tick(epoch.Add(9*time.Second)))
	assert.True(t, tm.Tick(epoch.Add(11*time.Second)))
	assert.False(t, tm.Running())
	assert.Equal(t, tm.Seconds(epoch.Add(11*time.Second)), 0)
}

func TestTimerTicksOnlyOnce(t *testing.T) {
	var tm Timer
	tm.Set(time.Second)
	tm.Start(epoch)
	assert.True(t, tm.Tick(epoch.Add(2*time.Second)))
	assert.False(t, tm.Tick(epoch.Add(3*time.Second)))
}

func TestTimerPauseKeepsTheRemainingTime(t *testing.T) {
	var tm Timer
	tm.Set(10 * time.Second)
	tm.Start(epoch)
	tm.Pause(epoch.Add(4 * time.Second))

	later := epoch.Add(time.Hour)
	assert.Equal(t, tm.Seconds(later), 6)
	tm.Start(later)
	assert.Equal(t, tm.Seconds(later.Add(time.Second)), 5)
}

func TestTimerResetGoesBackToTheStart(t *testing.T) {
	var tm Timer
	tm.Set(10 * time.Second)
	assert.False(t, tm.CanReset())
	tm.Start(epoch)
	tm.Pause(epoch.Add(3 * time.Second))
	assert.True(t, tm.CanReset())

	tm.Reset()
	assert.Equal(t, tm.Seconds(epoch), 10)
	assert.False(t, tm.CanReset())
}

func TestTimerDoesNotStartWithoutTime(t *testing.T) {
	var tm Timer
	assert.False(t, tm.CanStart())
	tm.Start(epoch)
	assert.False(t, tm.Running())
}

func TestTimerSnoozeRestartsWithTheExtraTime(t *testing.T) {
	var tm Timer
	tm.Set(time.Second)
	tm.Start(epoch)
	tm.Tick(epoch.Add(time.Second))

	tm.Snooze(5*time.Minute, epoch.Add(time.Second))
	assert.True(t, tm.Running())
	assert.Equal(t, tm.Seconds(epoch.Add(time.Second)), 300)
	assert.Equal(t, tm.Fraction(epoch.Add(time.Second)), 1.0)
}
