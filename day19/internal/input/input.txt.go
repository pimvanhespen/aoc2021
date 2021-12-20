package input

import (
	"io"
	"os"
	"strings"
)

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

func GroupLines(lines []string, sep string) [][]string {
	groups := make([][]string, 0)
	current := make([]string, 0)
	for _, line := range lines {
		if line == sep {
			groups = append(groups, current)
			current = make([]string, 0)
			continue
		}
		current = append(current, line)
	}
	groups = append(groups, current)
	return groups
}

func ReadLinesGrouped(in string) ([][]string, error){
	lines, err := ReadLines(in)
	if err != nil {
		return nil, err
	}

	return GroupLines(lines, ""), nil
}
