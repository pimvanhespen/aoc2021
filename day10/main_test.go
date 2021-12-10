package main

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Push('a')
	s.Push('b')

	b, ok := s.Pop()
	if ! ok {
		t.Error("b Not OK")
	}
	if b != 'b' {
		t.Error("b is not b")
	}
	a, ok := s.Pop()
	if ! ok {
		t.Error("a Not OK")
	}
	if a != 'a' {
		t.Error("A is not a")
	}
}
