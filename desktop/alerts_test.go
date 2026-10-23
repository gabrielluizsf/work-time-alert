package desktop

import (
	"testing"

	"github.com/gabrielluizsf/work-time-alert/store"
	"github.com/i9si-sistemas/assert"
)

func TestAlertsAddGivesAFreshID(t *testing.T) {
	list := NewAlerts(DefaultAlerts())
	added := list.Add(store.Alert{Name: "Reunião", Minutes: 30})
	assert.Equal(t, added.ID, int64(3))
	assert.Equal(t, list.Len(), 3)
}

func TestAlertsUpdateReplacesByID(t *testing.T) {
	list := NewAlerts(DefaultAlerts())
	assert.True(t, list.Update(store.Alert{ID: 2, Name: "Pausa", Minutes: 10}))
	got, ok := list.Find(2)
	assert.True(t, ok)
	assert.Equal(t, got.Name, "Pausa")
	assert.False(t, list.Update(store.Alert{ID: 99}))
}

func TestAlertsRemoveAndRestoreKeepThePosition(t *testing.T) {
	list := NewAlerts(DefaultAlerts())
	removed, at, ok := list.Remove(1)
	assert.True(t, ok)
	assert.Equal(t, at, 0)
	assert.Equal(t, list.All()[0].ID, int64(2))

	list.Restore(removed, at)
	assert.Equal(t, list.All()[0].ID, int64(1))

	_, _, ok = list.Remove(99)
	assert.False(t, ok)
}

func TestAlertsDoNotShareTheLoadedSlice(t *testing.T) {
	loaded := DefaultAlerts()
	list := NewAlerts(loaded)
	list.Update(store.Alert{ID: 1, Name: "Outro", Minutes: 1})
	assert.Equal(t, loaded[0].Name, "Hora de trabalhar")
}
