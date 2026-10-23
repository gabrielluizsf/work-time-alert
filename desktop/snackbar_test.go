package desktop

import (
	"testing"
	"time"

	"github.com/i9si-sistemas/assert"
)

func TestSnackbarHidesAfterItsTime(t *testing.T) {
	var s Snackbar
	s.Show(epoch, "oi", "", nil)
	assert.True(t, s.Visible(epoch.Add(3*time.Second)))
	assert.False(t, s.Visible(epoch.Add(5*time.Second)))
}

func TestSnackbarActionRunsOnceAndHides(t *testing.T) {
	var s Snackbar
	runs := 0
	s.Show(epoch, "oi", "Desfazer", func() { runs++ })
	assert.True(t, s.Visible(epoch.Add(9*time.Second)))
	s.Act()
	s.Act()
	assert.Equal(t, runs, 1)
	assert.False(t, s.Visible(epoch))
}
