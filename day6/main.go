package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main(){
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}


// Fish is useful for debugging
type Fish struct {
	Start int
}

func NewFish(start int) Fish {
	return Fish{
		Start: start,
	}
}

func (f Fish) GetBirths(limit int) []int {
	births := make([]int, 0)
	next := f.Start + 9
	for i := next; i <= limit; i += 7 {
		births = append(births, i)
	}
	return births
}

func (f Fish) GetValue(day int) int {
	// first day
	days := day - f.Start

	if days < 9 {
		return 8 - days
	}

	days -= 9

	return 6 - days % 7
}

func AddBirths(births []int, m map[int]int){
	for _, birth := range births {
		if _, exists := m[birth]; ! exists {
			m[birth] = 0
		}
		m[birth]++
	}
}


func Solve1(input []int, limit int) int {
	stack := make(map[int]int)

	total := 0

	for _, num := range input {
		start := num - 8
		f := NewFish(start)
		total++

		births := f.GetBirths(limit)
		AddBirths(births, stack)
	}

	// exec each day, add fish birth
	for i := 0; i <= limit; i++ {
		fmt.Printf("Day %3d; total: %10d\n", i, total)
		births, ok := stack[i]
		if ! ok {
			continue
		}

		for birth := 0; birth < births; birth++ {
			total++
			f := NewFish(i)

			bts := f.GetBirths(limit)
			AddBirths(bts, stack)
		}
	}

	return total
}

func exec(path string) error {
	numbers, err := readInputFile(path)
	if err != nil {
		return err
	}

	const limit = 80

	solution1 := Solve1(numbers, limit)
	fmt.Println(solution1)

	solution2 := Solve1(numbers, 256)
	fmt.Println(solution2)

	return nil
}

func readInputFile(path string) (_ []int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("readInputFile: %v", err)
		}
	}()

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	b = bytes.TrimSpace(b)
	fields := bytes.Split(b, []byte(","))

	numbers := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(string(field))
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
