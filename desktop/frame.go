package desktop

import (
	"github.com/gabrielluizsf/antui"
	"github.com/gabrielluizsf/antui/canvas"
)

// page is what survives between frames: the scroll position and the form
// fields drawn last frame, which decide what the keyboard does.
type page struct {
	scroll int
	fields []canvas.Area
}

// frame draws everything once and applies this frame's keys.
func (p *page) frame(v *view, app *App) {
	p.keys(v, app)
	app.Update()
	v.inert = app.Dialog.Open

	win := v.win
	win.Clear(v.pal.Background)

	pad := v.px(24)
	width := win.Width()
	timerA, alertsA, wide := columns(width, pad, v.px(900))
	timerA.Height = v.timerCardHeight()
	alertsA.Height = v.alertsCardHeight(app.Lang, alertsA.Width, app.Alerts.Len())

	headerH := v.px(56)
	content := pad + headerH + pad + timerA.Height + pad + alertsA.Height + pad
	if wide {
		content = pad + headerH + pad + max(timerA.Height, alertsA.Height) + pad
	}
	// Leave room below the last card for the snackbar.
	content += v.px(76)
	if !v.inert {
		p.scroll -= win.Wheel() * v.px(48)
	}
	p.scroll = clampScroll(p.scroll, content, win.Height())

	y := pad - p.scroll
	v.header(app, pad, y, width-2*pad)
	y += headerH + pad
	timerA.Y = y
	alertsA.Y = y
	if !wide {
		alertsA.Y = y + timerA.Height + pad
	}
	v.timerCard(app, timerA)
	p.fields = v.alertsCard(app, alertsA)

	v.snackbar(app)
	if app.Dialog.Open {
		v.inert = false
		v.dialog(app)
	}
}

// keys handles the shortcuts: Space starts or pauses when nothing has the
// focus, Enter in a form field submits, and Enter or Escape closes the dialog.
func (p *page) keys(v *view, app *App) {
	win := v.win
	if app.Dialog.Open {
		if win.KeyPressed(antui.KeyEnter) || win.KeyPressed(antui.KeyEscape) {
			app.CloseDialog()
		}
		return
	}
	focus := win.WidgetFocus()
	if focus == 0 && win.KeyPressed(antui.KeySpace) {
		app.Toggle()
		return
	}
	if win.KeyPressed(antui.KeyEnter) && p.inForm(win, focus) {
		app.Submit()
	}
}

func (p *page) inForm(win *antui.Window, focus uint32) bool {
	for _, a := range p.fields {
		if focus != 0 && win.WidgetID("input", a, "") == focus {
			return true
		}
	}
	return false
}
