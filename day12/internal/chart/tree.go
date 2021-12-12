package chart

import (
	"day12/internal/cave"
	"fmt"
)

type Tree struct {
	Node *cave.Cave
	Parent *Tree
	Children []*Tree
}

func NewTree(start *cave.Cave) *Tree {
	return &Tree{
		Node:     start,
		Parent:   nil,
		Children: make([]*Tree, 0, len(start.Links)),
	}
}

func (l *Tree) Add(child *Tree){
	l.Children = append(l.Children, child)
}

func (l *Tree) GetEndNodes() []*Tree {
	if len(l.Children) == 0 {
		return []*Tree{l}
	}

	endNodes := make([]*Tree, 0)
	for _, child := range l.Children {
		endNodes = append(endNodes, child.GetEndNodes()...)
	}
	return endNodes
}

func (l *Tree) String() string {
	s := l.Node.Name
	parent := l.Parent
	for parent != nil {
		s = fmt.Sprintf("%s,%s", parent.Node.Name, s)
		parent = parent.Parent
	}
	return s
}

func (l *Tree) ContainsNode(n *cave.Cave) int {
	total := 0
	parent := l.Parent
	for parent != nil {
		if parent.Node.Equals(n) {
			total++
		}
		parent = parent.Parent
	}
	return total
}

func (l *Tree) HasDouble() bool {
	m := map[string]struct{}{}

	parent := l
	for parent != nil {
		if ! parent.Node.IsSmall {
			parent = parent.Parent
			continue
		}
		name := parent.Node.Name
		if _, exists := m[name]; exists {
			return true
		}

		m[name] = struct{}{}

		parent = parent.Parent
	}
	return false
}

func (l *Tree) Expand(s Strategy) {
	for _, node := range l.Node.Links {
		if s.MustIgnore(node, l) {
			continue
		}

		child := &Tree{
			Node:     node,
			Parent:   l,
			Children: make([]*Tree, 0),
		}
		l.Add(child)

		if s.MustExpand(child){
			child.Expand(s)
		}
	}
}