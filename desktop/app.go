package desktop

import (
	"time"

	"github.com/gabrielluizsf/work-time-alert/store"
)

// SnoozeTime is how much a snooze adds to a finished alert.
const SnoozeTime = 5 * time.Minute

// App is everything the window shows, kept apart from the drawing so every
// action can be tested without a screen.
type App struct {
	Lang   Lang
	Alerts *Alerts
	Timer  Timer
	Form   Form
	Snack  Snackbar
	Dialog Dialog

	activeID int64
	store    store.Store // nil keeps everything in memory
	now      func() time.Time
	onDone   func(title, message string)
}

// Dialog is the "time is up" window.
type Dialog struct {
	Open           bool
	Title, Message string
}

// Options are what NewApp needs from the outside world.
type Options struct {
	Lang   Lang
	Store  store.Store
	Now    func() time.Time
	OnDone func(title, message string) // ring and notify; may be nil
}

// NewApp loads the saved alerts — or the defaults on a first run — and selects
// the first one, ready to start.
func NewApp(opt Options) *App {
	a := &App{Lang: opt.Lang, store: opt.Store, now: opt.Now, onDone: opt.OnDone}
	if a.now == nil {
		a.now = time.Now
	}
	if a.onDone == nil {
		a.onDone = func(string, string) {}
	}

	items, err := a.load()
	if err != nil {
		a.Snack.Show(a.now(), a.Lang.T("loadFailed"), "", nil)
	}
	if items == nil {
		items = DefaultAlerts()
	}
	a.Alerts = NewAlerts(items)
	if all := a.Alerts.All(); len(all) > 0 {
		a.Select(all[0].ID)
	}
	return a
}

func (a *App) load() ([]store.Alert, error) {
	if a.store == nil {
		return nil, nil
	}
	return a.store.Load()
}

// Active returns the alert the timer is loaded with.
func (a *App) Active() (store.Alert, bool) {
	if a.activeID == 0 {
		return store.Alert{}, false
	}
	return a.Alerts.Find(a.activeID)
}

// Now is the app's clock.
func (a *App) Now() time.Time { return a.now() }
