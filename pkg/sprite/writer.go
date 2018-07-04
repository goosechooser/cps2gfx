package sprite

import (
	"io"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
	"github.com/goosechooser/cps2gfx/pkg/tile"
)

// Encoder for sprites
type Encoder struct {
	te  tile.Encoder
	err error
}

// splits sprite data into a collection of Tiles
func encode(b []byte, dx, dy int) []tile.Tile {
	const tSizeUnpacked = 256

	rowLength := dx * tSizeUnpacked
	tiles := make([]tile.Tile, 0, dx*dy)
	raw := make([][]byte, 0, dx*dy)
	row := make([]byte, 16)
	for y := 0; y < dy; y++ {
		r, _ := byteutils.Deinterleave(b[y*rowLength:], row, dx)
		raw = append(raw, r...)
	}

	for i := range raw {
		tiles = append(tiles, tile.Tile{Data: raw[i]})
	}

	return tiles
}

// Splits sprite into tiles and then uses the Tile encoder to write them all out
func (e *Encoder) writeSprite(s Sprite) {
	tiles := encode(s.Data, s.Dx, s.Dy)
	for _, t := range tiles {
		err := e.te.Encode(t)
		if err != nil {
			e.err = err
			return
		}
	}
}

// NewEncoder constructs a sprite encoder
func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{te: *tile.NewEncoder(w)}
	return e
}

// Encode a sprite, packing it and then writing it to the underlying stream.
func (e *Encoder) Encode(s Sprite) (err error) {
	e.writeSprite(s)
	return e.err
}
