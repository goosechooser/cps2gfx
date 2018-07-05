package eprom_test

import (
	"bytes"
	"testing"
	"os"

	"github.com/goosechooser/cps2gfx/pkg/eprom"
)

func TestDeinterleave(t *testing.T) {
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

	got := make([]*bytes.Buffer, 4)
	for i := range got {
		got[i] = new(bytes.Buffer)
	}

	// There is some weirdness related to passing in slices of interfaces that I do not fully understand
	// so we just manually do it
	eprom.Encode(file, got[0], got[1], got[2], got[3])

	for i := range got {
		if bytes.Equal(got[i].Bytes(), expected[i]) != true {
			t.Errorf("Incorrect - file %d", i)
			t.Logf("len(got) %d len(expected) %d\n", len(got[i].Bytes()), len(expected[i]))
		}
	}

}
