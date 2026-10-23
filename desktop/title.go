package desktop

import "github.com/gabrielluizsf/work-time-alert/store"

// windowTitle shows the countdown in the taskbar while it runs, so the time
// left is visible with the window minimised.
func windowTitle(app *App) string {
	if !app.Timer.Running() {
		return appTitle
	}
	left := store.ClockTime(app.Timer.Seconds(app.Now()))
	if al, ok := app.Active(); ok {
		return left + " · " + al.Name
	}
	return left + " · " + appTitle
}
