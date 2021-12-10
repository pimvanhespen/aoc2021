package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type CostFunc func(int)int

type Move struct {
	Position int
	Cost int
}

func main() {
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}

func exec(path string) error {
	numbers, err := readInput(path)
	if err != nil {
		return err
	}

	c := Crabs{}

	for _, num := range numbers {
		c.Add(num)
	}

	solve1 := c.CalcCheapestMove(calcLinearCost)
	fmt.Printf("Move to %d is cheapest, costs %d\n", solve1.Position, solve1.Cost)

	solve2 := c.CalcCheapestMove(calcFancyCost)
	fmt.Printf("Move to %d is cheapest, costs %d\n", solve2.Position, solve2.Cost)

	return nil
}

func readInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	numStrings := strings.Split(string(b), ",")

	nums := make([]int, 0, len(numStrings))
	for _, str := range numStrings {
		num, e := strconv.Atoi(str)
		if e != nil {
			return nil, e
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func abs(in int )int {
	if in < 0 {
		return in * -1
	}
	return in
}

func calcLinearCost(dist int) int {
	return dist
}

func calcFancyCost(dist int) int {
	n := float64(dist)
	sum :=  ( n * (n + 1) ) / 2
	return int(math.Round(sum))
}

type Crabs map[int]int

func (c Crabs) Add(pos int) {
	if _, ok := c[pos]; ! ok {
		c[pos] = 0
	}
	c[pos]++
}

func (c Crabs) CalcMoveCost(target int, fn CostFunc) int {
	costs := 0
	for k, v := range c {
		dist := abs(k - target)
		cost := fn(dist)
		costs += cost * v
	}
	return costs
}

func (c Crabs) GetMinMax() (int, int) {
	lo, hi := 1 << 32, 0
	for k := range c {
		if k > hi {
			hi = k
		}
		if k < lo {
			lo = k
		}
	}
	return lo, hi
}

func (c Crabs) CalcCheapestMove(fn CostFunc) Move {
	cheapCost, cheapPosition := 1 << 32, 0

	min, max := c.GetMinMax()

	for pos := min; pos <= max; pos++ {
		cost := c.CalcMoveCost(pos, fn)
		//fmt.Printf("Move to %2d costs %5d\n", pos, cost)
		if cost < cheapCost{
			cheapCost = cost
			cheapPosition = pos
		}
	}

	return Move{
		Position: cheapPosition,
		Cost:     cheapCost,
	}
}
