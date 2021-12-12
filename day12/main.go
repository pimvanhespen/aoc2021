package main

import (
	"bufio"
	"day12/internal/cave"
	"day12/internal/chart"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := []string{
		"small",
		"demo",
		"large",
		"input",
	}

	for _, name := range inputs {
		fmt.Println()
		pth := fmt.Sprintf("input/%s.txt", name)
		fmt.Println(pth)
		if err := exec(pth); err != nil {
			panic(err)
		}
	}
}



type Strat1 struct{}

func (s Strat1) MustIgnore(node *cave.Cave, parent *chart.Tree) bool {
	if ! node.IsSmall {
		return false
	}

	return parent.ContainsNode(node) > 0
}

func (s Strat1) MustExpand(child *chart.Tree) bool {
	return child.Node.Name != "end"
}

func solve1(start *cave.Cave) int {
	root := &chart.Tree{
		Node:     start,
		Parent:   nil,
		Children: make([]*chart.Tree, 0, len(start.Links)),
	}

	root.Expand(&Strat1{})

	paths := 0
	for _, end := range root.GetEndNodes() {
		if end.Node.Name != "end" {
			continue
		}
		paths++
	}

	return paths
}


type Strat2 struct{}

func (s Strat2) MustIgnore(node *cave.Cave, parent *chart.Tree) bool {
	if ! node.IsSmall {
		return false
	}

	if node.Name == "start" {
		return true
	}

	return parent.ContainsNode(node) > 0 && parent.HasDouble()
}

func (s Strat2) MustExpand(child *chart.Tree) bool {
	return child.Node.Name != "end"
}

func solve2(start *cave.Cave) int {
	root := &chart.Tree{
		Node:     start,
		Parent:   nil,
		Children: make([]*chart.Tree, 0, len(start.Links)),
	}

	root.Expand(&Strat2{})

	paths := 0
	for _, end := range root.GetEndNodes() {
		if end.Node.Name != "end" {
			continue
		}
		paths++
	}

	return paths
}


func exec(pth string) error {
	start, err := readInput(pth)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	fmt.Println("Solve1: ", solve1(start))
	fmt.Println("Solve2: ", solve2(start))

	return nil
}

func find(name string, nodes []*cave.Cave) *cave.Cave {
	for _, node := range nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func readInput(pth string) (*cave.Cave, error) {
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	nodes := make([]*cave.Cave, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")

		nameA, nameB := parts[0], parts[1]
		a, b := find(nameA, nodes), find(nameB, nodes)

		if a == nil {
			a = cave.New(nameA)
			nodes = append(nodes, a)
		}
		if b == nil {
			b = cave.New(nameB)
			nodes = append(nodes, b)
		}

		a.Add(b)
		b.Add(a)
	}

	start := find("start", nodes)
	if start == nil {
		return nil, errors.New("start not found")
	}

	return start, nil
}