package desktop

import (
	"math"
	"time"
)

// Timer is the countdown behind the big clock. It keeps a deadline rather than
// counting frames, so a slow or skipped frame never makes the alert late.
type Timer struct {
	total    time.Duration // the length the ring is measured against
	initial  time.Duration // what Reset goes back to
	left     time.Duration // what was left when the timer last stopped
	deadline time.Time     // when it ends, while running
	running  bool
}

// Set loads a new duration and stops the timer.
func (t *Timer) Set(d time.Duration) {
	*t = Timer{total: d, initial: d, left: d}
}

// Start runs the timer from now. It does nothing when there is no time left.
func (t *Timer) Start(now time.Time) {
	if t.running || t.left <= 0 {
		return
	}
	t.deadline = now.Add(t.left)
	t.running = true
}

// Pause freezes the remaining time.
func (t *Timer) Pause(now time.Time) {
	if !t.running {
		return
	}
	t.left = t.Remaining(now)
	t.running = false
}

// Reset stops the timer and puts back its initial duration.
func (t *Timer) Reset() {
	t.left = t.initial
	t.total = t.initial
	t.running = false
}

// Snooze adds extra time to a finished timer and starts it again; the ring is
// measured against the extra time alone, as the web version does.
func (t *Timer) Snooze(extra time.Duration, now time.Time) {
	t.Pause(now)
	t.left += extra
	t.total = t.left
	t.Start(now)
}

// Tick reports whether the timer reached zero at now. It fires once: the
// timer stops itself so the next frame does not alert again.
func (t *Timer) Tick(now time.Time) bool {
	if !t.running || now.Before(t.deadline) {
		return false
	}
	t.left = 0
	t.running = false
	return true
}

// Remaining is the time still to go at now, never below zero.
func (t *Timer) Remaining(now time.Time) time.Duration {
	if !t.running {
		return t.left
	}
	return max(t.deadline.Sub(now), 0)
}

// Seconds is the remaining time rounded up to whole seconds, which is what the
// clock shows: 00:00:01 stays on screen until the very end.
func (t *Timer) Seconds(now time.Time) int {
	return int(math.Ceil(t.Remaining(now).Seconds()))
}

// Fraction is how much of the ring is still full, from 1 down to 0.
func (t *Timer) Fraction(now time.Time) float64 {
	if t.total <= 0 {
		return 1
	}
	return min(max(float64(t.Remaining(now))/float64(t.total), 0), 1)
}

// Running reports whether the countdown is going.
func (t *Timer) Running() bool { return t.running }

// CanStart reports whether Start would do something.
func (t *Timer) CanStart() bool { return !t.running && t.left > 0 }

// CanReset reports whether Reset would change anything.
func (t *Timer) CanReset() bool { return t.running || t.left != t.initial }
