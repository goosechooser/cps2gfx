// Package byteutils is for interleaving (combining) and deinterleaving (splitting) byte slices.
// Right now there is only support for interleaving 2 slices, or deinterleaving into 2 slices.
// This can be extended for more options (I've done it before), but I don't personally have a need for it.
package byteutils

import (
	"bytes"
)

// Interleave combines nbyte slices.
// n is the number of bytes to interleave by.
func Interleave(n int, b ...[]byte) (ibuf []byte) {
	ibuf = make([]byte, len(b)*len(b[0]))
	ilength := len(b) * n
	nInterleaves := len(b[0]) / n
	for i := 0; i < nInterleaves; i++ {
		start := i * ilength
		for j, buf := range b {
			offset := j * n
			copy(ibuf[start+offset:], buf[i:i+n])
		}
	}

	return ibuf
}

// Deinterleave seperates one slice into two slices. Could make this variadic tbh
// n is the number of bytes to deinterleave by.
func Deinterleave(buf []byte, n int) [][]byte {
	b := bytes.NewBuffer(buf)
	debuf := make([][]byte, 2)

	for b.Len() > 0 {
		debuf[0] = append(debuf[0], b.Next(n)...)
		debuf[1] = append(debuf[1], b.Next(n)...)
	}

	return debuf
}
