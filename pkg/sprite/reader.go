package sprite

import (
	"fmt"
	"io"
	"strconv"

	"github.com/goosechooser/cps2gfx/pkg/tile"
)

type reader interface {
	io.ReadSeeker
}

// Decoder for Sprite
type Decoder struct {
	r        reader
	tDecoder tile.Decoder
	BaseTile string
	// Width, Height int
	s Sprite
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
		d.s.Dx = w
	}
}

// Height is how tall a Sprite is (in Tiles)
func Height(h int) DecodeOption {
	return func(d *Decoder) {
		d.s.Dy = h
	}
}

// NewDecoder is a constructor ofc
func NewDecoder(r io.ReadSeeker, opts ...DecodeOption) *Decoder {
	d := &Decoder{r: r, tDecoder: *tile.NewDecoder(r), s: Sprite{}}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// decode handles seeking to the correct spot in the file
// and reading the data that makes up all the tiles in a Sprite.
// The actual reading/unpacking is handled by tDecoder
func (d *Decoder) decode() error {

	a, err := strconv.ParseInt(d.BaseTile, 0, 64)
	if err != nil {
		return err
	}

	_, err = d.seek(a*128, io.SeekStart)
	if err != nil {
		e := fmt.Errorf("%q: @ basetile %q", err, d.BaseTile)
		return e
	}

	d.s.Tiles, _ = d.decodeTiles(d.s.Dx, d.s.Dy)

	return nil
}

func (d *Decoder) seek(offset int64, whence int) (s int64, err error) {
	s, err = d.r.Seek(offset, whence)
	if err != nil {
		return 0, err
	}

	return s, nil
}

func (d *Decoder) decodeTiles(width, height int) (tiles []tile.Tile, err error) {
	tiles = make([]tile.Tile, width*height)
	//rowOffset represents the distance between rows that make up sprites
	rowOffset := int64(16 * 128)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			tiles[w+(h*width)], _ = d.decodeTile()
		}
		_, err = d.seek(rowOffset, io.SeekCurrent)
		if err != nil {
			e := fmt.Errorf("%q: @ basetile %q", err, d.BaseTile)
			return []tile.Tile{}, e
		}
	}
	return tiles, nil
}
func (d *Decoder) decodeTile() (tile.Tile, error) {
	return d.tDecoder.Decode()
}

// Decode reads a Sprite from the underlying stream
func (d *Decoder) Decode() (Sprite, error) {
	if err := d.decode(); err != nil {
		return Sprite{}, err
	}

	return d.s, nil
}
