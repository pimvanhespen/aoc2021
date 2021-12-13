package paper

import (
	"strings"
)

type Paper []Coord

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

func (p Paper) Fold(f Fold) Paper {
	m := map[string]Coord{}
	for _, cIn := range p {
		cOut := cIn.Fold(f)
		m[cOut.String()] = cOut
	}

	paper := make([]Coord, 0, len(m))
	for _, val := range m {
		paper = append(paper, val)
	}
	return paper
}
