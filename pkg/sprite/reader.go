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
	//rowOffset represents the distance between rows that make up sprites
	rowOffset int64
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
	// probably make this a DecodeOption later?
	d.rowOffset = int64(d.Decoder.Dimensions * 128)

	for _, opt := range opts {
		opt(d)
	}
	return d
}

// decode handles reading the data that makes up a Sprite.
// The actual reading/unpacking is handled by the tile.Decoder
func (d *Decoder) decode() (b []byte, err error) {
	row := make([]tile.Tile, d.Width)
	rowLength := d.Width * 256
	b = make([]byte, rowLength*d.Height)

	for h := 0; h < d.Height; h++ {
		for w := 0; w < d.Width; w++ {
			row[w], err = d.Decoder.Decode()
			if err != nil {
				return b, err
			}
		}

		combineHorizontal(row, b[h*rowLength:])
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
func combineHorizontal(tiles []tile.Tile, b []byte) {
	rowLength := tiles[0].Size

	for y := 0; y < rowLength; y++ {
		start := rowLength * y
		for x, t := range tiles {
			offset := x * rowLength
			copy(b[start*len(tiles)+offset:], t.Data[start:start+rowLength])
		}
	}
}

func (d *Decoder) nextRow() (err error) {
	_, err = d.r.Seek(d.rowOffset, io.SeekCurrent)
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
