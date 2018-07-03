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
	leng := len(b)*len(b[0])
	ibuf = make([]byte, 0, leng)
	nInterleaves := len(b[0]) / n

	for i := 0; i < nInterleaves; i++ {
		for _, buf := range b {
			ibuf = append(ibuf, buf[i*n:i*n+n]...)
		}
	}

	return ibuf
}

// Deinterleave seperates one slice into two slices. Could make this variadic tbh
// n is the number of bytes to deinterleave by.
// debuf is a variable number of []bytes
// deinterleave data is read into these
func Deinterleave(buf []byte, n int, debuf ...[]byte) {
	b := bytes.NewBuffer(buf)

	for i := range debuf {
		debuf[i] = make([]byte, 0, b.Len()/len(debuf))
	}

	for b.Len() > 0 {
		for i := range(debuf) {
			debuf[i] = append(debuf[i], b.Next(n)...)
		}
	}
}
