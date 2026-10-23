package desktop

import "github.com/gabrielluizsf/antui"

// paletteFor follows the system's light or dark mode, read every frame so a
// switch in the system settings lands at once.
func paletteFor(win *antui.Window) Palette {
	if dark, ok := win.SystemDark(); ok && dark {
		return DarkPalette()
	}
	return LightPalette()
}

// applyTheme hands the palette to AntUI's own widgets — the text fields and
// the focus ring — so they match the hand-drawn parts.
func (v *view) applyTheme() {
	th := v.win.Theme()
	th.Background = v.pal.Background
	th.Surface = v.pal.Surface
	th.SurfaceHover = v.pal.SurfaceVariant
	th.Border = v.pal.Outline
	th.Accent = v.pal.Primary
	th.AccentHover = v.pal.PrimaryHover
	th.Text = v.pal.OnSurface
	th.TextMuted = v.pal.OnSurfaceVariant
	th.TextOnAccent = v.pal.OnPrimary
	th.Radius = v.px(12)
	th.Padding = v.px(8)
}
