package desktop

import (
	"math"

	"github.com/gabrielluizsf/antui"
	"github.com/gabrielluizsf/antui/canvas"
)

// view is one frame's drawing context: the window, the colours, the display
// scale and the faces measured at it.
type view struct {
	win   *antui.Window
	pal   Palette
	scale float64
	inert bool // a dialog is open: draw the page, but let nothing be clicked

	body, small, title, clock *canvas.Face
}

// px turns a length in points into pixels at the display's scale.
func (v *view) px(points float64) int { return int(math.Round(points * v.scale)) }

// text draws s with its top at y and returns its width.
func (v *view) text(f *canvas.Face, x, y int, s string, c canvas.Color) int {
	return f.Draw(canvas.FaceDrawParams{
		Canvas: v.win.Canvas(),
		Pos:    canvas.Point{X: float64(x), Y: float64(y + f.Ascent())},
		Text:   s,
		Color:  c,
	})
}

// textCentered draws s centred in area.
func (v *view) textCentered(f *canvas.Face, area canvas.Area, s string, c canvas.Color) {
	x := area.X + (area.Width-f.Width(s))/2
	y := area.Y + (area.Height-f.Height())/2
	v.text(f, x, y, s, c)
}

// ellipsis shortens s with "…" until it fits in width.
func (v *view) ellipsis(f *canvas.Face, s string, width int) string {
	if f.Width(s) <= width {
		return s
	}
	r := []rune(s)
	for len(r) > 0 && f.Width(string(r)+"…") > width {
		r = r[:len(r)-1]
	}
	return string(r) + "…"
}

// card paints a raised surface.
func (v *view) card(a canvas.Area) {
	r := v.px(24)
	v.win.FillRoundRect(canvas.Round(canvas.Area{X: a.X, Y: a.Y + v.px(2), Width: a.Width, Height: a.Height}, r), canvas.Fade(0xFF101828, 0x14))
	v.win.FillRoundRect(canvas.Round(a, r), v.pal.Surface)
}

// chip paints a pill with a label and returns its width.
func (v *view) chip(x, y int, label string, bg, fg canvas.Color) int {
	h := v.small.Height() + v.px(8)
	w := v.small.Width(label) + v.px(20)
	v.win.FillRoundRect(canvas.Round(canvas.Area{X: x, Y: y, Width: w, Height: h}, h/2), bg)
	v.text(v.small, x+v.px(10), y+v.px(4), label, fg)
	return w
}
