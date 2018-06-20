package tile

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	packed, _ = hex.DecodeString("F1F2F3F4F1F2F3F4F1F2F3F4F1F2F3F4F1F2F3F4F1F2F3F4F1F2F3F4F1F2F3F4")
	pixels, _ = hex.DecodeString("0f0f0f0f000806050f0f0f0f000806050f0f0f0f000806050f0f0f0f000806050f0f0f0f000806050f0f0f0f000806050f0f0f0f000806050f0f0f0f00080605")
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


func TestUnpack16(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{packed, pixels},
	}
	for i, c := range cases {
		got := unpack16(c.input)
		t.Logf("case %d len got %d len ouput %d\n", i, len(got), len(c.output))
		buf := bytes.NewBuffer(got)
		buf2 := bytes.NewBuffer(c.output)
		row := make([]byte, 8)
		row2 := make([]byte, 8)
		i := 0
		for buf.Len() > 0 {
			buf.Read(row)
			buf2.Read(row2)
			if bytes.Equal(row, row2) != true {
				t.Errorf("Row %d\n", i)
				t.Errorf("%x, want %x", row, row2)
				// readRows(row, c.output, t)
			}
			i++
		}
	}
}

func TestUnpack32(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{packed, pixels},
	}
	for i, c := range cases {
		m := make([]byte, len(c.input))
		copy(m, c.input)

		for i := 0; i < 3; i++ {
			m = append(m, c.input...)
		}

		expected := make([]byte, len(c.output))
		for i := 0; i < 4; i++ {
			expected = append(c.output, c.output...)
		}

		got := unpack32(m)
		t.Logf("case %d len got %d len ouput %d\n", i, len(got), len(expected))
		buf := bytes.NewBuffer(got)
		buf2 := bytes.NewBuffer(expected)
		row := make([]byte, 8)
		row2 := make([]byte, 8)
		i := 0
		for buf.Len() > 0 {
			buf.Read(row)
			buf2.Read(row2)
			if bytes.Equal(row, row2) != true {
				t.Errorf("Row %d\n", i)
				t.Errorf("%x, want %x", row, row2)
				// readRows(row, c.output, t)
			}
			i++
		}
	}
}

func readRows(got, expected []byte, t *testing.T) {
	bif := bytes.NewBuffer(got)
	buf := bytes.NewBuffer(expected)
	ok := make([]byte, 8)
	ko := make([]byte, 8)

	for bif.Len() > 0 {
		bif.Read(ok)
		buf.Read(ko)
		t.Errorf("\ngot is %x\n expected is %x\n", ok, ko)
	}
}

func TestPack16(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{packed, pixels},
	}
	for i, c := range cases {
		u := unpack16(c.input)
		got := pack16(u)
		if bytes.Equal(got, c.input) != true {
			t.Errorf("Case %d\n", i)
			t.Errorf("PackT(%q) == %q, want %q\n", u, got, c.input)
			t.Errorf("Len of got %d\n", len(got))
		}
	}
}

func TestPack32(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{packed, pixels},
	}
	for i, c := range cases {
		m := make([]byte, len(c.input))
		copy(m, c.input)

		for i := 0; i < 3; i++ {
			m = append(m, c.input...)
		}
		expected := make([]byte, len(m))
		copy(expected, m)

		u := unpack32(m)
		got := pack32(u)
		if bytes.Equal(got, expected) != true {
			t.Errorf("Case %d\n", i)
			t.Errorf("PackT(%q) == %q, want %q\n", u, got, expected)
			t.Errorf("Len of got %d\n", len(got))
		}
	}
}

func BenchmarkUnpack16(b *testing.B) {
	cases := []struct {
		input []byte
	}{
		{packed},
	}
	for _, c := range cases {
		for n := 0; n < b.N; n++ {
			unpack16(c.input)
		}
	}
}

func BenchmarkUnpack32(b *testing.B) {
	cases := []struct {
		input []byte
	}{
		{packed},
	}
	for _, c := range cases {
		m := make([]byte, len(c.input))
		copy(m, c.input)

		for i := 0; i < 3; i++ {
			m = append(m, c.input...)
		}

		for n := 0; n < b.N; n++ {
			unpack32(m)
		}
	}
}

func BenchmarkPack16(b *testing.B) {
	cases := []struct {
		input []byte
	}{
		{packed},
	}
	for _, c := range cases {
		unpacked := unpack16(c.input)
		for n := 0; n < b.N; n++ {
			pack16(unpacked)
		}
	}
}

func BenchmarkPack32(b *testing.B) {
	cases := []struct {
		input []byte
	}{
		{packed},
	}
	for _, c := range cases {
		m := make([]byte, len(c.input))
		copy(m, c.input)

		for i := 0; i < 3; i++ {
			m = append(m, c.input...)
		}

		unpacked := unpack32(m)
		for n := 0; n < b.N; n++ {
			pack32(unpacked)
		}
	}
}
