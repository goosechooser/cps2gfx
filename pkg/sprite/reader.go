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
	Width, Height int
	// td wraps r
	td            tile.Decoder
	// r p much controls seeking in a consistent manner
	r             reader
	//rowOffset represents the distance between rows that make up sprites
	rowOffset int64
}

// DecodeOption just an alias for fun
type DecodeOption func(*Decoder)

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

//RowOffset is how many bytes between 'rows' in a spritesheet
func RowOffset(r int64) DecodeOption {
	return func(d *Decoder) {
		d.rowOffset = r
	}
}

// NewDecoder is a constructor ofc
func NewDecoder(r io.ReadSeeker, opts ...DecodeOption) *Decoder {
	d := &Decoder{
		td: *tile.NewDecoder(r),
		r: r,
	}

	for _, opt := range opts {
		opt(d)
	}
	return d
}

// decode handles reading the data that makes up a Sprite.
// The actual reading/unpacking is handled by the tile.Decoder
func (d *Decoder) decode() (b []byte, err error) {
	const tSizeUnpacked = 256

	rowLength := d.Width * tSizeUnpacked
	row := make([]tile.Tile, d.Width)
	b = make([]byte, rowLength*d.Height)

	for h := 0; h < d.Height; h++ {
		for w := 0; w < d.Width; w++ {
			row[w], err = d.td.Decode()
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
func (d *Decoder) Decode(opts ...DecodeOption) (Sprite, error) {
	for _, opt := range opts {
		opt(d)
	}

	data, err := d.decode()
	if err != nil {
		return Sprite{}, err
	}

	return Sprite{Dx: d.Width, Dy: d.Height, Data: data}, nil
}
