package desktop

import (
	"github.com/gabrielluizsf/antui"
	"github.com/gabrielluizsf/antui/canvas"
)

// loadFaces asks for the system's interface font at the display's scale and
// derives every size from it. A machine without one keeps AntUI's built-in
// face, blown up to roughly the same sizes.
func (v *view) loadFaces() {
	if f, err := antui.UIFace(15, v.scale); err == nil {
		canvas.SetDefaultFace(f)
	}
	base := canvas.DefaultFace()
	v.body = base
	v.small = faceAt(base, 13*v.scale)
	v.title = faceAt(base, 20*v.scale)
	v.clock = faceAt(base, 54*v.scale)
}

func faceAt(base *canvas.Face, pixels float64) *canvas.Face {
	if base.Fixed() {
		n := int(pixels/float64(base.Height()) + 0.5)
		return canvas.BuiltinScaled(max(n, 1))
	}
	return base.AtSize(pixels)
}
