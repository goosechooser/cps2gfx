package tile

import (
	"bufio"
	"encoding/binary"
	"io"
)

type writer interface {
	io.Writer
}

// Encoder woo!!
type Encoder struct {
	w   writer
	err error
}

func encode(b []byte) []byte {
	return pack32(b)
}

func pack32(b []byte) (t []byte) {
	b = reverse(b)
	row := make([]uint32, 32)

	for i := 0; i < 32; i++ {
		p := []byte{(b[i] << 4) | b[i+32], (b[i+64] << 4) | b[i+96], (b[i+128] << 4) | b[i+160], (b[i+192] << 4) | b[i+224]}
		row[i] = binary.LittleEndian.Uint32(p)
	}

	transpose32(row)

	pix := make([]byte, 4)
	t = make([]byte, 128)

	for i := range row {
		binary.LittleEndian.PutUint32(pix, row[i])
		t[i] = pix[0]
		t[i+32] = pix[1]
		t[i+64] = pix[2]
		t[i+96] = pix[3]
	}

	return t
}

// Pack16 uses transpose (currently for comparison with naive implementation)
func pack16(b []byte) (t []byte) {
	b = reverse(b)
	row := make([]uint16, 16)

	for i := 0; i < 16; i++ {
		p := []byte{(b[i] << 4) | b[i+8], (b[i+16] << 4) | b[i+32]}
		row[i] = binary.LittleEndian.Uint16(p)
	}

	transpose16(row)

	pix := make([]byte, 2)
	t = make([]byte, 32)

	for i := range row {
		binary.LittleEndian.PutUint16(pix, row[i])
		t[i] = pix[0]
		t[i+16] = pix[1]
	}

	return t
}

//Encode encodes the contents of a Tile and writes them to a stream
func (e *Encoder) writeTile(t Tile) {
	b := encode(t.Data)
	_, e.err = e.w.Write(b)
}

// NewEncoder is a constructor
func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{}
	e.initWriter(w)
	return e
}

func (e *Encoder) initWriter(w io.Writer) {
	if ww, ok := w.(writer); ok {
		e.w = ww
	} else {
		e.w = bufio.NewWriter(w)
	}
}

// Encode packs the Tile data and writes it to the underlying stream.
func (e *Encoder) Encode(t Tile) (err error) {
	e.writeTile(t)

	return e.err
}
