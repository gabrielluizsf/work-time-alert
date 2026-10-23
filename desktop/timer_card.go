package desktop

import (
	"github.com/gabrielluizsf/antui/canvas"
	"github.com/gabrielluizsf/work-time-alert/store"
)

// timerCardHeight is how tall timerCard draws at this scale.
func (v *view) timerCardHeight() int {
	return v.px(24+32+24) + 2*v.ringRadius() + v.px(28+44+16) + v.small.Height() + v.px(24)
}

func (v *view) ringRadius() int { return v.px(124) }

// timerCard draws the active alert, the ring and the start/pause/reset row.
func (v *view) timerCard(app *App, a canvas.Area) {
	v.card(a)
	now := app.Now()
	pad := v.px(24)
	y := a.Y + pad

	cw := v.chip(a.X+pad, y+v.px(4), app.Lang.T("activeLabel"), v.pal.PrimaryContainer, v.pal.OnPrimaryContainer)
	name := app.Lang.T("noAlert")
	if al, ok := app.Active(); ok {
		name = al.Name
	}
	nx := a.X + pad + cw + v.px(12)
	v.text(v.body, nx, y+v.px(16)-v.body.Height()/2, v.ellipsis(v.body, name, a.X+a.Width-pad-nx), v.pal.OnSurface)
	y += v.px(32 + 24)

	r := v.ringRadius()
	v.ring(a.X+a.Width/2, y+r, r, app.Timer.Fraction(now), store.ClockTime(app.Timer.Seconds(now)))
	y += 2*r + v.px(28)

	v.controls(app, a.X+pad, y, a.Width-2*pad)
	y += v.px(44 + 16)

	hint := v.ellipsis(v.small, app.Lang.T("hint"), a.Width-2*pad)
	v.text(v.small, a.X+(a.Width-v.small.Width(hint))/2, y, hint, v.pal.OnSurfaceVariant)
}

// controls draws the three timer buttons centred in width.
func (v *view) controls(app *App, x, y, width int) {
	labels := [3]string{app.Lang.T("start"), app.Lang.T("pause"), app.Lang.T("reset")}
	gap := v.px(10)
	bw := 0
	for _, l := range labels {
		bw = max(bw, v.buttonWidth(l))
	}
	bw = min(bw, (width-2*gap)/3)
	bx := x + (width-(3*bw+2*gap))/2
	h := v.px(44)

	if v.button(canvas.Area{X: bx, Y: y, Width: bw, Height: h}, labels[0], primaryButton, app.Timer.CanStart()) {
		app.Start()
	}
	bx += bw + gap
	if v.button(canvas.Area{X: bx, Y: y, Width: bw, Height: h}, labels[1], secondaryButton, app.Timer.Running()) {
		app.Pause()
	}
	bx += bw + gap
	if v.button(canvas.Area{X: bx, Y: y, Width: bw, Height: h}, labels[2], dangerButton, app.Timer.CanReset()) {
		app.Reset()
	}
}
