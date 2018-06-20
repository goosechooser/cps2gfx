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

type Pair struct {
	Even, Odd []byte
}

//Decode combines 4 seperated EPROM banks
func Decode(r []io.Reader) []byte {
	p := make([]Pair, len(r))
	for i, v := range r {
		p[i] = parse(v)
	}

	return Interleave(p)
}

func parse(r io.Reader) Pair {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	p := byteutils.Deinterleave(b, 2)
	return Pair{p[0], p[1]}
}

// InterleavePairs ... interleaves Pairs
func InterleavePairs(p []Pair, n int) (ip []Pair) {
	t := 0
	for i := 0; i < len(p)/2; i++ {
		t = i * 2
		even := byteutils.Interleave(p[t].Even, p[t+1].Even, n)
		odd := byteutils.Interleave(p[t].Odd, p[t+1].Odd, n)
		ip = append(ip, Pair{even, odd})
	}

	return ip
}

// Interleave but for FILES
func Interleave(p []Pair) []byte {
	firstPass := InterleavePairs(p, 2)
	secondPass := InterleavePairs(firstPass, 64)
	final := byteutils.Interleave(secondPass[0].Even, secondPass[0].Odd, 1048576)

	return final
}
