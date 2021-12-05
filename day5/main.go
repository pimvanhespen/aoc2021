package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}


type Field [][]int

func NewField(min, max Point) Field {
	cols := max.X - min.X
	rows := max.Y - min.Y

	field := make([][]int, rows+1)

	for n := range field {
		field[n] = make([]int, cols+1)
	}

	return field
}

func (f Field) String() string {
	s := ""
	for _, row := range f {
		for _, col := range row {
			if col == 0 {
				s += "."
			} else {
				s += strconv.Itoa(col)
			}
		}
		s += "\n"
	}
	return s
}

func (f Field) Copy() Field {
	field := make([][]int, len(f))
	for x := range field {
		field[x] = make([]int, len(f[x]))
		for y := range f[x] {
			field[x][y] = f[x][y]
		}
	}
	return field
}

func Max( a,b int) int {
	a, b = abs(a), abs(b)
	if a > b {
		return a
	}
	return b
}

func (f Field) Plot(l Line, allowDiagonals bool){
	dx := l.B.X - l.A.X
	dy := l.B.Y - l.A.Y

	if abs(dx) == abs(dy) {
		if ! allowDiagonals {
			return
		}
	} else if dx == 0 && dy != 0 {

	} else if dx != 0 && dy == 0 {

	} else {
		fmt.Printf("A: %+v; B: %+v; dx=%d, dy=%d", l.A, l.B, dx, dy)
		panic("wtf.exe")
	}

	stepX, stepY := 0, 0

	if dx < 0 {
		stepX = -1
	} else if dx > 0 {
		stepX = 1
	}

	if dy < 0 {
		stepY = -1
	} else if dy > 0 {
		stepY = 1
	}

	//fmt.Printf("A: %+v; B: %+v; dx=%d, sx=%d; dy=%d, sy=%d\n", l.A, l.B, dx, stepX, dy, stepY)

	max := Max(dx, dy)

	for i := 0; i <= max; i++ {
		x := l.A.X + i * stepX
		y := l.A.Y + i * stepY

		//fmt.Sprintf("X,Y = %d,%d", x, y)

		f[x][y]++
	}
}

type Line struct {
	A Point
	B Point
}

type Point struct {
	X, Y int
}

func TextToPoint(text string) Point {
	input := strings.Split(text, ",")
	sx, sy := input[0], input[1]
	x, _ := strconv.Atoi(sx)
	y, _ := strconv.Atoi(sy)

	return Point{x, y}
}

func readInput(reader io.Reader) (Field, []Line){
	scanner := bufio.NewScanner(reader)

	min, max := Point{}, Point{}

	lines :=  make([]Line, 0)

	for scanner.Scan() {
		text := scanner.Text()
		coords := strings.Split(text, " -> ")

		//fmt.Printf("%+v\n", coords)

		a := TextToPoint(coords[0])
		b := TextToPoint(coords[1])

		if a.X > max.X {
			max.X = a.X
		}
		if a.Y > max.Y {
			max.Y = a.Y
		}
		if b.X > max.X {
			max.X = b.X
		}
		if b.Y > max.Y {
			max.Y = b.Y
		}

		line := Line{a, b}
		lines = append(lines, line)
	}

	field := NewField(min, max)

	return field, lines
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func (f Field) CountMinHits(min int) int {
	hits := 0
	for _, row := range f {
		for _, col := range row {
			if col >= min {
				hits++
			}
		}
	}
	return hits
}

func exec(inpath string) error {
	f, err := os.Open(inpath)
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer f.Close()

	field, lines := readInput(f)

	field2 := field.Copy()

	for _, line := range lines {
		field.Plot(line, false)
		field2.Plot(line, true)
		//fmt.Println(field)
	}

	solve1 := field.CountMinHits(2)

	solve2 := field2.CountMinHits(2)

	fmt.Printf("Solve1: %d\n", solve1)
	fmt.Printf("Solve2: %d\n", solve2)

	// do stuff

	return nil
}