package chart

import "day12/internal/cave"

type Strategy interface {
	MustIgnore(node *cave.Cave, parent *Tree) bool
	MustExpand(child *Tree) bool
}