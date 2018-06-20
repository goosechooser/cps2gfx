package tile

const bitWidth8 = uint8(8)
const bitWidth16 = uint16(16)
const bitWidth32 = uint32(32)

func reverse(numbers []byte) []byte {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// transpose8 performs a bit matrix transpose on 8 uint8s
func transpose8(m []uint8) {
	width := bitWidth8
	mask := ^uint8(0)

	for width != 1 {
		width = width >> 1
		mask = mask ^ (mask >> width)
		swap8(m, width, mask)
	}
}

// Swap blocks within the bit matrix.
func swap8(m []uint8, width uint8, mask uint8) {
	var x, y *uint8

	for outer := uint8(0); outer < bitWidth8/(width*2); outer++ {
		for inner := uint8(0); inner < width; inner++ {
			x = &m[(inner)+(outer*width*2)]
			y = &m[(inner+width)+(outer*width*2)]

			*x = ((*y << width) & mask) ^ *x
			*y = ((*x & mask) >> width) ^ *y
			*x = ((*y << width) & mask) ^ *x
		}
	}
}

// transpose16 performs a bit matrix transpose on 16 uint16s
func transpose16(m []uint16) {
	width := bitWidth16
	mask := ^uint16(0)

	for width != 1 {
		width = width >> 1
		mask = mask ^ (mask >> width)
		swap16(m, width, mask)
	}
}

func swap16(m []uint16, width uint16, mask uint16) {
	var x, y *uint16
	var outer, inner uint16

	for outer = 0; outer < bitWidth16/(width*2); outer++ {
		for inner = 0; inner < width; inner++ {
			x = &m[(inner)+(outer*width*2)]
			y = &m[(inner+width)+(outer*width*2)]

			*x = ((*y << width) & mask) ^ *x
			*y = ((*x & mask) >> width) ^ *y
			*x = ((*y << width) & mask) ^ *x
		}
	}
}

// transpose32 performs a bit matrix transpose on 32 uint32s
func transpose32(m []uint32) {
	width := bitWidth32
	mask := ^uint32(0)

	for width != 1 {
		width = width >> 1
		mask = mask ^ (mask >> width)
		swap32(m, width, mask)
	}
}

func swap32(m []uint32, width uint32, mask uint32) {
	var x, y *uint32

	for outer := uint32(0); outer < bitWidth32/(width*2); outer++ {
		for inner := uint32(0); inner < width; inner++ {
			x = &m[(inner)+(outer*width*2)]
			y = &m[(inner+width)+(outer*width*2)]

			*x = ((*y << width) & mask) ^ *x
			*y = ((*x & mask) >> width) ^ *y
			*x = ((*y << width) & mask) ^ *x
		}
	}
}
