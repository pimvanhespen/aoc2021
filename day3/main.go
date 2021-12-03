package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type FilterFn func(col int, lines [][]byte) [][]byte

func grab(value byte, col int, lines [][]byte) [][]byte {
	matches := make([][]byte, 0) // <- not optimzed
	for _, line := range lines {
		if line[col] == value {
			matches = append(matches, line)
		}
	}
	return matches
}

func grabLeast(col int, lines [][]byte) [][]byte {

	zero, one := countOccurrences(col, lines)

	if zero <= one {
		return grab('0', col, lines)
	}

	return grab('1', col, lines)
}

func grabMost(col int, lines [][]byte) [][]byte {
	zero, one := countOccurrences(col, lines)

	if one >= zero {
		return grab('1', col, lines)
	}

	return grab('0', col, lines)
}

func countOccurrences(col int, lines [][]byte) (int, int) {
	zeroes := 0

	for _, line := range lines {
		if line[col] == '0' {
			zeroes++
		}
	}

	ones := len(lines) - zeroes

	return zeroes, ones
}

func runesToUint64(runes []byte) uint64 {
	var converted uint64
	for _, r := range runes {
		converted = converted << 1
		if r == '0' {
			continue
		}
		converted |= 1
	}
	return converted
}

func findRating(lines [][]byte, rateFn FilterFn) uint64 {
	remainder := lines
	length := len(lines[0])

	for i := 0; i < length; i++ {
		remainder = rateFn(i, remainder)

		if len(remainder) == 1 {
			return runesToUint64(remainder[0])
		}
	}

	// we only reach this stage when the loop has no good result
	return findRating(lines[1:], rateFn)
}

func FindOxygenGeneratorRating(lines [][]byte) uint64 {
	return findRating(lines, grabMost)
}

func FindCO2ScrubberRating(lines [][]byte) uint64 {
	return findRating(lines, grabLeast)
}

func CalculatePowerConsumption(lines [][]byte) uint64 {
	length := len(lines[0])

	balances := make([]int, length)

	for _, line := range lines {
		for index, b := range line {
			if b == '1' {
				balances[index]++
			} else {
				balances[index]--
			}
		}
	}

	var gamma, epsilon uint64

	for n, val := range balances {
		var bit uint64 = 1 << (length-n-1)
		if val > 0 {
			gamma |= bit
		} else {
			epsilon |= bit
		}
	}

	powerConsumption := gamma * epsilon

	return powerConsumption
}

func exec(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	lines := bytes.Split(data, []byte("\n"))

	powerConsumption := CalculatePowerConsumption(lines)
	fmt.Printf("Power Consumption: %d\n", powerConsumption)

	ogr := FindOxygenGeneratorRating(lines)
	fmt.Printf("Oxygen Genenerator Rating: %d\n", ogr)

	co2 := FindCO2ScrubberRating(lines)
	fmt.Printf("CO2 Scrubber Rating: %d\n", co2)

	lifeSupportRating := ogr * co2
	fmt.Printf("Life Support Rating: %d\n", lifeSupportRating)

	return nil
}

func main(){
	if err := exec("input.txt"); err != nil {
		panic(err)
	}
	if err := exec("demo.txt"); err != nil {
		panic(err)
	}
}
