package main

import (
	"bufio"
	"day13/internal/coord"
	"day13/internal/fold"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	paper, folds := input.Extract()
	for n, f := range folds {
		if n == 1 {
			fmt.Println("Solve1:", len(paper))
		}
		paper = paper.Fold(f)
	}
	fmt.Println(paper)
}

type Paper []coord.Coord

func (p Paper) String() string {
	x, y := 0, 0
	for _, c := range p {
		if c.X > x {
			x = c.X
		}
		if c.Y > y {
			y = c.Y
		}
	}

	field := make([][]rune, y+1)
	for i := range field {
		row := make([]rune, x+1)
		for n := range row {
			row[n] = ' '
		}
		field[i] = row
	}

	for _, c := range p {
		field[c.Y][c.X] = '█'
	}
	sb := strings.Builder{}
	for _, row := range field {
		for _, r := range row {
			sb.WriteRune(r)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p Paper) Fold(f fold.Fold) Paper {
	m := map[string]coord.Coord{}
	for _, cIn := range p {
		cOut := cIn.Fold(f)
		m[cOut.String()] = cOut
	}

	paper := make([]coord.Coord, 0, len(m))
	for _, val := range m {
		paper = append(paper, val)
	}
	return paper
}

type Input struct {
	Paper Paper
	Folds []fold.Fold
}

func (i Input) Extract() (Paper, []fold.Fold) {
	return i.Paper, i.Folds
}

func readInput(p string) (Input, error) {
	f, err := os.Open(p)
	if err != nil {
		return Input{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	paper := make([]coord.Coord, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		point, err := coord.FromText(text)
		if err != nil {
			return Input{}, err
		}

		paper = append(paper, point)
	}

	folds := make([]fold.Fold, 0)

	for scanner.Scan() {
		text := scanner.Text()
		fo, err := fold.FromText(text)
		if err != nil {
			return Input{}, err
		}
		folds = append(folds, fo)
	}

	input := Input{
		Paper: paper,
		Folds: folds,
	}
	return input, nil
}
