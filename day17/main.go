package main

import (
	"fmt"
)

func main() {
	//test := NewSquare(20, 30, -5, -10)
	//if err := exec(test); err != nil {
	//	panic(err)
	//}

	input := NewSquare(287, 309, -48, -76)
	if err := exec(input); err != nil {
		panic(err)
	}
}

var testCompare = `23,-10;25,-9;27,-5;29,-6;22,-6;21,-7;9,0;27,-7;24,-5;25,-7;26,-6;25,-5;6,8;11,-2;20,-5;29,-10;6,3;28,-7;8,0;30,-6;29,-8;20,-10;6,7;6,4;6,1;14,-4;21,-6;26,-10;7,-1;7,7;8,-1;21,-9;6,2;20,-7;30,-10;14,-3;20,-8;13,-2;7,3;28,-8;29,-9;15,-3;22,-5;26,-8;25,-8;25,-6;15,-4;9,-2;15,-2;12,-2;28,-9;12,-3;24,-6;23,-7;25,-10;7,8;11,-3;26,-7;7,1;23,-9;6,0;22,-10;27,-6;8,1;22,-8;13,-4;7,6;28,-6;11,-4;12,-4;26,-9;7,4;24,-10;23,-8;30,-8;7,0;9,-1;10,-1;26,-5;22,-9;6,5;7,5;23,-6;28,-10;10,-2;11,-1;20,-9;14,-2;29,-7;13,-3;23,-5;24,-8;27,-9;30,-7;28,-5;21,-10;7,9;6,6;21,-5;27,-10;7,2;30,-9;21,-8;22,-7;24,-9;20,-6;6,9;29,-5;8,-2;27,-8;30,-5;24,-7`

func exec(s Square) error {

	//fmt.Println("solve1:", solve1(s))
	x2, _ := solve2(s)
	fmt.Println("solve2:", x2)

	//parts := strings.Split(testCompare, ";")
	//given := make([]Point, len(parts))
	//for i, p := range parts {
	//	var x, y int
	//	_, _ = fmt.Sscanf(p, "%d,%d", &x, &y)
	//	given[i] = Point{x, y}
	//}
	//
	//sort.Slice(pts, func(i, j int) bool {
	//	if pts[i].X == pts[j].X {
	//		return pts[i].Y < pts[j].Y
	//	}
	//	return pts[i].X < pts[j].X
	//})
	//
	//sort.Slice(given, func(i, j int) bool {
	//	if given[i].X == given[j].X {
	//		return given[i].Y < given[j].Y
	//	}
	//	return given[i].X < given[j].X
	//})
	//
	//var notFound []Point
	//
	//for _, p1 := range given {
	//	var found bool
	//	for _, p2 := range pts {
	//		if p1 == p2 {
	//			found = true
	//			break
	//		}
	//	}
	//	if !found {
	//		notFound = append(notFound, p1)
	//	}
	//}
	//
	//for _, nf := range notFound {
	//	ps := plot(nf.X, nf.Y, s.RightBottom.X, s.RightBottom.Y)
	//	for _, p := range ps {
	//		if s.Contains(p) {
	//			fmt.Println("not found:", nf, "=>", p, "   ", ps)
	//			break
	//		}
	//	}
	//}

	return nil
}

func plot(xv, yv int, maxX, miny int) []Point {
	points := []Point{
		{0, 0},
	}

	iter := 0
	for {
		iter++

		point := Point{
			X: points[iter-1].X + xv,
			Y: points[iter-1].Y + yv,
		}

		points = append(points, point)

		if xv > 0 {
			xv = max(0, xv-1)
		} else if xv < 0 {
			xv = min(0, xv+1)
		}

		yv--

		if point.X > maxX || point.Y < miny {
			break
		}
	}

	return points
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//}
//
//func solve1(s Square) int {
//	var highest int
//
//	for x := 0; x < s.X+s.Width; x++ {
//		for y := s.Y; y < 100; y++ {
//			var maxY int
//			for _, p := range plot(x, y, s.X+s.Width, s.Y) {
//				maxY = max(maxY, p.Y)
//				if s.Contains(p) {
//					if maxY > highest {
//						//log.Printf("new highest: (%d, %d) => %d\n", x, y, maxY)
//						highest = maxY
//					}
//					continue
//				}
//			}
//		}
//	}
//	return highest
//}

func solve2(s Square) (int, []Point) {

	initialVelocities := map[Point]struct{}{}

	for xVelocity := 0; xVelocity <= s.RightBottom.X; xVelocity++ {
		for yVelocity := s.RightBottom.Y; yVelocity < 100; yVelocity++ {

			pts := plot(xVelocity, yVelocity, s.RightBottom.X, s.RightBottom.Y)
			if !s.ContainsAny(pts) {
				continue
			}

			initialVelocities[Point{xVelocity, yVelocity}] = struct{}{}
		}
	}

	pts := make([]Point, 0, len(initialVelocities))
	for p := range initialVelocities {
		pts = append(pts, p)
	}

	return len(pts), pts
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%3d,%3d)", p.X, p.Y)
}

type Square struct {
	LeftTop, RightBottom Point
}

func NewSquare(x1, x2, y1, y2 int) Square {
	return Square{
		LeftTop:     Point{x1, y1},
		RightBottom: Point{x2, y2},
	}
}

func (t Square) Contains(point Point) bool {

	return point.X >= t.LeftTop.X &&
		point.X <= t.RightBottom.X &&
		point.Y <= t.LeftTop.Y &&
		point.Y >= t.RightBottom.Y
}

func (t Square) ContainsAny(points []Point) bool {
	for _, p := range points {
		if t.Contains(p) {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
