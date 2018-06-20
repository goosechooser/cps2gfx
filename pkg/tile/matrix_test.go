package tile

import (
	"bytes"
	"reflect"
	"testing"
)

var (
	m8  = []uint8{1, 2, 3, 4, 1, 2, 3, 4}
	t8  = []uint8{0x55, 0x66, 0x88, 0, 0, 0, 0, 0}
	m16 = []uint16{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}
	t16 = []uint16{0x5555, 0x6666, 0x7878, 0x8080, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m32 = []uint32{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	t32 = []uint32{0x55555555, 0x66666666, 0x88888888, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

func printMatrix8(m []uint8, t *testing.T) {
	for i := range m {
		t.Logf("%08b\n", m[i])
	}
	t.Log("\n")
}

func printMatrix16(m []uint16, t *testing.T) {
	for i := range m {
		t.Logf("%016b\n", m[i])
	}
	t.Log("\n")
}

func printMatrix32(m []uint32, t *testing.T) {
	for i := range m {
		t.Logf("%032b\n", m[i])
	}
	t.Log("\n")
}

func TestTranspose8(t *testing.T) {
	cases := []struct {
		m          []uint8
		transposed []uint8
	}{
		{m8, t8},
	}
	for _, c := range cases {
		ref := make([]byte, len(m8))
		copy(ref, m8)
		printMatrix8(ref, t)
		transpose8(c.m)
		if bytes.Equal(c.m, c.transposed) != true {
			t.Errorf("transpose8(%x) == %x, want %x", ref, c.m, c.transposed)
			printMatrix8(c.m, t)
		}
	}
}

func TestTranspose16(t *testing.T) {
	cases := []struct {
		m          []uint16
		transposed []uint16
	}{
		{m16, t16},
	}
	for _, c := range cases {
		t.Logf("len of m %d, len of transposed %d\n", len(c.m), len(c.transposed))
		ref := make([]uint16, len(c.m))
		copy(ref, c.m)
		printMatrix16(ref, t)
		transpose16(c.m)
		if reflect.DeepEqual(c.m, c.transposed) != true {
			t.Errorf("transpose16(%x) == %x, want %x", ref, c.m, c.transposed)
			printMatrix16(c.m, t)
		}
	}
}

func TestTranspose32(t *testing.T) {
	cases := []struct {
		m          []uint32
		transposed []uint32
	}{
		{m32, t32},
	}
	for _, c := range cases {
		t.Logf("len of m %d, len of transposed %d\n", len(c.m), len(c.transposed))
		ref := make([]uint32, len(c.m))
		copy(ref, c.m)
		printMatrix32(ref, t)
		transpose32(c.m)
		if reflect.DeepEqual(c.m, c.transposed) != true {
			t.Errorf("transpose32(%x) == %x, want %x", ref, c.m, c.transposed)
			printMatrix32(c.m, t)
		}
	}
}

func BenchmarkTranspose8(b *testing.B) {
	cases := []struct {
		m []uint8
	}{
		{m8},
	}
	for _, c := range cases {
		for n := 0; n < b.N; n++ {
			transpose8(c.m)
		}
	}
}

func BenchmarkTranspose16(b *testing.B) {
	cases := []struct {
		m []uint16
	}{
		{m16},
	}
	for _, c := range cases {
		for n := 0; n < b.N; n++ {
			transpose16(c.m)
		}
	}
}

func BenchmarkTranspose32(b *testing.B) {
	cases := []struct {
		m []uint32
	}{
		{m32},
	}
	for _, c := range cases {
		for n := 0; n < b.N; n++ {
			transpose32(c.m)
		}
	}
}
