package coord

import (
	"day13/internal/fold"
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func calFold(pos, fold int) int {
	return fold - (pos - fold)
}

func (c Coord) Fold(f fold.Fold) Coord {
	if f.IsHorizontal {
		if c.Y < f.Value {
			return c
		}

		return Coord{
			X: c.X,
			Y: calFold(c.Y, f.Value),
		}
	} else {
		if c.X < f.Value {
			return c
		}

		return Coord{
			X: calFold(c.X, f.Value),
			Y: c.Y,
		}
	}
}

func FromText(in string) (Coord, error) {
	parts := strings.Split(in, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Coord{}, fmt.Errorf("coord.FromText: %v", err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Coord{}, fmt.Errorf("coord.FromText: %v", err)
	}

	return Coord{x, y}, nil
}
