package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

type Letters []rune

func (l Letters) Contains(r rune) bool {
	for _, c := range l {
		if c == r {
			return true
		}
	}
	return false
}

func (l Letters) ContainsAll(other Letters) bool {
	for _, letter := range other {
		if !l.Contains(letter) {
			return false
		}
	}
	return true
}

func (l Letters) String() string {
	return string(l)
}

// Difference returns the unique letters contained in l that are not in other
// example 'abc' | 'ab' returns 'c'
func (l Letters) Difference(other Letters) Letters {
	result := make([]rune, 0)
	for _, c := range l {
		if other.Contains(c) {
			continue
		}
		result = append(result, c)
	}
	return result
}

type Digit struct {
	Top, TopLeft, TopRight, Center, BottomLeft, BottomRight, Bottom rune
}

func (d Digit) GetLetters() Letters {
	all := []rune{d.TopLeft, d.Top, d.TopRight, d.Center, d.BottomLeft, d.Bottom, d.BottomRight}
	for i := len(all) - 1; i >= 0; i-- {
		if all[i] != 0 {
			continue
		}
		all = append(all[:i], all[i+1:]...)
	}
	return all
}

func (d Digit) ParseNumber(l Letters) int {
	length := len(l)

	if length == 2 {
		return 1
	}
	if length == 3 {
		return 7
	}
	if length == 4 {
		return 4
	}
	if length == 7 {
		return 8
	}

	if length == 5 {
		if l.Contains(d.TopRight) && l.Contains(d.BottomRight) {
			return 3
		} else if l.Contains(d.TopRight) {
			return 2
		} else {
			return 5
		}
	}

	if !l.Contains(d.Center) {
		return 0
	}

	if l.Contains(d.TopRight) {
		return 9
	}

	return 6
}

func (d Digit) String() string {
	p := func(r rune) string {
		if r == 0 {
			return " "
		}
		return string(r)
	}

	s := fmt.Sprintf(" %3s \n", strings.Repeat(string(d.Top), 3))
	s += strings.Repeat(fmt.Sprintf("%1s%3s%1s\n", string(d.TopLeft), strings.Repeat(" ", 3), string(d.TopRight)), 2)
	s += fmt.Sprintf(" %3s \n", strings.Repeat(string(d.Center), 3))
	s += strings.Repeat(fmt.Sprintf("%1s%3s%1s\n", p(d.BottomLeft), strings.Repeat(" ", 3), p(d.BottomRight)), 2)
	s += fmt.Sprintf(" %3s \n", strings.Repeat(string(d.Bottom), 3))
	return s
}

type Puzzle struct {
	Reference []Letters
	Payload   []Letters
}

func (p Puzzle) GetOneFourSevenEight() (Letters, Letters, Letters, Letters) {
	var one, four, seven, eight Letters
	for _, l := range p.Reference {
		if len(l) == 2 {
			one = l
		}
		if len(l) == 3 {
			seven = l
		}
		if len(l) == 4 {
			four = l
		}
		if len(l) == 7 {
			eight = l
		}
	}
	return one, four, seven, eight
}

func (p Puzzle) GetInput(size int, mustContain Letters) (Letters, bool) {
	for _, letters := range p.Reference {
		if len(letters) != size {
			continue
		}

		if letters.ContainsAll(mustContain) {
			//fmt.Printf("%s >= %s\n", letters, mustContain)
			return letters, true
		}
	}
	return nil, false
}

