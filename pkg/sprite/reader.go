package sprite

import (
	"io"

	"github.com/goosechooser/cps2gfx/pkg/tile"
)

type reader interface {
	io.ReadSeeker
}

// Decoder for Sprite
type Decoder struct {
	tile.Decoder
	r             reader
	BaseTile      string
	Width, Height int
}

// DecodeOption just an alias for fun
type DecodeOption func(*Decoder)

// BaseTile is the memory address of the Tile at (0,0) in the Sprite
func BaseTile(bt string) DecodeOption {
	return func(d *Decoder) {
		d.BaseTile = bt
	}
}

// Width is how long a Sprite is (in Tiles)
func Width(w int) DecodeOption {
	return func(d *Decoder) {
		d.Width = w
	}
}

// Height is how tall a Sprite is (in Tiles)
func Height(h int) DecodeOption {
	return func(d *Decoder) {
		d.Height = h
	}
}

// NewDecoder is a constructor ofc
func NewDecoder(r io.ReadSeeker, opts ...DecodeOption) *Decoder {
	d := &Decoder{Decoder: *tile.NewDecoder(r), r: r}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// decode handles reading the data that makes up a Sprite.
// The actual reading/unpacking is handled by tile.Decoder
func (d *Decoder) decode() (b []byte, err error) {
	row := make([]tile.Tile, d.Width)

	for h := 0; h < d.Height; h++ {
		for w := 0; w < d.Width; w++ {
			row[w], err = d.Decoder.Decode()
			if err != nil {
				return b, err
			}
		}

		b = append(b, combineHorizontal(row, d.Width)...)
		if err = d.nextRow(); err != nil {
			return b, err
		}
	}
	return b, nil
}

// To horizontally stitch tiles together you have to structure the bytes as:
// [tile0-row0] [tile1-row0] ... [tileN-row0]
// ... ... ...
// [tile0-rowN] [tile1-rowN] ... [tileN-rowN]
// this could be accomplished by a variadic interleave u kno
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

func (d *Decoder) nextRow() (err error) {
	//rowOffset represents the distance between rows that make up sprites
	rowOffset := int64(16 * 128)
	_, err = d.r.Seek(rowOffset, io.SeekCurrent)
	if err != nil {
		return err
	}
	return nil
}

// Decode reads a Sprite from the underlying stream
func (d *Decoder) Decode() (Sprite, error) {
	data, err := d.decode()
	if err != nil {
		return Sprite{}, err
	}
	return Sprite{Dx: d.Width, Dy: d.Height, Data: data}, nil
}
