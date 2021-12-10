package main

import (
	"bufio"
	"day9/internal/field"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}

func exec(path string) error {
	field, err := readInputFromFile(path)
	if err != nil {
		return err
	}


	fmt.Printf("Solution1: %4d\n", solve1(field))

	fmt.Println("Solution2:", solve2(field))

	return nil
}

func readInputFromFile(path string) (field.Field, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		nums := make([]int, 0)
		for _, char := range strings.Split(text, ""){
			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		rows = append(rows, nums)
	}
	return rows, nil
}

func solve1(f field.Field) int {
	lowPoints := f.GetLowPoints()

	sum := len(lowPoints)
	for _, n := range lowPoints {
		sum += n
	}

	return sum
}

func solve2(f field.Field) int {
	sizes := f.FindBasinSizes()

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	sum := sizes[0] * sizes[1] * sizes[2]
	return sum
}