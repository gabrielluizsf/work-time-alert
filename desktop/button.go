package desktop

import "github.com/gabrielluizsf/antui/canvas"

type buttonStyle int

const (
	primaryButton buttonStyle = iota
	secondaryButton
	dangerButton
	textButton
)

// colors gives a style's resting fill, hover fill and label colour.
func (s buttonStyle) colors(p Palette) (fill, hover, label canvas.Color) {
	switch s {
	case secondaryButton:
		return p.SecondaryContainer, canvas.Shade(p.SecondaryContainer, -0.06), p.OnSecondaryContainer
	case dangerButton:
		return p.ErrorContainer, canvas.Shade(p.ErrorContainer, -0.06), p.OnErrorContainer
	case textButton:
		return p.Surface, p.SurfaceVariant, p.Primary
	default:
		return p.Primary, p.PrimaryHover, p.OnPrimary
	}
}

// button draws an AntUI button in one of the app's styles and reports a
// click. A disabled or inert button is only painted, so it never takes focus
// or a click.
func (v *view) button(area canvas.Area, label string, style buttonStyle, enabled bool) bool {
	fill, hover, fg := style.colors(v.pal)
	if !enabled {
		if style != textButton {
			v.win.FillRoundRect(canvas.Round(area, area.Height/2), canvas.Fade(v.pal.OnSurface, 0x1F))
		}
		v.textCentered(v.body, area, label, canvas.Fade(v.pal.OnSurface, 0x61))
		return false
	}
	if v.inert {
		v.win.FillRoundRect(canvas.Round(area, area.Height/2), fill)
		v.textCentered(v.body, area, label, fg)
		return false
	}
	th := v.win.Theme()
	saved := *th
	th.Accent, th.AccentHover, th.TextOnAccent, th.Radius = fill, hover, fg, area.Height/2
	clicked := v.win.Button(area, label)
	*th = saved
	return clicked
}

// buttonWidth is the width a label needs inside a button.
func (v *view) buttonWidth(label string) int {
	return v.body.Width(label) + v.px(40)
}
