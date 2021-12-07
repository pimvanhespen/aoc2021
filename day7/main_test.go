package main

import "testing"

func TestMoveCost(t *testing.T) {
	testcases := []struct{
		input, expect int
	}{
		{0, 0},
		{1,1},
		{2,3},
		{3, 6},
		{4, 10},
		{5, 15},
		{6, 21},
	}

	for _, tc := range testcases {
		got := moveCost(tc.input)

		if got != tc.expect {
			t.Errorf("got %d; expect %d", got, tc.expect)
		}
	}
}
