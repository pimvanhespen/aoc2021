package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
)

func main() {
	files := []string{
		"demo.txt",
		"input.txt",
	}

	for _, file := range files {
		fpath := path.Join("inputs", file)
		fmt.Println("File: ", fpath)
		if err := exec(fpath); err != nil {
			panic(fmt.Errorf("%s: %v", fpath, err))
		}
	}
}

type Node struct {
	X, Y int
	Cost int
}

func NewNode(x, y, cost int) *Node {
	return &Node{x, y, cost}
}

type HeuristicFn func(*Node) int

func MhtnDist(tX, tY int) func(*Node) int {
	return func(node *Node) int {
		return tX - node.X + tY - node.Y
	}
}

type Field [][]*Node

func (f Field) NeighboursOf(n *Node) []*Node {
	var neighbours []*Node

	if n.Y-1 > 0 {
		neighbours = append(neighbours, f[n.Y-1][n.X])
	}

	if n.X+1 < len(f[n.Y]) {
		neighbours = append(neighbours, f[n.Y][n.X+1])
	}

	if n.Y+1 < len(f) {
		neighbours = append(neighbours, f[n.Y+1][n.X])
	}

	if n.X-1 > 0 {
		neighbours = append(neighbours, f[n.Y][n.X-1])
	}

	return neighbours
}

func (f Field) AStar(start, goal *Node, h HeuristicFn) []*Node {
	openSet := make(OpenSet)
	openSet.Add(start)

	cameFrom := make(map[*Node]*Node)

	gScore := make(map[*Node]int)
	fScore := make(map[*Node]int)
	for _, row := range f {
		for _, n := range row {
			gScore[n] = math.MaxInt
			fScore[n] = math.MaxInt
		}
	}

	gScore[start] = 0
	fScore[start] = h(start)

	for len(openSet) > 0 {
		current := openSet.MinCost(fScore)

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		openSet.Del(current)

		for _, neighbour := range f.NeighboursOf(current) {

			tentativeGScore := gScore[current] + neighbour.Cost

			if tentativeGScore < gScore[neighbour] {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + h(neighbour)
				openSet.Add(neighbour)
			}
		}
	}

	return nil
}

func (f Field) Multiply(size int) Field {
	multiplied := make(Field, len(f)*size)
	for y, row := range f {
		for i := 0; i < size; i++ {
			yOff := y + i*len(f)

			multiplied[yOff] = make([]*Node, len(f[y])*size)

			for x, node := range row {
				for j := 0; j < size; j++ {
					xOff := x + j*len(f[y])

					cost := node.Cost + i + j
					if cost > 9 {
						cost -= 9
					}

					multiplied[yOff][xOff] = NewNode(xOff, yOff, cost)
				}
			}
		}
	}
	return multiplied
}

type OpenSet map[*Node]struct{}

func (o OpenSet) Add(n *Node) {
	if _, ok := o[n]; !ok {
		o[n] = struct{}{}
	}
}

func (o OpenSet) Del(n *Node) {
	if _, ok := o[n]; ok {
		delete(o, n)
	}
}

func (o OpenSet) MinCost(d map[*Node]int) *Node {
	var lowest *Node
	min := math.MaxInt
	for n := range o {
		score := d[n]
		//fmt.Printf("MinCost: (%d,%d) = %d\n", n.X, n.Y, score)
		if score < min {
			min = score
			lowest = n
		}
	}
	return lowest
}

func reconstructPath(cameFrom map[*Node]*Node, current *Node) []*Node {
	totalPath := []*Node{current}
	var ok bool
	for {
		if current, ok = cameFrom[current]; !ok {
			return totalPath
		} else {
			totalPath = append([]*Node{current}, totalPath...)
		}
	}
}

func solve1(f Field) int {
	//f.Print(nil)
	//fmt.Println()

	start := f[0][0]
	eY := len(f) - 1
	eX := len(f[eY]) - 1
	end := f[eY][eX]

	result := f.AStar(start, end, MhtnDist(eX, eY))

	risk := -result[0].Cost
	for _, n := range result {
		risk += n.Cost
	}

	return risk
}

func solve2(f Field) int {
	multi := f.Multiply(5)

	return solve1(multi)
}

func readInput(in string) ([][]*Node, error) {
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}

	field := make([][]*Node, 0)

	scanner := bufio.NewScanner(f)
	lCount := 0
	for scanner.Scan() {
		text := scanner.Text()

		row := make([]*Node, len(text))
		for i, c := range text {
			row[i] = NewNode(i, lCount, int(c-'0'))
		}
		field = append(field, row)

		lCount++
	}

	return field, nil
}

func printField(field [][]*Node) {
	for _, row := range field {
		for _, n := range row {
			fmt.Print(n.Cost)
		}
		fmt.Println()
	}
}

func exec(in string) error {
	field, err := readInput(in)
	if err != nil {
		return err
	}

	printField(field)
	fmt.Println()

	fmt.Println("solve1:", solve1(field))
	fmt.Println("solve2:", solve2(field))

	return nil
}
