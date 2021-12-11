package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Octopuses [][]*Octopus

type Octopus struct {
	Value int
	Flashed bool
}

func NewOctopus(val int) *Octopus {
	return &Octopus{
		Value:   val,
		Flashed: false,
	}
}

func (o *Octopus) Start() {
	o.Flashed = false
	o.Value++
}

func (o *Octopus) Inc() {
	if ! o.Flashed {
		o.Value++
	}
}

func (o *Octopus) Flash() bool {
	if o.Value <= 9 {
		return false
	}
	o.Flashed = true
	o.Value = 0
	return true
}

var (
	exes = []int{-1, 0, 1, -1, 1, -1, 0, 1}
	whys = []int{-1, -1, -1, 0, 0, 1, 1, 1}
)

func (o Octopuses) FlashNeighbours(xIn,yIn int) int {
	flashes := 0
	for i := 0; i < 8; i++ {
		x, y := xIn + exes[i], yIn + whys[i]

		if x < 0 || y < 0 || y >= len(o) || x >= len(o[y]){
			continue
		}

		neighbour := o[y][x]
		neighbour.Inc()
		if neighbour.Flash() {
			flashes += 1 + o.FlashNeighbours(x, y)
		}
	}
	return flashes
}

func(o Octopuses) NextGen() int {
	flashes := 0

	for _, row := range o {
		for _, octopus := range row {
			octopus.Start()
		}
	}

	for y, row := range o {
		for x, octo := range row {
			if octo.Flash() {
				flashes += 1 + o.FlashNeighbours(x, y)
			}
		}
	}

	return flashes
}

func (o Octopuses) DeepCopy() Octopuses {
	rows := make([][]*Octopus, len(o))
	for y, row := range o {
		rowCopy := make([]*Octopus, len(row))
		for x, octo := range row {
			rowCopy[x] = NewOctopus(octo.Value)
		}
		rows[y] = rowCopy
	}
	return rows
}

func (o Octopuses) Count() int {
	count := 0
	for _, row := range o {
		count += len(row)
	}
	return count
}

func main(){
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}

func solve2(o Octopuses) int {
	flashes := 0
	target := o.Count()
	generations := 0
	for target != flashes {
		generations++
		flashes = o.NextGen()
	}
	return generations
}

func exec(path string) error {
	octo, err := readInput(path)
	if err != nil {
		return err
	}

	octo2 := octo.DeepCopy()

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += octo.NextGen()
	}
	fmt.Println("Solve1: ", totalFlashes)

	fmt.Println("Solve2: ", solve2(octo2))

	return nil
}

func readInput(path string) (Octopuses, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows := make([][]*Octopus, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		line := scanner.Text()
		octos := make([]*Octopus, 0, len(line))
		for _, r := range line {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			octos = append(octos, NewOctopus(i))
		}
		rows = append(rows, octos)
	}
	return rows, nil
}