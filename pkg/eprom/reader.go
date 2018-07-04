// This package is for reading/decoding dumps of EPROMS used by the CPS2
// I know for sure it works for Vampire Savior, other games not supported
// This could be extended to parse mame's driver code?
// example https://github.com/mamedev/mame/blob/8179a84458204a5e767446fcf7d10f032a40fd0c/src/mame/drivers/cps2.cpp#L8578

package eprom

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

type pair struct {
	even, odd []byte
}

//Decode combines 4 seperated EPROM banks
func Decode(r []io.Reader) []byte {
	p := make([]pair, len(r))
	for i, v := range r {
		p[i] = parse(v)
	}

	return interleave(p)
}

// prepares
func parse(r io.Reader) pair {
	n := make([]byte, 2)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	p, _ := byteutils.Deinterleave(b, n, 2)
	return pair{p[0], p[1]}
}

// does the entire interleave process for the files
func interleave(p []pair) []byte {
	firstPass := interleavePairs(p, 2)
	secondPass := interleavePairs(firstPass, 64)
	final := byteutils.Interleave(1048576, secondPass[0].even, secondPass[0].odd)

	return final
}

// interleaves even and odd parts
func interleavePairs(p []pair, n int) (ip []pair) {
	for i := 0; i < len(p)/2; i++ {
		even := byteutils.Interleave(n, p[i*2].even, p[i*2+1].even)
		odd := byteutils.Interleave(n, p[i*2].odd, p[i*2+1].odd)
		ip = append(ip, pair{even, odd})
	}

	return ip
}
