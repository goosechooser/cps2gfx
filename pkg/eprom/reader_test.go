package eprom_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/eprom"
)

func open(f string, t *testing.T) []byte {
	file, err := os.Open(f)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		t.Error(err)
	}

	return b
}

func TestDecode(t *testing.T) {
	files := []string{"testdata/mock.13", "testdata/mock.15", "testdata/mock.17", "testdata/mock.19"}

	b := make([]io.Reader, len(files))
	for i, v := range files {
		file, err := os.Open(v)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		b[i] = file
	}

	result := eprom.Decode(b)
	expected := open("testdata/mock.final", t)

	if bytes.Equal(result, expected) != true {
		t.Errorf("Incorrect")
	}
}
