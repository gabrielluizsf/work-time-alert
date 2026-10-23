package desktop

import (
	"strconv"

	"github.com/gabrielluizsf/antui/canvas"
	"github.com/gabrielluizsf/work-time-alert/store"
)

// rowLayout reports whether a row is narrow enough to put its buttons under
// the name, and the row's height either way.
func (v *view) rowLayout(lang Lang, width int) (stacked bool, height int) {
	buttons := v.rowButtonsWidth(lang)
	if width-v.px(32)-buttons < v.px(140) {
		return true, v.px(14) + v.body.Height() + v.px(4) + v.small.Height() + v.px(12+36+14)
	}
	return false, v.px(68)
}

func (v *view) rowButtonsWidth(lang Lang) int {
	w := 0
	for _, k := range []string{"use", "edit", "delete"} {
		w += v.rowButtonWidth(lang.T(k))
	}
	return w + 2*v.px(8)
}

func (v *view) rowButtonWidth(label string) int { return v.body.Width(label) + v.px(28) }

// alertsCardHeight is how tall alertsCard draws for n alerts.
func (v *view) alertsCardHeight(lang Lang, width, n int) int {
	pad := v.px(24)
	h := pad + v.px(32+20) + v.formHeight() + v.px(24)
	if n == 0 {
		h += v.px(120)
	} else {
		_, row := v.rowLayout(lang, width-2*pad)
		h += n*row + (n-1)*v.px(8)
	}
	return h + pad
}

// alertsCard draws the title, the form and the list, and returns the form's
// field areas.
func (v *view) alertsCard(app *App, a canvas.Area) []canvas.Area {
	v.card(a)
	pad := v.px(24)
	x, y, width := a.X+pad, a.Y+pad, a.Width-2*pad

	tw := v.text(v.title, x, y+v.px(16)-v.title.Height()/2, app.Lang.T("alertsTitle"), v.pal.OnSurface)
	v.chip(x+tw+v.px(10), y+v.px(16)-(v.small.Height()+v.px(8))/2, strconv.Itoa(app.Alerts.Len()), v.pal.SecondaryContainer, v.pal.OnSecondaryContainer)
	y += v.px(32 + 20)

	fields := v.alertForm(app, x, y, width)
	y += v.formHeight() + v.px(24)

	if app.Alerts.Len() == 0 {
		v.emptyState(app.Lang, canvas.Area{X: x, Y: y, Width: width, Height: v.px(120)})
		return fields
	}
	active, _ := app.Active()
	// The list can change under a click (Delete), so it is walked over a copy.
	for _, al := range append([]store.Alert(nil), app.Alerts.All()...) {
		y += v.alertRow(app, al, al.ID == active.ID, x, y, width) + v.px(8)
	}
	return fields
}

// alertRow draws one alert and its Use, Edit and Delete buttons.
func (v *view) alertRow(app *App, al store.Alert, active bool, x, y, width int) int {
	stacked, h := v.rowLayout(app.Lang, width)
	bg, fg := v.pal.SurfaceVariant, v.pal.OnSurface
	if active {
		bg, fg = v.pal.PrimaryContainer, v.pal.OnPrimaryContainer
	}
	v.win.FillRoundRect(canvas.Round(canvas.Area{X: x, Y: y, Width: width, Height: h}, v.px(16)), bg)

	inner := v.px(16)
	textW := width - 2*inner
	if !stacked {
		textW -= v.rowButtonsWidth(app.Lang) + inner
	}
	ty := y + v.px(14)
	if !stacked {
		ty = y + (h-v.body.Height()-v.px(4)-v.small.Height())/2
	}
	v.text(v.body, x+inner, ty, v.ellipsis(v.body, al.Name, textW), fg)
	v.text(v.small, x+inner, ty+v.body.Height()+v.px(4), al.Duration(), v.pal.OnSurfaceVariant)

	bh := v.px(36)
	by := y + (h-bh)/2
	bx := x + width - inner - v.rowButtonsWidth(app.Lang)
	if stacked {
		by = y + h - v.px(14) - bh
		bx = x + inner
	}
	buttons := []struct {
		key   string
		style buttonStyle
		do    func(int64)
	}{
		{"use", secondaryButton, app.Select},
		{"edit", textButton, app.Edit},
		{"delete", dangerButton, app.Delete},
	}
	for _, b := range buttons {
		label := app.Lang.T(b.key)
		w := v.rowButtonWidth(label)
		if v.button(canvas.Area{X: bx, Y: by, Width: w, Height: bh}, label, b.style, true) {
			b.do(al.ID)
		}
		bx += w + v.px(8)
	}
	return h
}

// emptyState is shown in place of the list when there are no alerts.
func (v *view) emptyState(lang Lang, a canvas.Area) {
	title, desc := lang.T("emptyTitle"), lang.T("emptyDesc")
	ty := a.Y + (a.Height-v.body.Height()-v.px(6)-v.small.Height())/2
	v.text(v.body, a.X+(a.Width-v.body.Width(title))/2, ty, title, v.pal.OnSurface)
	desc = v.ellipsis(v.small, desc, a.Width)
	v.text(v.small, a.X+(a.Width-v.small.Width(desc))/2, ty+v.body.Height()+v.px(6), desc, v.pal.OnSurfaceVariant)
}
