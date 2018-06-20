package tile_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/tile"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		fname    string
		expected []byte
	}{
		{"testdata/cooldata", make([]byte, 128)},
	}
	for _, c := range cases {
		r := open(c.fname)
		r.Read(c.expected)
		r.Seek(0, io.SeekStart)

		d := tile.NewDecoder(r, tile.Dimensions(16))
		decodedTile, _ := d.Decode()

		buf := bytes.NewBuffer([]byte{})
		e := tile.NewEncoder(buf)

		err := e.Encode(decodedTile)

		if err != nil {
			t.Errorf("Oh no %q\n", err)
		}

		got := make([]byte, 128)
		buf.Read(got)

		if bytes.Equal(got, c.expected) != true {
			t.Errorf("tile.Encode(tmpfile, %x) == %x, expected: %x\n", decodedTile.Data, got, c.expected)
		}
	}

}
