package eprom_test

import (
	"bytes"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/eprom"
)

func TestDeinterleave(t *testing.T) {
	files := []string{"testdata/mock.13", "testdata/mock.15", "testdata/mock.17", "testdata/mock.19"}
	expected := make([][]byte, len(files))

	for i, v := range files {
		expected[i] = open(v, t)
	}

	b := open("testdata/mock.final", t)
	got := eprom.Deinterleave(b)

	for i := range got {
		if bytes.Equal(got[i], expected[i]) != true {
			t.Errorf("Incorrect - file %d", i)
		}
	}

}
