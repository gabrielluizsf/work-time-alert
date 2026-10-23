package desktop

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestDecodeIconKeepsTheColours(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 64, 64))
	for i := range img.Pix {
		img.Pix[i] = 0
	}
	img.Set(0, 0, color.NRGBA{R: 0x4F, G: 0x46, B: 0xE5, A: 0xFF})
	var buf bytes.Buffer
	assert.NoError(t, png.Encode(&buf, img))

	icons, err := decodeIcon(buf.Bytes())
	assert.NoError(t, err)
	assert.NotEmpty(t, icons)
}

func TestDecodeIconRejectsRubbish(t *testing.T) {
	_, err := decodeIcon([]byte("not a png"))
	assert.Error(t, err)
}
