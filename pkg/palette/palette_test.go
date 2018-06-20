package palette

import (
	"image/color"
	"testing"
)

var palData = []string{"ff00", "ffda", "ffc9", "fd97", "fb75", "fee0", "fc00", "f842", "f620", "f345", "f468", "f68a", "f8ac", "fbce", "ffff", "f001"}
var expected = color.Palette{
		color.RGBA{0xF, 0x0, 0x0, 0xF},
		color.RGBA{0xF, 0xD, 0xA, 0xF},
	}
func TestParseString(t *testing.T) {
	got := parseStrings(palData[:2])

	if len(got) != len(expected) {
		t.Errorf("len(got): %d, len(expected): %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("parseString(%x) = %x, expected %x", palData[:2], got[i], expected[i])
		}
	}

}
