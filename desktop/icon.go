package desktop

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"

	"github.com/gabrielluizsf/antui/canvas"
)

// decodeIcon turns PNG bytes into the canvases a window icon is made of.
func decodeIcon(data []byte) ([]*canvas.Canvas, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	b := img.Bounds()
	cv, err := canvas.NewCanvas(b.Dx(), b.Dy())
	if err != nil {
		return nil, err
	}
	for y := range b.Dy() {
		for x := range b.Dx() {
			c := color.NRGBAModel.Convert(img.At(b.Min.X+x, b.Min.Y+y)).(color.NRGBA)
			cv.Pixels[y*cv.Stride+x] = canvas.Color(uint32(c.A)<<24 | uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B))
		}
	}
	return canvas.IconSet(cv), nil
}
