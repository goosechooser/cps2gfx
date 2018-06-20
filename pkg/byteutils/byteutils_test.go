package byteutils_test

import (
	"bytes"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/byteutils"
)

func TestInterleave(t *testing.T) {
	cases := []struct {
		file1, file2, interleaved []byte
		n                         int
	}{
		{[]byte("ac"), []byte("bd"), []byte("abcd"), 1},
		{[]byte("ac"), []byte("bd"), []byte("acbd"), 2},
	}
	for _, c := range cases {
		got := byteutils.Interleave(c.file1, c.file2, c.n)
		if bytes.Equal(got, c.interleaved) != true {
			t.Errorf("Interleave(%q, %q) == %q, want %q", c.file1, c.file2, got, c.interleaved)
		}
	}
}

func TestDeinterleave(t *testing.T) {
	cases := []struct {
		buf   []byte
		n     int
		debuf [][]byte
	}{
		{[]byte("abcd"), 1, [][]byte{
			[]byte("ac"),
			[]byte("bd"),
		}},
		{[]byte("abcd"), 2, [][]byte{
			[]byte("ab"),
			[]byte("cd"),
		}},
	}
	for _, c := range cases {
		got := byteutils.Deinterleave(c.buf, c.n)
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
		got := byteutils.Interleave(c.file1, c.file2, c.n)
		final := byteutils.Deinterleave(got, c.n)

		if bytes.Equal(final[0], c.file1) != true {
			t.Errorf("Did not interleave then deinterleave properly, got %q want %q", final[0], c.file1)
		}

		if bytes.Equal(final[1], c.file2) != true {
			t.Errorf("Did not interleave then deinterleave properly, got %q want %q", final[1], c.file2)
		}
	}
}
