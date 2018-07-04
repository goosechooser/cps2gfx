// This package is for writing/encoding dumps of EPROMS used by the CPS2
// I know for sure it works for Vampire Savior, other games not supported
// This could be extended to parse mame's driver code?
// example https://github.com/mamedev/mame/blob/8179a84458204a5e767446fcf7d10f032a40fd0c/src/mame/drivers/cps2.cpp#L8578

package eprom

import (
	"io"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

// Encode writes out b to the given streams
func Encode(b []byte, w ...io.Writer) (err error) {
	d := deinterleave(b)
	for i, wr := range w {
		_, err = wr.Write(d[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Deinterleave ... deinterleave
func deinterleave(b []byte) (d [][]byte) {
	firstPass := make([][]byte, 2)
	byteutils.Deinterleave(b, 1048576, firstPass...)

	secondPass := make([][]byte, len(firstPass)*2)
	for i, f := range firstPass {
		byteutils.Deinterleave(f, 64, secondPass[2*i:2*i+2]...)
	}

	final := make([][]byte, len(secondPass)*2)
	for i, s := range secondPass {
		byteutils.Deinterleave(s, 2, final[2*i:2*i+2]...)
	}

	for i := 0; i < 4; i++ {
		d = append(d, byteutils.Interleave(2, final[i], final[i+4]))
	}

	return d
}
