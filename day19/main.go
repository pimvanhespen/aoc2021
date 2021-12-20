
package main

import (
	"day19/internal/input"
	"fmt"
	"path"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func NewPoint(x, y, z int) Point {
	return Point{x, y, z}
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

type Scanner struct {
	Name int
	Position Point
	Readings []Point
}

func NewScanner(name int, readings []Point) Scanner {
	return Scanner{
		Name:     name,
		Position: Point{},
		Readings: readings,
	}
}

func (s Scanner) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("--- scanner %d ---\n", s.Name))
	for _, r := range s.Readings {
		sb.WriteString(fmt.Sprintf("%s\n", r))
	}
	return sb.String()
}

func (s Scanner) SimilarTo(other Scanner) bool {
	return false
}

func main(){
	files := []string{
		"facings.txt",
		//"demo.txt",
		//"input.txt",
	}

	for _, file := range files {
		fpath := path.Join("inputs", file)
		fmt.Println("File: ", fpath)
		if err := exec(fpath); err != nil {
			panic(fmt.Errorf("%s: %v", fpath, err))
		}
	}
}

func solve1() int {
	return -1
}

func solve2() int {
	return -1
}

func readInput(in string) ([]Scanner, error){
	groups, err := input.ReadLinesGrouped(in)
	if err != nil {
		return nil, err
	}

	scanners := make([]Scanner, 0, len(groups))
	for _, group := range groups {
		readings := make([]Point, 0, len(group[1:]))
		for _, text := range group[1:] {
			var x, y, z int
			_, err := fmt.Sscanf(text, "%d,%d,%d", &x, &y, &z)
			if err != nil {
				return nil, err
			}
			reading := Point{x, y, z}
			readings = append(readings, reading)
		}

		var name int
		_, err := fmt.Sscanf(group[0], "--- scanner %d ---", &name)
		if err != nil {
			return nil, err
		}

		scanner := NewScanner(name, readings)
		scanners = append(scanners, scanner)
	}

	return scanners, nil
}

func exec(in string) error {
	scanners, err := readInput(in)
	if err != nil {
		return err
	}

	base := scanners[0]
	for _, s := range scanners {
		if base.SimilarTo(s) {
			fmt.Println("Success!")
		} else {
			fmt.Println("FAIL!")
		}
	}

	fmt.Println("solve1:", solve1())
	fmt.Println("solve2:", solve2())

	return nil
}