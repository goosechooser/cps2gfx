package tile

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Tile represents a collection of pixels (16x16)
// Need to seperate image stuff from tile stuffs
type Tile struct {
	Address int64
	Data    []byte
	Size    int
}


// DefaultPalette is the default enjoy
var DefaultPalette = color.Palette{
	color.RGBA{0x55, 0x55, 0x00, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x4c, 0x4c, 0x99, 0xff},
	color.RGBA{0x49, 0x49, 0xdd, 0xff},
	color.RGBA{0x4c, 0x99, 0x00, 0xff},
	color.RGBA{0x4c, 0x99, 0x4c, 0xff},
	color.RGBA{0x4c, 0x99, 0x99, 0xff},
	color.RGBA{0x49, 0x93, 0xdd, 0xff},
	color.RGBA{0x49, 0xdd, 0x00, 0xff},
	color.RGBA{0x49, 0xdd, 0x49, 0xff},
	color.RGBA{0x49, 0xdd, 0x93, 0xff},
	color.RGBA{0x49, 0xdd, 0xdd, 0xff},
	color.RGBA{0x4f, 0xee, 0xee, 0xff},
	color.RGBA{0x66, 0x00, 0x00, 0xff},
	color.RGBA{0x66, 0x00, 0x66, 0xff},
	color.RGBA{0x55, 0x00, 0xaa, 0xff},
	color.RGBA{0x4f, 0x00, 0xee, 0xff},
}

// ToImage turns a Tile into an Image
func (t *Tile) ToImage(p color.Palette) image.Image {
	pal := image.NewPaletted(image.Rect(0, 0, t.Size, t.Size), p)
	pal.Pix = t.Data
	pal.Stride = t.Size
	return pal
}

//String make it do
func (t *Tile) String() string {
	var b strings.Builder
	for i, v := range t.Data {
		if i%t.Size == 0 {
			fmt.Fprintf(&b, "\n")
		}
		fmt.Fprintf(&b, "%d, ", v)
	}
	fmt.Fprintf(&b, "\n")
	return b.String()
}
