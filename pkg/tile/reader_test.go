package tile_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/tile"
)

func open(f string) *bytes.Reader {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewReader(b)
}

func expectedResults() (e []byte) {
	e = make([]byte, 256)
	for i := range e {
		e[i] = 0x0F
	}
	return e
}

func TestNewDecoder(t *testing.T) {
	r := open("testdata/cooldata")
	d := tile.NewDecoder(r, tile.Dimensions(16))

	if d.Dimensions != 16 {
		t.Errorf("Dimensions should be %d, got %d\n", 16, d.Dimensions)
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		fname    string
		off      int64
		expected []byte
	}{
		{"testdata/cooldata", 0, expectedResults()},
	}
	for _, c := range cases {
		r := open(c.fname)
		d := tile.NewDecoder(r)
		r.Seek(c.off, io.SeekStart)
		got, _ := d.Decode()

		if bytes.Equal(got.Data, c.expected) != true {
			t.Errorf("DecodedAt(%d) incorrectly\n", c.off)
			t.Errorf("Got %x\nExpected %x\n", got.Data, c.expected)
			// TURN THIS OUTPUT INTO RESULTS WOO
			t.Log(got.String())
		}
	}
}
