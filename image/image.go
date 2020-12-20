//go:generate msgp
package quantdimage

import (
	"image"
	"image/color"
)

type ImageRequest struct {
	Image     []byte
	DX        int
	DY        int
	Speed     int
	MaxColors int
	Dither    float64
}

type ImageResponse struct {
	Image   []byte
	DX      int
	DY      int
	Palette []RGBA
}

type RGBA struct {
	R, G, B, A uint8
}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

func NewImageRequest(img image.Image, speed, maxColors int, dither float64) *ImageRequest {
	return &ImageRequest{
		Image:     GoImageToRgba32(img),
		DX:        img.Bounds().Dx(),
		DY:        img.Bounds().Dy(),
		Speed:     speed,
		MaxColors: maxColors,
		Dither:    dither,
	}
}

func (r *ImageResponse) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *ImageResponse) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.DX, r.DY)
}

func (r *ImageResponse) At(x, y int) color.Color {
	return r.Palette[r.ColorIndexAt(x, y)]
}

func (r *ImageResponse) ColorIndexAt(x, y int) uint8 {
	return r.Image[(x%r.DX)+(y*r.DX)]
}

func GoImageToRgba32(im image.Image) []byte {
	ret := make([]byte, im.Bounds().Dx()*im.Bounds().Dy()*4)

	p := 0

	for y := im.Bounds().Min.Y; y < im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Min.X; x < im.Bounds().Max.X; x++ {
			r16, g16, b16, a16 := im.At(x, y).RGBA() // Each value ranges within [0, 0xffff]

			ret[p+0] = uint8(r16 >> 8)
			ret[p+1] = uint8(g16 >> 8)
			ret[p+2] = uint8(b16 >> 8)
			ret[p+3] = uint8(a16 >> 8)
			p += 4
		}
	}

	return ret
}
