package desktop

import (
	"strconv"
	"strings"

	"github.com/gabrielluizsf/work-time-alert/store"
)

// Form is the text the alert fields hold. The numbers stay as text while the
// user types and are only read when the form is submitted.
type Form struct {
	Name, Hours, Minutes, Seconds string
	EditingID                     int64 // 0 while adding a new alert
}

// Alert reads the fields into an alert. Blank or broken numbers count as zero,
// like the web form; ok is false when the alert could not be saved.
func (f Form) Alert() (store.Alert, bool) {
	al := store.Alert{
		ID:      f.EditingID,
		Name:    strings.TrimSpace(f.Name),
		Hours:   number(f.Hours),
		Minutes: number(f.Minutes),
		Seconds: number(f.Seconds),
	}
	return al, al.Valid()
}

// Editing reports whether the form is changing an existing alert.
func (f Form) Editing() bool { return f.EditingID != 0 }

// FormFor fills the form with an alert to edit.
func FormFor(al store.Alert) Form {
	return Form{
		Name:      al.Name,
		Hours:     strconv.Itoa(al.Hours),
		Minutes:   strconv.Itoa(al.Minutes),
		Seconds:   strconv.Itoa(al.Seconds),
		EditingID: al.ID,
	}
}

func number(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || n < 0 {
		return 0
	}
	return n
}
