package sprite

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/goosechooser/cps2gfx/pkg/tile"
)

// Sprite is multiple 16x16 Tiles + Palette
type Sprite struct {
	// BaseTile is the address of the Tile located at (0,0)
	BaseTile int
	// Dx is the width of the Sprite, in Tiles
	// Dx is the height of the Sprite, in Tiles
	Dx, Dy int
	// The Tiles that make up the sprite
	Tiles []tile.Tile
}

// ToImage massages all the Tiles together and returns an Image.
func (s *Sprite) ToImage(p color.Palette) image.Image {
	dy := s.Tiles[0].Size * s.Dy
	dx := s.Tiles[0].Size * s.Dx

	r := image.Rect(0, 0, dx, dy)
	img := image.NewPaletted(r, p)
	img.Pix = s.CombineTiles()
	img.Stride = dx
	return img
}

// CombineTiles concats the Sprite's individual Tiles
// into the appropiate 2D representation
func (s *Sprite) CombineTiles() []byte {
	return combineTiles(s.Tiles, s.Dx, s.Dy)
}

func combineTiles(tiles []tile.Tile, dx, dy int) (b []byte) {
	b = make([]byte, dx*dy*len(tiles[0].Data))
	rowLength := dx * len(tiles[0].Data)
	nRows := len(tiles) / dx

	for i := 0; i < nRows; i++ {
		row := combineHorizontal(tiles[i*dx:(i*dx)+dx], dx)
		copy(b[i*rowLength:i*rowLength+rowLength], row)
	}

	return b
}

// To horizontally stitch tiles together you have to structure the bytes as:
// [tile0-row0] [tile1-row0] ... [tileN-row0]
// ... ... ...
// [tile0-rowN] [tile1-rowN] ... [tileN-rowN]
func combineHorizontal(tiles []tile.Tile, dx int) (b []byte) {
	tSize := len(tiles[0].Data)
	rowLength := tiles[0].Size
	b = make([]byte, dx*tSize)

	for y := 0; y < rowLength; y++ {
		start := rowLength * y
		for x, t := range tiles {
			offset := x * rowLength
			copy(b[start*dx+offset:start*dx+offset+rowLength], t.Data[start:start+rowLength])
		}
	}

	return b
}

//String make it do
func (s *Sprite) String() string {
	var b strings.Builder
	data := combineTiles(s.Tiles, s.Dx, s.Dy)
	for i, v := range data {
		if i%(16*s.Dx) == 0 {
			fmt.Fprintf(&b, "\n")
		}
		fmt.Fprintf(&b, "%d, ", v)
	}
	fmt.Fprintf(&b, "\n")
	return b.String()
}
