package desktop

// Start runs the countdown.
func (a *App) Start() { a.Timer.Start(a.now()) }

// Pause freezes the countdown.
func (a *App) Pause() { a.Timer.Pause(a.now()) }

// Reset puts the active alert's full time back.
func (a *App) Reset() { a.Timer.Reset() }

// Toggle is what the Space key does: start when stopped, pause when running.
func (a *App) Toggle() {
	if a.Timer.Running() {
		a.Pause()
		return
	}
	a.Start()
}

// Update advances the clock; when the time is up it opens the dialog and
// rings once.
func (a *App) Update() {
	if !a.Timer.Tick(a.now()) {
		return
	}
	title := a.Lang.T("attention")
	if al, ok := a.Active(); ok {
		title = al.Name
	}
	message := a.Lang.T("timeUpMessage")
	a.Dialog = Dialog{Open: true, Title: title, Message: message}
	a.onDone(title, message)
}

// CloseDialog dismisses the "time is up" dialog.
func (a *App) CloseDialog() { a.Dialog = Dialog{} }

// Snooze dismisses the dialog and runs five more minutes.
func (a *App) Snooze() {
	a.CloseDialog()
	a.Timer.Snooze(SnoozeTime, a.now())
	a.Snack.Show(a.now(), a.Lang.T("snoozed"), "", nil)
}
