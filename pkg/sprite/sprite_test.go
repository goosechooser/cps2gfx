package sprite_test

import (
	"io"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/goosechooser/cps2gfx/pkg/sprite"
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
	opts := []sprite.DecodeOption{
		sprite.Width(1),
		sprite.Height(1),
	}
	expected := make([]byte, 256)
	r.Read(expected)
	r.Seek(0, io.SeekStart)

	d := sprite.NewDecoder(r, opts...)

	s, err := d.Decode()
	if err != nil {
		log.Fatalf("Decoding error, %v", err)
	}

	b := new(bytes.Buffer)
	e := sprite.NewEncoder(b)

	e.Encode(s)
	t.Logf("who %x\n", b.Bytes())

	if bytes.Equal(b.Bytes(), expected) != true {
		t.Logf("who %x\n", expected)
		t.Logf("who %x\n", b.Bytes())
	}
}

func BenchmarkDecode(b *testing.B) {
	r := open("testdata/cooldata")
	opts := []sprite.DecodeOption{
		sprite.Width(1),
		sprite.Height(1),
	}
	expected := make([]byte, 256)
	r.Read(expected)
	r.Seek(0, io.SeekStart)

	d := sprite.NewDecoder(r, opts...)

	for i := 0; i < b.N; i++ {
		d.Decode()
	}
}

func BenchmarkEncode(b *testing.B) {
	r := open("testdata/cooldata")
	opts := []sprite.DecodeOption{
		sprite.Width(1),
		sprite.Height(1),
	}
	expected := make([]byte, 256)
	r.Read(expected)
	r.Seek(0, io.SeekStart)

	d := sprite.NewDecoder(r, opts...)

	s, err := d.Decode()
	if err != nil {
		log.Fatalf("Decoding error, %v", err)
	}

	buf := new(bytes.Buffer)
	e := sprite.NewEncoder(buf)

	for i := 0; i < b.N; i++ {
		e.Encode(s)
	}
}
