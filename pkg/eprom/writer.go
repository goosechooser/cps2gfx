// This package is for writing/encoding dumps of EPROMS used by the CPS2
// I know for sure it works for Vampire Savior, other games not supported
// This could be extended to parse mame's driver code?
// example https://github.com/mamedev/mame/blob/8179a84458204a5e767446fcf7d10f032a40fd0c/src/mame/drivers/cps2.cpp#L8578

package eprom

import (
	"bytes"
	"io"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

// Encode writes out b to the given streams
// func Encode(b []byte, w ...io.Writer) (err error) {
func Encode(r io.Reader, w ...io.Writer) (err error) {
	d := deinterleave(r)
	for i, wr := range w {
		_, err = wr.Write(d[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Deinterleave ... deinterleave
func deinterleave(r io.Reader) (d [][]byte) {
	n := make([]byte, 1048576)
	firstPass, _ := byteutils.Deinterleave(r, n, 2)

	n = make([]byte, 64)
	secondPass := make([][]byte, 0, len(firstPass)*2)
	for i := range firstPass {
		r = bytes.NewReader(firstPass[i])
		s, _ := byteutils.Deinterleave(r, n, 2)
		secondPass = append(secondPass, s...)
	}

	n = make([]byte, 2)
	final := make([][]byte, 0, len(secondPass)*2)
	for i := range secondPass {
		r = bytes.NewReader(secondPass[i])
		f, _ := byteutils.Deinterleave(r, n, 2)
		final = append(final, f...)
	}

	for i := 0; i < 4; i++ {
		d = append(d, byteutils.Interleave(2, final[i], final[i+4]))
	}

	return d
}
