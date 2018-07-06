package eprom_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/eprom"
)

func TestEncode(t *testing.T) {
	files := []string{"testdata/mock.13", "testdata/mock.15", "testdata/mock.17", "testdata/mock.19"}
	expected := make([][]byte, len(files))

	for i, v := range files {
		expected[i] = open(v, t)
	}

	file, err := os.Open("testdata/mock.final")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	got := eprom.Encode(file)

	for i := range got {
		if bytes.Equal(got[i], expected[i]) != true {
			t.Errorf("Incorrect - file %d", i)
			t.Logf("len(got) %d len(expected) %d\n", len(got[i]), len(expected[i]))
		}
	}

}

func BenchmarkEncode(b *testing.B) {
	file, _ := os.Open("testdata/mock.final")
	defer file.Close()

	for i := 0; i < b.N; i++ {
		eprom.Encode(file)
	}

}

// pkg: github.com/goosechooser/cps2gfx/pkg/eprom
// BenchmarkEncode-8   	    5000	    201583 ns/op	 1141664 B/op	      67 allocs/op
