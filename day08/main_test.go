package main

import (
	"fmt"
	"testing"
)

func TestLetters_Difference(t *testing.T) {
	testcases := []struct {
		A, B, Expect Letters
	}{
		{[]rune("abc"), []rune("ab"), []rune("c")},
		{[]rune("sadf"), []rune("fds"), []rune("a")},
	}

	for _, tc := range testcases {
		got := tc.A.Difference(tc.B)

		if !got.ContainsAll(tc.Expect) {
			t.Errorf("gpt: %+v; expect %+v", got, tc.Expect)
		}
	}
}

func TestDigit_String(t *testing.T) {
	digit := Digit{
		Top:         'a',
		TopLeft:     'b',
		TopRight:    'c',
		Center:      'd',
		BottomLeft:  'e',
		BottomRight: 'f',
		Bottom:      'g',
	}

	fmt.Println(digit)
}

func TestDigit_GetNum(t *testing.T) {
	digit := Digit{
		Top:         'a',
		TopLeft:     'b',
		TopRight:    'c',
		Center:      'd',
		BottomLeft:  'e',
		BottomRight: 'f',
		Bottom:      'g',
	}

	testcases := []struct {
		Input  Letters
		Expect int
	}{
		{
			Input:  []rune("abcefg"),
			Expect: 0,
		},
		{
			Input:  []rune("cf"),
			Expect: 1,
		},
		{
			Input:  []rune("acdeg"),
			Expect: 2,
		},
		{
			Input:  []rune("acdfg"),
			Expect: 3,
		},
		{
			Input:  []rune("bdcf"),
			Expect: 4,
		},
		{
			Input:  []rune("abdfg"),
			Expect: 5,
		},
		{
			Input:  []rune("abdefg"),
			Expect: 6,
		},
		{
			Input:  []rune("acf"),
			Expect: 7,
		},
		{
			Input:  []rune("abcdefg"),
			Expect: 8,
		},
		{
			Input:  []rune("abcdfg"),
			Expect: 9,
		},
	}

	for _, tc := range testcases {
		got := digit.ParseNumber(tc.Input)
		if got != tc.Expect {
			t.Errorf("got: %d, expect: %d", got, tc.Expect)
		}
	}
}
