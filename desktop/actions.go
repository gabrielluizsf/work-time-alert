package desktop

import (
	"time"

	"github.com/gabrielluizsf/work-time-alert/store"
)

// Select loads an alert into the timer, stopped and full.
func (a *App) Select(id int64) {
	al, ok := a.Alerts.Find(id)
	if !ok {
		return
	}
	a.activeID = id
	a.Timer.Set(time.Duration(al.Total()) * time.Second)
}

// Submit adds the alert in the form, or saves the one being edited.
func (a *App) Submit() {
	al, ok := a.Form.Alert()
	if !ok {
		a.Snack.Show(a.now(), a.Lang.T("invalidAlert"), "", nil)
		return
	}
	if a.Form.Editing() {
		if !a.Alerts.Update(al) {
			return
		}
		a.Form = Form{}
		if a.activeID == al.ID {
			a.Select(al.ID)
		}
		a.Snack.Show(a.now(), a.Lang.T("updated")+` "`+al.Name+`"`, "", nil)
		a.persist()
		return
	}
	added := a.Alerts.Add(al)
	a.Form = Form{}
	a.Select(added.ID)
	a.Snack.Show(a.now(), a.Lang.T("added")+` "`+added.Name+`"`, "", nil)
	a.persist()
}

// Edit fills the form with an alert.
func (a *App) Edit(id int64) {
	if al, ok := a.Alerts.Find(id); ok {
		a.Form = FormFor(al)
	}
}

// ClearForm empties the form and leaves edit mode.
func (a *App) ClearForm() { a.Form = Form{} }

// Delete removes an alert and offers to undo it. Deleting the active alert
// moves the timer to the first one left.
func (a *App) Delete(id int64) {
	removed, at, ok := a.Alerts.Remove(id)
	if !ok {
		return
	}
	if a.Form.EditingID == id {
		a.Form = Form{}
	}
	if a.activeID == id {
		a.activeID = 0
		a.Timer.Set(0)
		if all := a.Alerts.All(); len(all) > 0 {
			a.Select(all[0].ID)
		}
	}
	a.Snack.Show(a.now(), a.Lang.T("deleted")+` "`+removed.Name+`"`, a.Lang.T("undo"), func() {
		a.undoDelete(removed, at)
	})
	a.persist()
}

func (a *App) undoDelete(al store.Alert, at int) {
	a.Alerts.Restore(al, at)
	if a.activeID == 0 {
		a.Select(al.ID)
	}
	a.persist()
}

// persist saves the list; a failure is told to the user, not swallowed.
func (a *App) persist() {
	if a.store == nil {
		return
	}
	if err := a.store.Save(a.Alerts.All()...); err != nil {
		a.Snack.Show(a.now(), a.Lang.T("saveFailed"), "", nil)
	}
}
