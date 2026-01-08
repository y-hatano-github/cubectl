package graphics

import "sort"

// Coordinate
type Point struct {
	X int
	Y int
}

type PolygonData struct {
	Outline []Point // 輪郭線
	Fill    []Point // 塗りつぶし
}

// Convert a list of vertices of a face into a set of line segments
func polygon(x, y []int) PolygonData {
	return PolygonData{
		Outline: outline(x, y),
		Fill:    fll(x, y),
	}
}

func outline(x, y []int) []Point {
	ps := []Point{}

	// Connect each vertex to the next
	for i := 0; i < len(x)-1; i++ {
		ps = append(ps, line(x[i], y[i], x[i+1], y[i+1])...)
	}
	// Connect the last vertex to the first
	ps = append(ps, line(x[len(x)-1], y[len(y)-1], x[0], y[0])...)

	return ps
}

func fll(x, y []int) []Point {
	ps := []Point{}

	ymin, ymax := y[0], y[0]
	for _, v := range y {
		if v < ymin {
			ymin = v
		}
		if v > ymax {
			ymax = v
		}
	}

	for sy := ymin; sy <= ymax; sy++ {
		var xs []int

		for i := 0; i < len(x); i++ {
			j := (i + 1) % len(x)

			y1, y2 := y[i], y[j]
			x1, x2 := x[i], x[j]

			if (y1 <= sy && y2 > sy) || (y2 <= sy && y1 > sy) {
				ix := x1 + (sy-y1)*(x2-x1)/(y2-y1)
				xs = append(xs, ix)
			}
		}

		sort.Ints(xs)

		for i := 0; i+1 < len(xs); i += 2 {
			for sx := xs[i]; sx <= xs[i+1]; sx++ {
				ps = append(ps, Point{X: sx, Y: sy})
			}
		}
	}

	return ps
}

// Bresenham's line algorithm
func line(x1, y1, x2, y2 int) []Point {
	var dx, dy, sx, sy int

	if x2 >= x1 {
		sx = 1
		dx = x2 - x1
	} else {
		sx = -1
		dx = x1 - x2
	}

	if y2 >= y1 {
		sy = 1
		dy = y2 - y1
	} else {
		sy = -1
		dy = y1 - y2
	}

	ps := []Point{}
	x := x1
	y := y1

	if dx >= dy {
		e := dx / 2
		for i := 0; i <= dx; i++ {
			ps = append(ps, Point{X: x, Y: y})
			x += sx
			e -= dy
			if e < 0 {
				y += sy
				e += dx
			}
		}
	} else {
		e := dy / 2
		for i := 0; i <= dy; i++ {
			ps = append(ps, Point{X: x, Y: y})
			y += sy
			e -= dx
			if e < 0 {
				x += sx
				e += dy
			}
		}
	}

	return ps
}
