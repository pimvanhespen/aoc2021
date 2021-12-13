package coord

import (
	"day13/internal/fold"
	"testing"
)

func TestCoord_Fold(t *testing.T) {
	testcases := []struct {
		In     Coord
		Fold   fold.Fold
		Expect Coord
	}{
		{In: Coord{2, 2}, Fold: fold.Fold{IsHorizontal: true, Value: 1}, Expect: Coord{2, 0}},
		{In: Coord{4, 4}, Fold: fold.Fold{IsHorizontal: true, Value: 2}, Expect: Coord{4, 0}},
		{In: Coord{2, 2}, Fold: fold.Fold{IsHorizontal: false, Value: 1}, Expect: Coord{0, 2}},
		{In: Coord{4, 4}, Fold: fold.Fold{IsHorizontal: false, Value: 2}, Expect: Coord{0, 4}},
	}

	for _, tc := range testcases {
		got := tc.In.Fold(tc.Fold)
		if got.X != tc.Expect.X || got.Y != tc.Expect.Y {
			t.Errorf("got %v, expect %v", got, tc.Expect)
		}
	}
}
