package tile

import (
	"bufio"
	"encoding/binary"
	"io"
)

const lower = 0x0F

type reader interface {
	io.Reader
}

// Decoder for Tiles
type Decoder struct {
	r          reader
	Dimensions int // Dimensions is the size of the Tile(s) to make
	buf        []byte
	t          Tile
}

// DecodeOption just an alias
type DecodeOption func(*Decoder)

// Dimensions are 8 for 8x8 Tile or 16 for 16x16 Tile
// Default value is 16.
func Dimensions(dim int) DecodeOption {
	return func(d *Decoder) {
		d.Dimensions = dim
	}
}

// NewDecoder is a constructor ofc
func NewDecoder(r io.Reader, opts ...DecodeOption) *Decoder {
	d := &Decoder{Dimensions: 16}
	d.initReader(r)

	for _, opt := range opts {
		opt(d)
	}

	d.initBuf()

	return d
}

func (d *Decoder) initBuf() {
	if d.Dimensions == 8 {
		d.buf = make([]byte, 32)
	}
	if d.Dimensions == 16 {
		d.buf = make([]byte, 128)
	}
}

func (d *Decoder) initReader(r io.Reader) {
	if rr, ok := r.(reader); ok {
		d.r = rr
	} else {
		d.r = bufio.NewReader(r)
	}
}

// Decode reads data and then returns it as a Tile.
func (d *Decoder) decode() error {
	_, err := d.r.Read(d.buf)
	if err != nil {
		return err
	}

	data := unpack32(d.buf)
	d.t = Tile{0, data, 16}

	return nil
}

// Decode reads a Tile from r
func (d *Decoder) Decode() (Tile, error) {
	if err := d.decode(); err != nil {
		return Tile{}, err
	}

	return d.t, nil
}

// unpack32 unpacks all 128bytes of a 16x16 tile 'at once'
func unpack32(b []byte) (t []byte) {
	row := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		//effectively: tile3, tile0, tile2, tile4
		p := []byte{b[i+64], b[i], b[i+96], b[i+32]}
		row[i] = binary.LittleEndian.Uint32(p)
	}

	transpose32(row)
	t = toPixel32(row)
	return reverse(t)
}

// Massages 32 uint32s into 64 'pixels'
func toPixel32(row []uint32) (t []byte) {
	pix := make([]byte, 4)
	t = make([]byte, 256)

	for i, v := range row[16:] {
		binary.LittleEndian.PutUint32(pix, v)
		t[i] = pix[3] >> 4
		t[i+16] = pix[3] & lower
		t[i+32] = pix[2] >> 4
		t[i+48] = pix[2] & lower
		t[i+64] = pix[1] >> 4
		t[i+80] = pix[1] & lower
		t[i+96] = pix[0] >> 4
		t[i+112] = pix[0] & lower
	}

	for i, v := range row[:16] {
		binary.LittleEndian.PutUint32(pix, v)
		t[i+128] = pix[3] >> 4
		t[i+144] = pix[3] & lower
		t[i+160] = pix[2] >> 4
		t[i+176] = pix[2] & lower
		t[i+192] = pix[1] >> 4
		t[i+208] = pix[1] & lower
		t[i+224] = pix[0] >> 4
		t[i+240] = pix[0] & lower
	}
	return t
}

// Unpack16 unpacks all 32bytes of a 8x8 tile 'at once'
func unpack16(b []byte) (t []byte) {
	row := make([]uint16, 16)
	for i := 0; i < 16; i++ {
		p := []byte{b[i], b[i+16]}
		row[i] = binary.LittleEndian.Uint16(p)
	}

	transpose16(row)
	t = toPixel16(row)
	return reverse(t)
}

func toPixel16(row []uint16) (t []byte) {
	lower := byte(15)
	pix := make([]byte, 2)
	t = make([]byte, 64)

	for i, v := range row[8:] {
		binary.LittleEndian.PutUint16(pix, v)
		t[i] = pix[1] >> 4
		t[i+8] = pix[1] & lower
		t[i+16] = pix[0] >> 4
		t[i+24] = pix[0] & lower
	}

	for i, v := range row[:8] {
		binary.LittleEndian.PutUint16(pix, v)
		t[i+32] = pix[1] >> 4
		t[i+40] = pix[1] & lower
		t[i+48] = pix[0] >> 4
		t[i+56] = pix[0] & lower
	}

	return t
}
