package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Counts map[string]int64
type Rules map[string]rune

func toS(a, b rune) string {
	return string([]rune{a, b})
}

func splitS(twoLetters string) (rune, rune){
	return rune(twoLetters[0]), rune(twoLetters[1])
}

func NextCounts(pairCounts Counts, rules Rules) Counts {
	out := make(Counts)
	for pair, count := range pairCounts {
		middle, ok := rules[pair]
		if !ok {
			panic(pair)
		}
		left, right := splitS(pair)

		out[toS(left, middle)] += count
		out[toS(middle, right)] += count
	}
	return out
}

func CountsFromTemplate(template string) Counts {
	chars := []rune(template)
	res := make(Counts)
	for i := 1; i < len(chars); i++ {
		pair := string([]rune{chars[i-1], chars[i]})
		res[pair] += 1
	}
	return res
}

func GetCharCounts(initTemplate string, counts Counts) map[rune]int64 {
	out := make(map[rune]int64)
	for pair, count := range counts {
		chars := []rune(pair)
		if len(chars) != 2 {
			panic(pair)
		}
		b := chars[1]
		out[b] += count
	}

	initRunes := []rune(initTemplate)
	out[initRunes[0]] += 1
	return out
}

func ReadLines(_path string) ([]string, error){
	f, err := os.Open(_path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

func main() {
	lines, err := ReadLines("inputs/input.txt")
	if err != nil {
		panic(err)
	}

	template := lines[0]
	rulesText := lines[2:]

	rules := make(map[string]rune)
	for _, line := range rulesText {
		var pre string
		var post rune
		_, err := fmt.Sscanf(line, "%s -> %c", &pre, &post)
		if err != nil {
			panic(line)
		}
		rules[pre] = post
	}

	fmt.Printf("Template: %s\n", template)
	pairs := CountsFromTemplate(template)
	for i := 1; i <= 40; i++ {
		pairs = NextCounts(pairs, rules)
	}

	counts := GetCharCounts(template, pairs)

	chars := make([]rune, 0, len(counts))
	for k := range counts {
		chars = append(chars, k)
	}

	sort.Slice(chars, func(i, j int) bool {
		return counts[chars[j]] > counts[chars[i]]
	})

	most, least := counts[chars[len(chars)-1]], counts[chars[0]]
	fmt.Printf("Result: %d - %d = %d\n", most, least, most-least)
}
