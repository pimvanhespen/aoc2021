package field

import (
	"fmt"
)

var (
	// coord map for finding basins
	exes = []int{ 0, -1, 1, 0}
	whys = []int{-1,  0, 0, 1}
)

type Field [][]int

func (f Field) IsLowPoint(x, y int) bool {
	val := f[y][x]

	if ((y-1) >= 0) && val >= f[y-1][x] {
		return false
	}

	if ((x+1) < len(f[y])) && val >= f[y][x+1]{
		return false
	}

	if ((x-1) >= 0) && val >= f[y][x-1]{
		return false
	}

	if ((y+1) < len(f)) && val >= f[y+1][x]{
		return false
	}

	return true
}

func (f Field) GetLowPoints() []int {
	lowPoints := make([]int, 0)

	for y, row := range f {
		for x, num := range row {
			if f.IsLowPoint(x, y) {
				lowPoints = append(lowPoints, num)
			}
		}
	}

	return lowPoints
}

func (f Field) IsValidCoord(x,y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if y >= len(f) {
		return false
	}

	if x >= len(f[y]){
		return false
	}

	return true
}

func (f Field) MarkBasin(xPos, yPos int, checked *[][]bool, basins *[][]bool) int {

	(*checked)[yPos][xPos] = true
	(*basins)[yPos][xPos] = true
	size := 1

	for i := range exes {
		x, y := exes[i] + xPos, whys[i] + yPos

		if ! f.IsValidCoord(x, y) {
			continue
		}

		if (*checked)[y][x] {
			continue
		}

		if f[y][x] < 9 {
			size += f.MarkBasin(x, y, checked, basins)
		}
	}

	return size
}

func (f Field) FindBasinSizes() []int {
	checked := make([][]bool, 0, len(f))
	basins := make([][]bool, 0, len(f))
	for _, row := range f {
		checked = append(checked, make([]bool, len(row)))
		basins = append(basins, make([]bool, len(row)))
	}

	sizes := make([]int, 0)
	for y, row := range f {
		for x := range row {
			if checked[y][x] {
				continue
			}
			checked[y][x] = true

			if f[y][x] < 9 {
				size := f.MarkBasin(x, y, &checked, &basins)
				sizes = append(sizes, size)
			}
		}
	}

	printField(&basins, &checked)

	return sizes
}


// nice for debugging!
func printField(f , m*[][]bool){
	for y, row := range *f {
		for _, val := range row {
			if val {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" | ")
		for _, val := range (*m)[y] {
			if val {
				fmt.Print("M")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

