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

func BenchmarkDecode(b *testing.B) {
	files := []string{"testdata/mock.13", "testdata/mock.15", "testdata/mock.17", "testdata/mock.19"}

	r := make([]io.Reader, len(files))
	for i, v := range files {
		file, _ := os.Open(v)
		defer file.Close()
		r[i] = file
	}

	for i := 0; i < b.N; i++ {
		eprom.Decode(r)
	}
}

// pkg: github.com/goosechooser/cps2gfx/pkg/eprom
// BenchmarkDecode-8   	    2000	    606662 ns/op	 3468455 B/op	      81 allocs/op
// PASS
// ok  	github.com/goosechooser/cps2gfx/pkg/eprom	4.849s
