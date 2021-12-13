package main

import (
	"bufio"
	"day13/internal/paper"
	"fmt"
	"os"
)

func main() {
	in, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	sheet, folds := in.Paper, in.Folds
	for n, f := range folds {
		if n == 1 {
			fmt.Println("Solve1:", len(sheet))
		}
		sheet = sheet.Fold(f)
	}
	fmt.Println(sheet)
}

type input struct {
	Paper paper.Paper
	Folds []paper.Fold
}

func readInput(p string) (input, error) {
	f, err := os.Open(p)
	if err != nil {
		return input{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ppr := make([]paper.Coord, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		point, err := paper.FromText(text)
		if err != nil {
			return input{}, err
		}

		ppr = append(ppr, point)
	}

	folds := make([]paper.Fold, 0)

	for scanner.Scan() {
		text := scanner.Text()
		fo, err := paper.FoldFromText(text)
		if err != nil {
			return input{}, err
		}
		folds = append(folds, fo)
	}

	input := input{
		Paper: ppr,
		Folds: folds,
	}
	return input, nil
}
