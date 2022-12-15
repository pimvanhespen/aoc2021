package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Algorithm []byte

var infPixel byte

func (a Algorithm) Process(img Image, uneven bool) Image {

	// The value of infinite pixels in the pixel processing algorithm may change each iteration,
	// so we switch it between 0 and 511 (0b111111111) each iteration. As infinite pixels are
	// always 0 (0b000000000) or 511 (0b111111111)
	if uneven {
		// [. . .]
		// [. . .]  -> 0b000000000
		// [. . .]
		infPixel = a[0b000000000]
	} else {
		// [# # #]
		// [# # #]  -> 0b111111111
		// [# # #]
		infPixel = a[0b111111111]
	}

	const spacing = 1
	width := len(img[0]) + 2*spacing
	height := len(img) + 2*spacing

	out := NewImage(width, height, Low)

	for y := 0; y < len(out); y++ {
		for x := 0; x < len(out[y]); x++ {
			out[y][x] = a.ProcessPixel(img, x-spacing, y-spacing)
		}
	}
	return out
}

func (a Algorithm) ProcessPixel(img Image, x, y int) byte {

	var n uint
	var xOff, yOff int

	for i := 0; i < 9; i++ {
		xOff, yOff = i%3-1, i/3-1
		if !img.has(High, x+xOff, y+yOff) {
			continue
		}
		// Add the '1' bit on bit position i
		// 1 << 0 = 00000001 and corresponds to the bottom right pixel
		// 1 << 8 = 10000000 and corresponds to the top left pixel
		n |= 1 << (8 - i)
	}

	return a[n]
}

func NewImage(width int, height int, c byte) Image {
	out := make(Image, height)
	for y := range out {
		out[y] = bytes.Repeat([]byte{c}, width)
	}
	return out
}

type Image [][]byte

func (i Image) Count(high byte) int {
	var count int
	for y := 0; y < len(i); y++ {
		count += bytes.Count(i[y], []byte{high})
	}
	return count
}

func (i Image) String() string {
	var b bytes.Buffer

	for y := 0; y < len(i); y++ {
		b.Write(i[y])
		b.WriteByte('\n')
	}

	return b.String()
}

func (i Image) has(value byte, x int, y int) bool {
	if y < 0 || y >= len(i) || x < 0 || x >= len(i[y]) {
		// the value of infinite pixels swaps each iteration
		return value == infPixel
	}
	return i[y][x] == value
}

func parseInput() (Algorithm, Image) {
	f, err := os.Open("day20/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(b, []byte{'\n'})

	a := Algorithm(lines[0])
	i := Image(lines[2:])

	return a, i
}

const (
	High byte = '#'
	Low  byte = '.'
)

func main() {
	alg, img := parseInput()

	fmt.Println("Part 1:", solve(alg, img, 2))
	fmt.Println("Part 2:", solve(alg, img, 50))

}

func solve(alg Algorithm, img Image, times int) int {
	for i := 0; i < times; i++ {
		img = alg.Process(img, i%2 == 1)
	}

	return img.Count(High)
}
