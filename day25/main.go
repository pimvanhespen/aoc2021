package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func readInput() [][]byte {
	f, err := os.Open("day25/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return bytes.Split(b, []byte{'\n'})
}

func main() {

	fmt.Println(solve1(readInput()))
}

func solveVerbose(max int, input [][]byte) int {
	s := seafloor(input)

	var count int

	for {
		count++

		if s.step() == 0 {
			return count
		}

		fmt.Printf("Step %d\n", count)
		fmt.Println(s)

		if count >= max {
			return count
		}
	}
}

const (
	Empty = '.'
	Right = '>'
	Down  = 'v'
)

type seafloor [][]byte

type pos struct {
	X, Y int
}

func (p pos) right() pos {
	return pos{p.X + 1, p.Y}
}

func (p pos) down() pos {
	return pos{p.X, p.Y + 1}
}

func (s seafloor) is(p pos, typ byte) bool {
	return s.get(p) == typ
}

func (s seafloor) get(p pos) byte {
	return s[p.Y][p.X]
}

func (s seafloor) set(p pos, v byte) {
	s[p.Y][p.X] = v
}

func (s seafloor) right(p pos) pos {
	if p.X+1 < len(s[0]) {
		return pos{p.X + 1, p.Y}
	}
	return pos{0, p.Y}
}

func (s seafloor) down(p pos) pos {
	if p.Y+1 < len(s) {
		return pos{p.X, p.Y + 1}
	}
	return pos{p.X, 0}
}

func (s seafloor) step() int {
	var count int
	var moves [][2]pos

	width, height := len(s[0]), len(s)

	for i := 0; i < (width * height); i++ {
		x, y := i%width, i/width

		p := pos{x, y}

		if !s.is(p, Right) {
			continue
		}

		next := s.right(p)

		if s.is(next, Empty) {
			moves = append(moves, [2]pos{p, next})
		}
	}

	count += len(moves)

	for _, move := range moves {
		s.set(move[0], Empty)
		s.set(move[1], Right)
	}

	moves = moves[:0]

	for i := 0; i < (width * height); i++ {
		x, y := i%width, i/width

		p := pos{x, y}

		if !s.is(p, Down) {
			continue
		}

		next := s.down(p)

		if s.is(next, Empty) {
			moves = append(moves, [2]pos{p, next})
		}
	}

	count += len(moves)

	for _, move := range moves {
		s.set(move[0], Empty)
		s.set(move[1], Down)
	}

	return count
}

func (s seafloor) String() string {
	var b bytes.Buffer
	for _, row := range s {
		b.Write(row)
		b.WriteByte('\n')
	}
	return b.String()
}

func solve1(input [][]byte) int {
	s := seafloor(input)

	var count int
	for {
		count++

		if s.step() == 0 {
			return count
		}

	}
}
