package desktop

import "time"

// Snackbar is the short message at the bottom of the window, optionally with
// one action such as "Undo".
type Snackbar struct {
	Message  string
	Action   string
	OnAction func()
	until    time.Time
}

const (
	snackbarTime       = 4 * time.Second
	snackbarActionTime = 10 * time.Second
)

// Show replaces whatever was showing. A message with an action stays longer,
// giving the user time to reach the button.
func (s *Snackbar) Show(now time.Time, message, action string, onAction func()) {
	d := snackbarTime
	if action != "" && onAction != nil {
		d = snackbarActionTime
	}
	*s = Snackbar{Message: message, Action: action, OnAction: onAction, until: now.Add(d)}
}

// Visible reports whether the snackbar is on screen at now.
func (s *Snackbar) Visible(now time.Time) bool {
	return s.Message != "" && now.Before(s.until)
}

// Act runs the action once and hides the snackbar.
func (s *Snackbar) Act() {
	run := s.OnAction
	*s = Snackbar{}
	if run != nil {
		run()
	}
}
