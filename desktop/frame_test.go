package desktop

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/gabrielluizsf/antui"
	"github.com/gabrielluizsf/antui/canvas"
	"github.com/i9si-sistemas/assert"
)

// renderFrame draws one frame offscreen. With WTA_SNAPSHOTS set to a
// directory, the frame is also saved there as a PNG to look at.
func renderFrame(t *testing.T, name string, w, h int, pal Palette, setup func(*App)) {
	t.Helper()
	win, cv, err := antui.Offscreen(w, h)
	assert.NoError(t, err)
	app, _, _ := newTestApp(nil)
	if setup != nil {
		setup(app)
	}
	v := &view{win: win, scale: 1, pal: pal}
	v.loadFaces()
	v.applyTheme()
	var p page
	assert.True(t, win.Begin())
	p.frame(v, app)
	win.End()

	if dir := os.Getenv("WTA_SNAPSHOTS"); dir != "" {
		savePNG(t, filepath.Join(dir, name+".png"), cv)
	}
}

func savePNG(t *testing.T, path string, cv *canvas.Canvas) {
	img := image.NewNRGBA(image.Rect(0, 0, cv.Width, cv.Height))
	for y := range cv.Height {
		for x := range cv.Width {
			c := uint32(cv.At(x, y))
			img.Set(x, y, color.NRGBA{R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c), A: 0xFF})
		}
	}
	f, err := os.Create(path)
	assert.NoError(t, err)
	defer f.Close()
	assert.NoError(t, png.Encode(f, img))
}

func TestFramesDrawAtEverySize(t *testing.T) {
	renderFrame(t, "wide-light", 1100, 760, LightPalette(), func(a *App) {
		a.Start()
	})
	renderFrame(t, "narrow-dark", 420, 900, DarkPalette(), nil)
	renderFrame(t, "dialog", 1100, 760, LightPalette(), func(a *App) {
		a.Dialog = Dialog{Open: true, Title: "Hora de trabalhar", Message: a.Lang.T("timeUpMessage")}
	})
	renderFrame(t, "empty-snackbar", 1100, 760, LightPalette(), func(a *App) {
		a.Delete(1)
		a.Delete(2)
		a.Form = Form{Name: "Reunião", Minutes: "30"}
	})
}
