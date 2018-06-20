// This package is for creating palettes used by the image package.

package palette

import (
	"image/color"
	"log"
	"strconv"
)

func parseString(s string) color.RGBA {
	i, err := strconv.ParseInt("0x" + s, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	return color.RGBA{uint8(i&0x0F00 >> 8), uint8(i&0x00F0 >> 4), uint8(i&0x000F), uint8(i >> 12)}
}

func parseStrings(s []string) (p color.Palette) {
	for _, v := range s {
		p = append(p, parseString(v))
	}
	return p
}
