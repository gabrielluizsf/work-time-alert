package desktop

import (
	"errors"
	"testing"
	"time"

	"github.com/gabrielluizsf/work-time-alert/store"
	"github.com/i9si-sistemas/assert"
)

type memStore struct {
	saved   []store.Alert
	saves   int
	loadErr error
	saveErr error
}

func (m *memStore) Load() ([]store.Alert, error) { return m.saved, m.loadErr }
func (m *memStore) Save(al ...store.Alert) error {
	m.saves++
	if m.saveErr != nil {
		return m.saveErr
	}
	m.saved = append([]store.Alert(nil), al...)
	return nil
}

type clock struct{ t time.Time }

func (c *clock) now() time.Time          { return c.t }
func (c *clock) advance(d time.Duration) { c.t = c.t.Add(d) }

func newTestApp(st store.Store) (*App, *clock, *[]string) {
	c := &clock{t: epoch}
	rang := &[]string{}
	app := NewApp(Options{
		Lang:   Portuguese,
		Store:  st,
		Now:    c.now,
		OnDone: func(title, _ string) { *rang = append(*rang, title) },
	})
	return app, c, rang
}

func TestNewAppStartsWithTheDefaultsSelected(t *testing.T) {
	app, _, _ := newTestApp(&memStore{})
	assert.Equal(t, app.Alerts.Len(), 2)
	active, ok := app.Active()
	assert.True(t, ok)
	assert.Equal(t, active.Name, "Hora de trabalhar")
	assert.Equal(t, app.Timer.Seconds(epoch), 25*60)
}

func TestNewAppTellsWhenTheSavedAlertsCannotBeRead(t *testing.T) {
	app, _, _ := newTestApp(&memStore{loadErr: errors.New("tampered")})
	assert.True(t, app.Snack.Visible(epoch))
	assert.Equal(t, app.Snack.Message, Portuguese.T("loadFailed"))
	assert.Equal(t, app.Alerts.Len(), 2)
}

func TestSubmitAddsSelectsAndSaves(t *testing.T) {
	st := &memStore{}
	app, _, _ := newTestApp(st)
	app.Form = Form{Name: "Reunião", Minutes: "30"}
	app.Submit()

	active, _ := app.Active()
	assert.Equal(t, active.Name, "Reunião")
	assert.Equal(t, len(st.saved), 3)
	assert.Equal(t, app.Form, Form{})
}

func TestSubmitRefusesAnInvalidAlert(t *testing.T) {
	st := &memStore{}
	app, _, _ := newTestApp(st)
	app.Form = Form{Name: "", Minutes: "30"}
	app.Submit()
	assert.Equal(t, st.saves, 0)
	assert.Equal(t, app.Snack.Message, Portuguese.T("invalidAlert"))
	assert.Equal(t, app.Form.Minutes, "30")
}

func TestEditingTheActiveAlertReloadsTheTimer(t *testing.T) {
	app, _, _ := newTestApp(&memStore{})
	app.Edit(1)
	app.Form.Minutes = "50"
	app.Submit()
	assert.Equal(t, app.Timer.Seconds(epoch), 50*60)
}

func TestDeletingTheActiveAlertMovesToTheNextAndUndoBringsItBack(t *testing.T) {
	st := &memStore{}
	app, c, _ := newTestApp(st)
	app.Delete(1)

	active, _ := app.Active()
	assert.Equal(t, active.ID, int64(2))
	assert.Equal(t, len(st.saved), 1)

	c.advance(5 * time.Second) // the undo stays longer than a plain message
	assert.True(t, app.Snack.Visible(c.now()))
	app.Snack.Act()
	assert.Equal(t, app.Alerts.All()[0].ID, int64(1))
	assert.Equal(t, len(st.saved), 2)
}

func TestDeletingTheLastAlertEmptiesTheTimer(t *testing.T) {
	app, _, _ := newTestApp(&memStore{})
	app.Delete(1)
	app.Delete(2)
	_, ok := app.Active()
	assert.False(t, ok)
	assert.False(t, app.Timer.CanStart())

	app.Snack.Act()
	active, ok := app.Active()
	assert.True(t, ok)
	assert.Equal(t, active.ID, int64(2))
}

func TestASaveFailureIsShown(t *testing.T) {
	app, _, _ := newTestApp(&memStore{saveErr: errors.New("disk full")})
	app.Form = Form{Name: "X", Seconds: "1"}
	app.Submit()
	assert.Equal(t, app.Snack.Message, Portuguese.T("saveFailed"))
}

func TestTimeUpOpensTheDialogAndRingsOnce(t *testing.T) {
	app, c, rang := newTestApp(nil)
	app.Toggle()
	c.advance(25 * time.Minute)
	app.Update()
	app.Update()

	assert.True(t, app.Dialog.Open)
	assert.Equal(t, app.Dialog.Title, "Hora de trabalhar")
	assert.Equal(t, len(*rang), 1)

	app.Snooze()
	assert.False(t, app.Dialog.Open)
	assert.True(t, app.Timer.Running())
	assert.Equal(t, app.Timer.Seconds(c.now()), 300)
}

func TestToggleStartsAndPauses(t *testing.T) {
	app, c, _ := newTestApp(nil)
	app.Toggle()
	assert.True(t, app.Timer.Running())
	c.advance(time.Minute)
	app.Toggle()
	assert.False(t, app.Timer.Running())
	assert.Equal(t, app.Timer.Seconds(c.now()), 24*60)
}
