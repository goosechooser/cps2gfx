// Package byteutils is for interleaving (combining) and deinterleaving (splitting) byte slices.
// Right now there is only support for interleaving 2 slices, or deinterleaving into 2 slices.
// This can be extended for more options (I've done it before), but I don't personally have a need for it.

package byteutils

import (
	"bytes"
)

// Interleave combines two byte slices.
// n is the number of bytes to interleave by.
func Interleave(buf1 []byte, buf2 []byte, n int) (ibuf []byte) {
	b1 := bytes.NewBuffer(buf1)
	b2 := bytes.NewBuffer(buf2)

	for b1.Len() > 0 {
		ibuf = append(ibuf, b1.Next(n)...)
		ibuf = append(ibuf, b2.Next(n)...)
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
