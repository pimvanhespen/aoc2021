package main

import "testing"

func TestFish_GetBirths(t *testing.T) {
	cases := []struct{
		Start, Limit int
		Expect []int
	}{
		{0, 10, []int{9}},
		{0, 24, []int{9, 16, 23}},
	}

	for _, tc := range cases {
		fish := NewFish(tc.Start)
		births := fish.GetBirths(tc.Limit)

		if len(births) != len(tc.Expect) {
			t.Errorf("Got %v, expect: %v", births, tc.Expect)
		} else {
			for i := range tc.Expect {
				if tc.Expect[i] != births[i] {
					t.Errorf("diff at index %d; got %d; expect %d", i, births[i], tc.Expect[i])
				}
			}
		}
	}
}

func TestFish_GetValue(t *testing.T) {
	cases := []struct {
		Start, Days, Expect int
	}{
		{0, 0, 8},
		{0, 1, 7},
		{0, 2, 6},
		{0, 3, 5},
		{0, 4, 4},
		{0, 5, 3},
		{0, 6, 2},
		{0, 7, 1},
		{0, 8, 0},
		{0, 9, 6},
		{0, 10, 5},
		{0, 11, 4},
		{0, 12, 3},
		{0, 13, 2},
		{0, 14, 1},
		{0, 15, 0},
	}

	for _, tc := range cases {
		f := NewFish(tc.Start)
		got := f.GetValue(tc.Days)

		if got != tc.Expect {
			t.Errorf("Got %d, Expected %d", got, tc.Expect)
		} else {
			t.Logf("Got %d", got)
		}
	}
}
