package sprite_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/sprite"
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

func TestCombineTiles(t *testing.T) {
	r := open("testdata/cooldata")
	d := tile.NewDecoder(r)
	tiles := make([]tile.Tile, 3)

	for i := range tiles {
		tiles[i], _ = d.Decode()

	}

	s := sprite.Sprite{0, 3, 1, tiles}

	if false {
		t.Error(s.String())
	}
}
