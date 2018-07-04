package byteutils_test

import (
	"bytes"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

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
		got := byteutils.Interleave(c.n, c.files...)
		if bytes.Equal(got, c.interleaved) != true {
			t.Errorf("Interleave == %q, want %q", got, c.interleaved)
		}
	}
}

func TestDeinterleave(t *testing.T) {
	cases := []struct {
		buf, n []byte
		o      int
		debuf  [][]byte
	}{
		{[]byte("abcd"), make([]byte, 1), 2, [][]byte{
			[]byte("ac"),
			[]byte("bd"),
		}},
		{[]byte("abcd"), make([]byte, 2), 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
		}},
		{[]byte("abcdef"), make([]byte, 2), 3, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
			[]byte("ef"),
		}},
	}
	for _, c := range cases {
		got, _ := byteutils.Deinterleave(c.buf, c.n, c.o)

		if bytes.Equal(got[0], c.debuf[0]) != true {
			t.Errorf("Deinterleave(%q) == %q, want %q", c.buf, got[0], c.debuf[0])
		}
		if bytes.Equal(got[1], c.debuf[1]) != true {
			t.Errorf("Deinterleave(%q) == %q, want %q", c.buf, got[1], c.debuf[1])
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
		got := byteutils.Interleave(c.n, c.file1, c.file2)
		n := make([]byte, c.n)
		final, _ := byteutils.Deinterleave(got, n, 2)

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
		for i := 0; i < b.N; i++ {
			byteutils.Interleave(c.n, c.file1, c.file2)
		}
	}
}

func BenchmarkDeinterleave(b *testing.B) {
	cases := []struct {
		buf, n []byte
		o      int
		debuf  [][]byte
	}{
		{[]byte("abcd"), make([]byte, 2), 1, [][]byte{
			[]byte("ac"),
			[]byte("bd"),
		}},
		{[]byte("abcd"), make([]byte, 2), 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
		}},
		{[]byte("abcdef"), make([]byte, 3), 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
			[]byte("ef"),
		}},
	}
	for _, c := range cases {
		for i := 0; i < b.N; i++ {
			byteutils.Deinterleave(c.buf, c.n, c.o)
		}
	}
}
