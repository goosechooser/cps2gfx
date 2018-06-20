package sprite

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Sprite is multiple 16x16 Tiles + Palette
type Sprite struct {
	// BaseTile is the address of the Tile located at (0,0)
	BaseTile string
	// Dx is the width of the Sprite, in Tiles
	// Dx is the height of the Sprite, in Tiles
	Dx, Dy int
	// The data that makes up the sprite
	Data []byte
}

// ToImage massages all the Tiles together and returns an Image.
func (s *Sprite) ToImage(p color.Palette) image.Image {
	// 16 is a magic number that should be explained/refactored
	dy := 16 * s.Dy
	dx := 16 * s.Dx

	r := image.Rect(0, 0, dx, dy)
	img := image.NewPaletted(r, p)
	img.Pix = s.Data
	img.Stride = dx
	return img
}

// String make it do
// since this is p much for debugging
// you should refactor to use dependency injection
// then move the function out of the package
func (s *Sprite) String() string {
	var b strings.Builder
	for i, v := range s.Data {
		if i%(16*s.Dx) == 0 {
			fmt.Fprintf(&b, "\n")
		}
		fmt.Fprintf(&b, "%d, ", v)
	}
	fmt.Fprintf(&b, "\n")
	return b.String()
}
