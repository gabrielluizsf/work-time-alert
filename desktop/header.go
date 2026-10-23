package desktop

import (
	"time"

	"github.com/gabrielluizsf/antui/canvas"
)

// header draws the brand on the left and the current time on the right, and
// returns its height.
func (v *view) header(app *App, x, y, width int) int {
	h := v.px(56)
	icon := canvas.Area{X: x, Y: y + (h-v.px(44))/2, Width: v.px(44), Height: v.px(44)}
	v.win.FillRoundRect(canvas.Round(icon, v.px(14)), v.pal.Primary)
	v.clockGlyph(icon)

	tx := icon.X + icon.Width + v.px(14)
	v.text(v.title, tx, y+v.px(6), "Work Time Alert", v.pal.OnSurface)
	v.text(v.small, tx, y+v.px(10)+v.title.Height(), app.Lang.T("tagline"), v.pal.OnSurfaceVariant)

	now := time.Now().Format("15:04:05")
	cw := v.small.Width(now) + v.px(20)
	if tx+v.title.Width("Work Time Alert")+v.px(16) < x+width-cw {
		v.chip(x+width-cw, y+(h-v.small.Height()-v.px(8))/2, now, v.pal.Surface, v.pal.OnSurfaceVariant)
	}
	return h
}

// clockGlyph draws a small clock face in the brand icon.
func (v *view) clockGlyph(a canvas.Area) {
	cx, cy := a.X+a.Width/2, a.Y+a.Height/2
	style := canvas.StrokeStyle{Width: float64(v.px(2.5)), Cap: canvas.CapRound}
	face := canvas.NewPath()
	face.AddPolyline(arcPoints(float64(cx), float64(cy), float64(a.Width)*0.28, 1), true)
	v.win.Canvas().StrokePath(face, v.pal.OnPrimary, style)

	hands := canvas.NewPath()
	hands.AddPolyline([]canvas.Point{
		{X: float64(cx), Y: float64(cy) - float64(a.Width)*0.16},
		{X: float64(cx), Y: float64(cy)},
		{X: float64(cx) + float64(a.Width)*0.12, Y: float64(cy) + float64(a.Width)*0.08},
	}, false)
	v.win.Canvas().StrokePath(hands, v.pal.OnPrimary, style)
}
