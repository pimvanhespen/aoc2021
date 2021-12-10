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
type Fish int

func NewFish(start int) Fish {
	return Fish(start)
}

func (f Fish) GetBirths(limit int) []int {
	births := make([]int, 0)
	next := int(f) + 9
	for i := next; i <= limit; i += 7 {
		births = append(births, i)
	}
	return births
}

func (f Fish) GetValue(day int) int {
	// first day
	days := day - int(f)

	if days < 9 {
		return 8 - days
	}

	days -= 9

	return 6 - days % 7
}

func AddBirths(births []int, m map[int]int){
	for _, birth := range births {
		if _, ok := m[birth]; ! ok {
			m[birth] = 0
		}
		m[birth]++
	}
}

type Stack map[int]int

func (m Stack) Add(addition, index int) {
	if _, ok := m[index]; ! ok {
		m[index] = 0
	}
	m[index] += addition
}

func printFish(day int, fishes []Fish){
	s := fmt.Sprintf("%2d (%d): ", day, len(fishes))
	for _, fish := range fishes {
		s += strconv.Itoa(fish.GetValue(day)) + " "
	}
	fmt.Println(s)
}

func SolveVerbose(input []int, limit int) int {
	stack := make(Stack)
	fishes := make([]Fish, 0, len(input))

	for _, num := range input {
		start := num - 8
		f := NewFish(start)
		fishes = append(fishes, f)

		for _, date := range f.GetBirths(limit) {
			stack.Add(1, date)
		}
	}

	for day := 0; day <= limit; day++ {
		births, ok := stack[day]
		if ! ok {
			printFish(day, fishes)
			continue
		}

		for birth := 0; birth < births; birth++ {
			f := NewFish(day)
			fishes = append(fishes, f)

			bts := f.GetBirths(limit)
			AddBirths(bts, stack)
		}
		printFish(day, fishes)
	}

	return len(fishes)
}

func Solve2(input []int, limit int) int {
	stack := make(Stack)

	total := 0

	for _, num := range input {
		start := num - 8
		f := NewFish(start)
		total++

		for _, date := range f.GetBirths(limit) {
			stack.Add(1, date)
		}
	}

	for days := 0; days <= limit; days++ {
		births, ok := stack[days]
		if ! ok {
			continue
		}

		total += births
		f := NewFish(days)

		for _, date := range f.GetBirths(limit) {
			stack.Add(births, date)
		}
	}

	return total
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
		//fmt.Printf("Day %3d; total: %10d\n", i, total)
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

	SolveVerbose(numbers, 18)

	solution1a := Solve1(numbers, 80)
	fmt.Println(solution1a)
	solution1b := Solve2(numbers, 80)
	fmt.Println(solution1b)


	solution2 := Solve2(numbers, 256)
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
