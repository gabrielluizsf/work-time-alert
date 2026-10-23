package desktop

import "github.com/gabrielluizsf/antui/canvas"

// input draws a labelled text field. While a dialog is open the field is only
// painted, so a click behind the dialog cannot focus it.
func (v *view) input(area canvas.Area, label string, text *string) {
	v.text(v.small, area.X+v.px(4), area.Y-v.small.Height()-v.px(6), label, v.pal.OnSurfaceVariant)
	if !v.inert {
		v.win.Input(area, text)
		return
	}
	r := canvas.Round(area, v.win.Theme().Radius)
	v.win.FillRoundRect(r, v.pal.Surface)
	v.win.RoundRect(r, v.pal.Outline)
	v.text(v.body, area.X+8, area.Y+(area.Height-v.body.Height())/2, v.ellipsis(v.body, *text, area.Width-16), v.pal.OnSurface)
}
