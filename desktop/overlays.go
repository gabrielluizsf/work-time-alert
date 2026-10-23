package desktop

import "github.com/gabrielluizsf/antui/canvas"

// snackbar draws the message pill at the bottom of the window.
func (v *view) snackbar(app *App) {
	s := &app.Snack
	if !s.Visible(app.Now()) {
		return
	}
	w, h := v.win.Width(), v.px(52)
	pad := v.px(16)
	actionW := 0
	if s.Action != "" {
		actionW = v.buttonWidth(s.Action)
	}
	msgW := v.body.Width(s.Message)
	sw := min(msgW+actionW+3*pad, w-2*pad)
	a := canvas.Area{X: (w - sw) / 2, Y: v.win.Height() - h - v.px(24), Width: sw, Height: h}
	v.win.FillRoundRect(canvas.Round(a, v.px(12)), v.pal.InverseSurface)

	msg := v.ellipsis(v.body, s.Message, sw-actionW-3*pad)
	v.text(v.body, a.X+pad, a.Y+(h-v.body.Height())/2, msg, v.pal.InverseOnSurface)
	if s.Action == "" {
		return
	}
	btn := canvas.Area{X: a.X + sw - pad/2 - actionW, Y: a.Y + v.px(8), Width: actionW, Height: h - v.px(16)}
	saved := v.pal
	v.pal.Surface, v.pal.SurfaceVariant, v.pal.Primary = v.pal.InverseSurface, canvas.Shade(v.pal.InverseSurface, 0.08), v.pal.InversePrimary
	if v.button(btn, s.Action, textButton, true) {
		s.Act()
	}
	v.pal = saved
}

// dialog draws the "time is up" window over a dimmed page.
func (v *view) dialog(app *App) {
	d := app.Dialog
	w, h := v.win.Width(), v.win.Height()
	v.win.FillRect(canvas.Area{Width: w, Height: h}, v.pal.Scrim)

	pad := v.px(24)
	dw := min(v.px(400), w-2*v.px(16))
	bh := v.px(44)
	dh := pad + v.px(48+16) + v.title.Height() + v.px(8) + v.body.Height() + v.px(24) + bh + pad
	a := canvas.Area{X: (w - dw) / 2, Y: (h - dh) / 2, Width: dw, Height: dh}
	v.win.FillRoundRect(canvas.Round(a, v.px(28)), v.pal.Surface)

	icon := canvas.Area{X: a.X + (dw-v.px(48))/2, Y: a.Y + pad, Width: v.px(48), Height: v.px(48)}
	v.win.FillRoundRect(canvas.Round(icon, icon.Width/2), v.pal.PrimaryContainer)
	saved := v.pal.OnPrimary
	v.pal.OnPrimary = v.pal.OnPrimaryContainer
	v.clockGlyph(icon)
	v.pal.OnPrimary = saved

	y := icon.Y + icon.Height + v.px(16)
	title := v.ellipsis(v.title, d.Title, dw-2*pad)
	v.text(v.title, a.X+(dw-v.title.Width(title))/2, y, title, v.pal.OnSurface)
	y += v.title.Height() + v.px(8)
	msg := v.ellipsis(v.body, d.Message, dw-2*pad)
	v.text(v.body, a.X+(dw-v.body.Width(msg))/2, y, msg, v.pal.OnSurfaceVariant)
	y += v.body.Height() + v.px(24)

	ok, snooze := app.Lang.T("ok"), app.Lang.T("snooze")
	ow, sw := v.buttonWidth(ok), v.buttonWidth(snooze)
	right := a.X + dw - pad
	if v.button(canvas.Area{X: right - ow, Y: y, Width: ow, Height: bh}, ok, primaryButton, true) {
		app.CloseDialog()
	}
	if v.button(canvas.Area{X: right - ow - v.px(8) - sw, Y: y, Width: sw, Height: bh}, snooze, secondaryButton, true) {
		app.Snooze()
	}
}
