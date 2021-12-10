package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
}

type Stack []rune

func NewStack() *Stack {
	stack := Stack(make([]rune, 0))
	return &stack
}

func(s *Stack) Push(r rune){
	*s = append(*s, r)
}

func(s *Stack) Pop() (rune, bool) {
	if len(*s) < 1 {
		return 0, false
	}

	head := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return head, true
}

func(s *Stack) Peek() (rune, bool) {
	if len(*s) < 1 {
		return 0, false
	}

	return (*s)[len(*s)-1], true
}

func(s *Stack) HasPop() bool {
	return len(*s) > 0
}

type Status int

const (
	Invalid Status = iota
	Corrupted
	Incomplete
	Ok
)

type Evaluation struct {
	Status Status
	Index int
	Char rune
	Stack *Stack
}

func EvalLine(line string) Evaluation {
	stack := NewStack()
	for i, chr := range line {
		switch chr {
		case '(':
			stack.Push(')')
		case '[':
			stack.Push(']')
		case '{':
			stack.Push('}')
		case '<':
			stack.Push('>')
		case ')':
			fallthrough
		case '>':
			fallthrough
		case '}':
			fallthrough
		case ']':
			pop, ok := stack.Pop()
			if ! ok || pop != chr{
				return Evaluation{
					Status: Corrupted,
					Index:  i,
					Char:   chr,
					Stack: stack,
				}
			}
		}
	}

	if len(*stack) > 0 {
		return Evaluation{
			Status: Incomplete,
			Index:  len(line),
			Char:   0,
			Stack: stack,
		}
	}

	return Evaluation{
		Status: Ok,
		Index:  len(line) - 1,
		Char:   0,
		Stack: stack,
	}
}


func exec(path string) error {
	input, err := readInput(path)
	if err != nil {
		return err
	}

	score := 0
	scores := make([]int, 0)
	for _, line := range input {
		eval := EvalLine(line)
		if eval.Status == Corrupted {
			switch eval.Char {
			case ')':
				score += 3
			case ']':
				score += 57
			case '}':
				score += 1197
			case '>':
				score += 25137
			}
		} else if eval.Status == Incomplete {
			score2 := 0
			for eval.Stack.HasPop() {
				score2 *= 5
				pop, _ := eval.Stack.Pop()
				switch pop {
				case ')':
					score2 += 1
				case ']':
					score2 += 2
				case '}':
					score2 += 3
				case '>':
					score2 += 4
				}
			}
			scores = append(scores, score2)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})



	fmt.Println("Score", score)
	fmt.Println("Score2", scores[len(scores)/2])

	return nil
}

func readInput(path string) ([]string, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}