package desktop

import (
	"slices"

	"github.com/gabrielluizsf/work-time-alert/store"
)

// Alerts is the list the user manages: add, edit, delete and undo a delete.
type Alerts struct {
	items []store.Alert
}

// DefaultAlerts are the two timers a first run starts with, the same ones the
// web version offers.
func DefaultAlerts() []store.Alert {
	return []store.Alert{
		{ID: 1, Name: "Hora de trabalhar", Minutes: 25},
		{ID: 2, Name: "Pausa Curta", Minutes: 5},
	}
}

// NewAlerts wraps a loaded list.
func NewAlerts(items []store.Alert) *Alerts {
	return &Alerts{items: slices.Clone(items)}
}

// All is the list in display order. Callers must not keep it across edits.
func (a *Alerts) All() []store.Alert { return a.items }

// Len is how many alerts there are.
func (a *Alerts) Len() int { return len(a.items) }

// Add appends an alert under a fresh id and returns it.
func (a *Alerts) Add(al store.Alert) store.Alert {
	al.ID = a.nextID()
	a.items = append(a.items, al)
	return al
}

// Update replaces the alert with the same id; it reports whether one existed.
func (a *Alerts) Update(al store.Alert) bool {
	i := a.index(al.ID)
	if i < 0 {
		return false
	}
	a.items[i] = al
	return true
}

// Remove deletes the alert with id and returns it with its old position, so
// an undo can put it back where it was.
func (a *Alerts) Remove(id int64) (store.Alert, int, bool) {
	i := a.index(id)
	if i < 0 {
		return store.Alert{}, -1, false
	}
	removed := a.items[i]
	a.items = slices.Delete(a.items, i, i+1)
	return removed, i, true
}

// Restore puts a removed alert back at its old position.
func (a *Alerts) Restore(al store.Alert, at int) {
	at = min(max(at, 0), len(a.items))
	a.items = slices.Insert(a.items, at, al)
}

// Find returns the alert with id.
func (a *Alerts) Find(id int64) (store.Alert, bool) {
	i := a.index(id)
	if i < 0 {
		return store.Alert{}, false
	}
	return a.items[i], true
}

func (a *Alerts) index(id int64) int {
	return slices.IndexFunc(a.items, func(al store.Alert) bool { return al.ID == id })
}

func (a *Alerts) nextID() int64 {
	var top int64
	for _, al := range a.items {
		top = max(top, al.ID)
	}
	return top + 1
}
