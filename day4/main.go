package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Rows    = 5
	Columns = 5
)

// Box Single Bingo Game Box
type Box struct {
	Value  int
	Marked bool
}

func NewBox(num int) Box {
	return Box{
		Value:  num,
		Marked: false,
	}
}

// Board is the repesentation of the Bingo Game
// It is stateful
type Board [Rows * Columns]*Box

func (b Board) String() string {
	s := ""
	for n, box := range b {
		value := "."
		if box.Marked {
			value = "X"
		}

		s += fmt.Sprintf(" %2d %s ", box.Value, value)

		if n % Columns == Columns - 1 {
			s += "\n"
		}
	}
	return s
}

// Bingo has all data of the moment a Bingo occurred on a Board
type Bingo struct {
	Hits []int
	Miss []int
	Turns int
	Board Board
}

func NewBingo(board Board) *Bingo {
	return &Bingo{
		Hits:  nil,
		Miss:  nil,
		Turns: 0,
		Board: board,
	}
}

func (b *Bingo) HasRowBingo(row int) bool {
	begin := Columns * row
	end := Columns * (row + 1)
	for i := begin; i < end; i++ {
		if ! b.Board[i].Marked {
			return false
		}
	}
	return true
}

func (b *Bingo) HasColumnBingo(col int) bool {
	begin := col
	end := col + 20
	for i := begin; i <= end; i += Rows {
		if ! b.Board[i].Marked {
			return false
		}
	}
	return true
}

func (b *Bingo) HasBingo() bool {
	for i := 0; i < 5; i++ {
		if b.HasRowBingo(i) || b.HasColumnBingo(i) {
			return true
		}
	}
	return false
}

func (b *Bingo) CheckNumber(num int) bool {
	for _, box := range b.Board {
		if box.Value == num {
			box.Marked = true
			return true
		}
	}
	return false
}

func (b *Bingo) Score() int {
	sum := 0
	for _, n := range b.Miss {
		sum += n
	}

	lastHit := b.Hits[len(b.Hits) - 1]
	score := sum * lastHit
	return score
}

// Solve1 solves this bingo Board in the least turns possible
// The function returns after putting the board in a winning state
// this function returns an error when nu possibility has been found
func (b *Bingo) Solve1(numbers []int) error {
	var round, number int
	for round, number = range numbers {
		if b.CheckNumber(number) {
			b.Hits = append(b.Hits, number)
		}

		if ( len(b.Hits) >= Rows || len(b.Hits) >= Columns) && b.HasBingo(){
			break
		}
	}

	if !b.HasBingo() {
		return errors.New("no solution found")
	}

	for _, box := range b.Board {
		if ! box.Marked {
			b.Miss = append(b.Miss, box.Value)
		}
	}

	b.Turns = round

	return nil
}

func getBingoNumbers(scanner *bufio.Scanner) ([]int, error){
	if !scanner.Scan(){
		return nil, errors.New("no text found")
	}
	line := scanner.Text()
	strs := strings.Split(line, ",")
	numbers := make([]int, 0, len(strs))
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func LineToNumbers(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))
	for _, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func getBingoBoard(lines []string) (Board, error) {
	board := Board{}

	for row, line := range lines {
		numbers, err := LineToNumbers(line)
		if err != nil {
			return Board{}, err
		}
		if len(numbers) != Columns {
			return Board{}, fmt.Errorf("missing numbers, expeced len=%d, found %+v", Columns, numbers)
		}

		for col, num := range numbers {
			box := Box{
				Value:  num,
				Marked: false,
			}

			index := col + row * Columns
			board[index] = &box
		}
	}
	return board, nil
}

func getBingoBoards(scanner *bufio.Scanner) ([]Board, error){
	boards := make([]Board, 0)

	for scanner.Scan(){
		lines := make([]string, 0, Rows)
		for i := 0; i < Rows; i++ {
			if ! scanner.Scan() {
				return nil, errors.New("unexpected EOF")
			}
			line := scanner.Text()
			if line == "" {
				return nil, errors.New("unexpected empty line")
			}
			lines = append(lines, line)
		}
		board, err := getBingoBoard(lines)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	return boards, nil
}

func main(){
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers, err := getBingoNumbers(scanner)
	if err != nil {
		fmt.Println(err.Error())
		panic("something")
	}
	boards, err := getBingoBoards(scanner)
	if err != nil {
		panic(err.Error())
	}


	bingos := make([]*Bingo, 0, len(boards))

	for _, board := range boards {
		bingo := NewBingo(board)
		if err := bingo.Solve1(numbers); err != nil {
			continue
		}
		bingos = append(bingos, bingo)
	}

	sort.Slice(bingos, func(i, j int) bool {
		return bingos[i].Turns < bingos[j].Turns
	})

	best := bingos[0]
	worst := bingos[len(bingos) - 1]


	fmt.Printf("Best: %d\n", best.Score())
	fmt.Printf("Worst: %d\n", worst.Score())
}
