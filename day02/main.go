package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main(){
	if err := exec(); err != nil {
		panic(err)
	}
}

func exec() error {
	lines, err := getLines("input.txt")
	if err != nil {
		return err
	}

	solution2, err := solve2(lines)
	if err != nil {
		return err
	}

	fmt.Printf("The result is %d\n", solution2)

	return nil
}

func getLines(path string)([]string, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	return lines, nil
}

func solve2 (input []string) (int, error) {
	depth, aim, hori := 0, 0, 0

	for _, line := range input {
		dir, steps, err := parseMove(line)
		if err != nil {
			return 0, err
		}

		switch dir {
		case "up":
			aim -= steps
		case "down":
			aim += steps
		case "forward":
			hori += steps
			depth += steps * aim
		}
	}

	result := depth * hori

	return result, nil
}


func parseMove(line string) (string, int, error) {
	split := strings.Split(strings.TrimSpace(line), " ")
	move, textAmount := split[0], split[1]

	amount, err := strconv.Atoi(textAmount)
	if err != nil {
		return "", 0, fmt.Errorf("parseMove: %v", err)
	}

	return move, amount, nil
}