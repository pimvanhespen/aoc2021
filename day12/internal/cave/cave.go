package cave

import "strings"

type Cave struct {
	Name    string
	IsSmall bool
	Links   []*Cave
}

func New(name string) *Cave {
	singleUse := name != strings.ToUpper(name)
	return &Cave{
		Name:    name,
		IsSmall: singleUse,
		Links:   make([]*Cave, 0, 1),
	}
}

func(n *Cave) Add(other *Cave){
	n.Links = append(n.Links, other)
}

func(n *Cave) Equals(other *Cave) bool {
	return n.Name == other.Name
}
