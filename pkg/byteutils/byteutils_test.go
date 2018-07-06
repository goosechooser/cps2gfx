package byteutils_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

// because while a slice of bytes.Reader is a collection of things that
// satisfy the io.Reader interface, it is not /a slice of io.Reader/
func why(b ...[]byte) []io.Reader {
	r := make([]io.Reader, len(b))
	for i := range b {
		r[i] = bytes.NewReader(b[i])
	}
	return r
}

func TestInterleave(t *testing.T) {
	cases := []struct {
		files       [][]byte
		interleaved []byte
		n           int
	}{
		{
			[][]byte{[]byte("ac"), []byte("bd")},
			[]byte("abcd"), 1,
		},
		{
			[][]byte{[]byte("ac"), []byte("bd")},
			[]byte("acbd"), 2,
		},
		{
			[][]byte{[]byte("ac"), []byte("bd"), []byte("ef")},
			[]byte("abecdf"), 1,
		},
		{
			[][]byte{[]byte("acde"), []byte("bdef"), []byte("efgh")},
			[]byte("acbdefdeefgh"), 2,
		},
	}

	for _, c := range cases {
		br := make([]*bytes.Reader, len(c.files))
		for i := range br {
			br[i] = bytes.NewReader(c.files[i])
		}
		got := byteutils.Interleave(c.n, why(c.files...)...)
		if bytes.Equal(got, c.interleaved) != true {
			t.Errorf("Interleave == %q, want %q", got, c.interleaved)
		}
	}
}

func TestDeinterleave(t *testing.T) {
	cases := []struct {
		buf []byte
		n, o      int
		debuf  [][]byte
	}{
		{[]byte("abcd"), 1, 2, [][]byte{
			[]byte("ac"),
			[]byte("bd"),
		}},
		{[]byte("abcd"), 2, 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
		}},
		{[]byte("abcdef"), 2, 3, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
			[]byte("ef"),
		}},
	}
	for _, c := range cases {
		got, _ := byteutils.Deinterleave(bytes.NewReader(c.buf), c.n, c.o)

		for i := range got {
			if bytes.Equal(got[i], c.debuf[i]) != true {
				t.Errorf("Deinterleave(%q) == %q, want %q", c.buf, got[i], c.debuf[i])
			}
		}
	}
}

func TestInterleaveThenDeinterleave(t *testing.T) {
	cases := []struct {
		file1, file2, interleaved []byte
		n                         int
	}{
		{[]byte("ac"), []byte("bd"), []byte("abcd"), 1},
		{[]byte("ac"), []byte("bd"), []byte("acbd"), 2},
	}
	for _, c := range cases {
		got := byteutils.Interleave(c.n, why(c.file1, c.file2)...)
		final, _ := byteutils.Deinterleave(bytes.NewReader(got), c.n, 2)

		if bytes.Equal(final[0], c.file1) != true {
			t.Errorf("Did not interleave then deinterleave properly, got %q want %q", final[0], c.file1)
		}

		if bytes.Equal(final[1], c.file2) != true {
			t.Errorf("Did not interleave then deinterleave properly, got %q want %q", final[1], c.file2)
		}
	}
}

func BenchmarkInterleave(b *testing.B) {
	cases := []struct {
		file1, file2, interleaved []byte
		n                         int
	}{
		{[]byte("ac"), []byte("bd"), []byte("abcd"), 1},
		{[]byte("ac"), []byte("bd"), []byte("acbd"), 2},
	}
	for _, c := range cases {
		nice := why(c.file1, c.file2)
		for i := 0; i < b.N; i++ {
			byteutils.Interleave(c.n, nice...)
		}
	}
}

func BenchmarkDeinterleave(b *testing.B) {
	cases := []struct {
		buf []byte
		n, o      int
		debuf  [][]byte
	}{
		{[]byte("abcd"), 2, 1, [][]byte{
			[]byte("ac"),
			[]byte("bd"),
		}},
		{[]byte("abcd"), 2, 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
		}},
		{[]byte("abcdef"), 3, 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
			[]byte("ef"),
		}},
	}
	for _, c := range cases {
		buf := bytes.NewReader(c.buf)
		for i := 0; i < b.N; i++ {
			byteutils.Deinterleave(buf, c.n, c.o)
		}
	}
}

// pkg: github.com/goosechooser/cps2gfx/pkg/byteutils
// BenchmarkDeinterleave-8   	  300000	      3891 ns/op	   20615 B/op	      11 allocs/op