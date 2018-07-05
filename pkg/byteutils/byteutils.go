// Package byteutils is for interleaving (combining) and deinterleaving (splitting) byte slices.
// Right now there is only support for interleaving 2 slices, or deinterleaving into 2 slices.
// This can be extended for more options (I've done it before), but I don't personally have a need for it.
package byteutils

import (
	"bytes"
	"io"
)

// Interleave combines multiple readers by alternating every len(b) bytes
func Interleave(n int, r...io.Reader) (ibuf []byte) {
	b := make([]byte, n)
	ibuf = make([]byte, 0, len(r) * len(b))
	bufs := make([]bytes.Buffer, len(r))

	for i := range r {
		bufs[i].ReadFrom(r[i])
	}

	for bufs[0].Len() > 0 {
		for i := range bufs {
			bufs[i].Read(b)
			ibuf = append(ibuf, b...)
		}
	}

	return ibuf
}

// Deinterleave seperates a stream into o number of slices.
// the size of n will determine the number of bytes to deinterleave by
func Deinterleave(r io.Reader, n int, o int) (debuf [][]byte, err error) {
	b := make([]byte, n)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	debuf = make([][]byte, o)

	for i := range debuf {
		debuf[i] = make([]byte, 0, buf.Len()/o)
	}

	for buf.Len() > 0 {
		for i := range debuf {
			_, err = buf.Read(b)
			if err != nil {
				break
			}
			debuf[i] = append(debuf[i], b...)
		}
	}

	return debuf, err
}
