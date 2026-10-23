package desktop

import (
	"testing"

	"github.com/gabrielluizsf/work-time-alert/store"
	"github.com/i9si-sistemas/assert"
)

func TestFormReadsTheFields(t *testing.T) {
	al, ok := Form{Name: "  Foco ", Hours: "1", Minutes: "05", Seconds: ""}.Alert()
	assert.True(t, ok)
	assert.Equal(t, al.Name, "Foco")
	assert.Equal(t, al.Total(), 3900)
}

func TestFormRejectsBlankNameOrZeroTime(t *testing.T) {
	_, ok := Form{Name: "  ", Minutes: "5"}.Alert()
	assert.False(t, ok)
	_, ok = Form{Name: "Foco", Minutes: "abc", Seconds: "-3"}.Alert()
	assert.False(t, ok)
}

func TestFormForRoundTripsAnAlert(t *testing.T) {
	in := store.Alert{ID: 7, Name: "Pausa", Minutes: 5, Seconds: 30}
	f := FormFor(in)
	assert.True(t, f.Editing())
	out, ok := f.Alert()
	assert.True(t, ok)
	assert.Equal(t, out, in)
}