func (p Puzzle) Solve() (Solution, error) {
	solution := Digit{}
	one, four, seven, eight := p.GetOneFourSevenEight()

	// TOP
	topCandidate := seven.Difference(one)
	if len(topCandidate) != 1 {
		return Solution{}, errors.New("top must have only one candidate")
	}
	solution.Top = topCandidate[0]

	three, ok := p.GetInput(5, seven)
	if !ok {
		return Solution{}, errors.New("missing three")
	}

	nine, ok := p.GetInput(6, three)
	if !ok {
		return Solution{}, errors.New("missing nine")
	}

	// TOP LEFT
	topLeftCandidate := nine.Difference(three)
	if len(topLeftCandidate) != 1 {
		return Solution{}, errors.New("topLeft must have only one candidate")
	}
	solution.TopLeft = topLeftCandidate[0]

	// CENTER
	centerCandidate := four.Difference(topLeftCandidate).Difference(one)
	if len(centerCandidate) != 1 {
		return Solution{}, errors.New("centerCandidate must have only one candidate")
	}
	solution.Center = centerCandidate[0]

	// BOTTOM
	centerAndBottom := three.Difference(seven)
	bottomCandidate := centerAndBottom.Difference(centerCandidate)
	if len(bottomCandidate) != 1 {
		return Solution{}, errors.New("bottomCandidate must have only one candidate")
	}
	solution.Bottom = bottomCandidate[0]

	// BOTTOM RIGHT
	five, ok := p.GetInput(5, solution.GetLetters())
	if !ok {
		return Solution{}, errors.New("missing five")
	}

	bottomRightCandidate := five.Difference(solution.GetLetters())
	if len(bottomRightCandidate) != 1 {
		return Solution{}, fmt.Errorf("bottomRightCandidate must have only one candidate\ngot: %+v\nfrom: [%+v] DIFF [%+v]\n%s\n", bottomRightCandidate, five, solution.GetLetters(), solution)
	}
	solution.BottomRight = bottomRightCandidate[0]

	// TOP RIGHT
	topRightCandidate := one.Difference(solution.GetLetters())
	if len(topRightCandidate) != 1 {
		return Solution{}, errors.New("topRightCandidate must have only one candidate")
	}
	solution.TopRight = topRightCandidate[0]

	// BOTTOM LEFT
	bottomLeftCandidate := eight.Difference(solution.GetLetters())
	if len(bottomLeftCandidate) != 1 {
		return Solution{}, errors.New("bottomLeftCandidate must have only one candidate")
	}
	solution.BottomLeft = bottomLeftCandidate[0]

	return Solution{
		Puzzle: p,
		Digit:  solution,
	}, nil
}

type Solution struct {
	Puzzle Puzzle
	Digit  Digit
}

func (s Solution) GetDigits() []int {
	var result []int
	for _, payload := range s.Puzzle.Payload {
		result = append(result, s.Digit.ParseNumber(payload))
	}
	return result
}

func (s Solution) GetNumber() int {
	length := len(s.Puzzle.Payload)
	sum := 0
	for n, digit := range s.Puzzle.Payload {
		multiplier := int(math.Pow10(length - n - 1))
		sum += multiplier * s.Digit.ParseNumber(digit)
	}
	return sum
}

func solve1(s []Solution) int {
	count := 0
	for _, solution := range s {
		for _, digit := range solution.GetDigits() {
			if digit == 1 || digit == 4 || digit == 7 || digit == 8 {
				count++
			}
		}
	}
	return count
}

func solve2(solutions []Solution) int {
	total := 0
	for _, solution := range solutions {
		total += solution.GetNumber()
	}
	return total
}

func main() {
	puzzles, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	var solutions []Solution

	for _, puzzle := range puzzles {
		solution, err := puzzle.Solve()
		if err != nil {
			panic(err)
		}
		solutions = append(solutions, solution)
	}

	solution1 := solve1(solutions)
	fmt.Printf("Solution1: %d\n", solution1)

	solution2 := solve2(solutions)
	fmt.Printf("Solution2: %d\n", solution2)
}

func stringToDigits(input string) []Letters {
	split := strings.Split(input, " ")
	all := make([]Letters, 0, len(split))
	for _, s := range split {
		all = append(all, []rune(s))
	}
	return all
}

func readInput(path string) ([]Puzzle, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	puzzles := make([]Puzzle, 0)
	for scanner.Scan() {
		line := scanner.Text()
		BeforeAfter := strings.Split(line, " | ")
		before, after := BeforeAfter[0], BeforeAfter[1]
		p := Puzzle{
			Reference: stringToDigits(before),
			Payload:   stringToDigits(after),
		}
		puzzles = append(puzzles, p)
	}

	return puzzles, nil
}
