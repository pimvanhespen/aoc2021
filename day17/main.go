package main

import (
	"fmt"
	"path"
)

func main(){
	files := []string{
		"demo.txt",
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

func solve1(t TargetArea) int {
	return -1
}

func solve2() int {
	return -1
}

func exec(in string) error {

	t, err := readFile(in)
	if err != nil {
		return err
	}

	fmt.Println("solve1:", solve1(t))
	fmt.Println("solve2:", solve2())

	return nil
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type TargetArea struct {
	Point
	Width, Height int
}

func (t TargetArea) 

func (t TargetArea) PlotHighestCourse() Point {
	return Point{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func readFile(in string) (TargetArea, error) {
	var xa, xb int
	var ya, yb int

	_, err := fmt.Sscanf(in, "target area: x=%d..%d, y=-%d..-%d", &xa, &xb, &ya, &yb)
	if err != nil {
		return TargetArea{}, err
	}

	t := TargetArea{
		Point:  Point{
			X: min(xa,xb),
			Y: min(ya,yb),
		},
		Width:  abs(abs(xa)-abs(xb)),
		Height: abs(abs(ya)-abs(yb)),
	}
	return t, nil
}
