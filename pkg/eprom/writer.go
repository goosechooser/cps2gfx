// This package is for writing/encoding dumps of EPROMS used by the CPS2
// I know for sure it works for Vampire Savior, other games not supported
// This could be extended to parse mame's driver code?
// example https://github.com/mamedev/mame/blob/8179a84458204a5e767446fcf7d10f032a40fd0c/src/mame/drivers/cps2.cpp#L8578

package eprom

import (
	"io"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

// Encode do a trick
func Encode(w []io.Writer) {

}

// Deinterleave ... deinterleave
func Deinterleave(b []byte) (d [][]byte) {
	firstPass := byteutils.Deinterleave(b, 1048576)

	secondPass := [][]byte{}
	for _, i := range firstPass {
		secondPass = append(secondPass, byteutils.Deinterleave(i, 64)...)
	}

	final := [][]byte{}
	for _, i := range secondPass {
		final = append(final, byteutils.Deinterleave(i, 2)...)
	}

	for i := 0; i < 4; i++ {
		d = append(d, byteutils.Interleave(final[i], final[i+4], 2))
	}

	return d
}
