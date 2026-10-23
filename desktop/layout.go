package desktop

import "github.com/gabrielluizsf/antui/canvas"

// columns places the timer card and the alerts card: side by side on a wide
// window, one above the other on a narrow one. y and heights are filled in by
// the caller, which knows how tall each card is.
func columns(width, pad, minWide int) (timer, alerts canvas.Area, wide bool) {
	inner := width - 2*pad
	if width < minWide {
		return canvas.Area{X: pad, Width: inner}, canvas.Area{X: pad, Width: inner}, false
	}
	tw := max((inner-pad)*5/12, minWide*5/12-pad)
	timer = canvas.Area{X: pad, Width: tw}
	alerts = canvas.Area{X: pad + tw + pad, Width: inner - tw - pad}
	return timer, alerts, true
}

// clampScroll keeps the page between its top and the point where its bottom
// meets the bottom of the window.
func clampScroll(scroll, content, window int) int {
	return min(max(scroll, 0), max(content-window, 0))
}
