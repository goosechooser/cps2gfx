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
	leng := len(b) * len(b[0])
	ibuf = make([]byte, 0, leng)
	nInterleaves := len(b[0]) / n

	for i := 0; i < nInterleaves; i++ {
		for _, buf := range b {
			ibuf = append(ibuf, buf[i*n:i*n+n]...)
		}
	}

	return ibuf
}

// Deinterleave seperates one slice into o number of slices.
// the size of n will determine the number of bytes to deinterleave by
func Deinterleave(b, n []byte, o int) (debuf [][]byte, err error) {
	buf := bytes.NewBuffer(b)
	debuf = make([][]byte, o)

	for i := range debuf {
		debuf[i] = make([]byte, 0, buf.Len()/o)
	}

	for buf.Len() > 0 {
		for i := range debuf {
			_, err = buf.Read(n)
			if err != nil {
				break
			}
			debuf[i] = append(debuf[i], n...)
		}
	}

	return debuf, err
}
