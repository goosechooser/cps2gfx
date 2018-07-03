package sprite

import (
	"io"

	"github.com/goosechooser/cps2gfx/pkg/tile"
	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

type writer interface {
	io.Writer
}

// Encoder for sprites
type Encoder struct {
	tile.Encoder
	w   writer
	err error
}

// splits sprite data into a collection of Tiles
func encode(b []byte, dx, dy int) []tile.Tile {
	tiles := make([]tile.Tile, 0, dx*dy)
	rowLength := dx * 256
	temp := make([]byte, rowLength)
	raw := make([][]byte, 0, dx*dy)

	for y := 0; y < dy; y++ {
		copy(temp, b[y*rowLength:])
		byteutils.Deinterleave(temp, 16, raw[y*dx:y*dx+dx]...)
	}

	for i := range raw {
		tiles = append(tiles, tile.Tile{Data: raw[i]})
	}

	return tiles
}

// Splits sprite into tiles and then uses the Tile encoder to write them all out
func (e *Encoder) writeSprite (s Sprite) {
	tiles := encode(s.Data, s.Dx, s.Dy)
	for _, t := range tiles {
		err := e.Encoder.Encode(t)
		if err != nil {
			e.err = err
		}
	}
}

// NewEncoder constructs a sprite encoder
func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{Encoder: *tile.NewEncoder(w)}
	return e
}

// Encode a sprite, packing it and then writing it to the underlying stream.
func (e *Encoder) Encode(s Sprite) (err error) {
	e.writeSprite(s)
	return e.err
}



