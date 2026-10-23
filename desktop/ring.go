package desktop

import (
	"math"

	"github.com/gabrielluizsf/antui/canvas"
)

// arcPoints walks a circle clockwise from twelve o'clock for fraction of a
// turn. The step count follows the arc's length, so a full ring is smooth and
// a sliver costs a handful of points.
func arcPoints(cx, cy, r, fraction float64) []canvas.Point {
	fraction = min(max(fraction, 0), 1)
	steps := max(int(math.Ceil(fraction*r*2*math.Pi/3)), 2)
	pts := make([]canvas.Point, steps+1)
	for i := range pts {
		a := -math.Pi/2 + 2*math.Pi*fraction*float64(i)/float64(steps)
		pts[i] = canvas.Point{X: cx + r*math.Cos(a), Y: cy + r*math.Sin(a)}
	}
	return pts
}

// ring draws the progress ring with the remaining time in the middle.
func (v *view) ring(cx, cy, r int, fraction float64, label string) {
	style := canvas.StrokeStyle{Width: float64(v.px(12)), Cap: canvas.CapRound}
	track := canvas.NewPath()
	track.AddPolyline(arcPoints(float64(cx), float64(cy), float64(r), 1), true)
	v.win.Canvas().StrokePath(track, v.pal.SurfaceVariant, style)

	if fraction > 0 {
		arc := canvas.NewPath()
		arc.AddPolyline(arcPoints(float64(cx), float64(cy), float64(r), fraction), false)
		v.win.Canvas().StrokePath(arc, v.pal.RingColor(fraction), style)
	}

	box := canvas.Area{X: cx - r, Y: cy - r, Width: 2 * r, Height: 2 * r}
	v.textCentered(v.clock, box, label, v.pal.OnSurface)
}
