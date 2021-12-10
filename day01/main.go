package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputUrl = "https://adventofcode.com/2021/day/1/input"

func solve2(ints []int) (int, int){
	sum := func(ints []int) int {
		return ints[0]+ints[1]+ints[2]
	}

	increase, decrease := 0, 0

	var current int
	set := ints[:3]
	last := sum(set)
	for _, num := range ints[3:] {
		set = append(set[1:], num)
		current = sum(set)

		if current > last {
			increase++
		} else {
			decrease++
		}
		last = current
	}

	return increase, decrease
}

func solve1(ints []int) (int, int){
	increase, decrease := 0, 0

	last := ints[0]
	for _, num := range ints[1:] {
		if num > last {
			increase++
		} else {
			decrease++
		}
		last = num
	}

	return increase, decrease
}

func execute() ([]int, error) {

	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	ints := make([]int, 0)
	for scan.Scan() {
		text := scan.Text()
		n, e := strconv.Atoi(text)
		if e != nil {
			return nil, e
		}
		ints = append(ints, n)
	}

	return ints, nil
}

func main(){
	ints, err := execute()
	if err != nil {
		panic(err)
	}

	inc1, dec1 := solve1(ints)
	fmt.Println("Puzzle 1a")
	fmt.Printf("Up %4d - Down %4d\n", inc1, dec1)

	inc2, dec2 := solve2(ints)
	fmt.Println("Puzzle 1b")
	fmt.Printf("Up %4d - Down %4d\n", inc2, dec2)
}
