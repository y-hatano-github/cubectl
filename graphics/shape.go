package graphics

// Coordinate
type Point struct {
	X int
	Y int
}

// Convert a list of vertices of a face into a set of line segments
func polygon(x, y []int) []Point {
	ps := []Point{}

	// Connect each vertex to the next
	for i := 0; i < len(x)-1; i++ {
		ps = append(ps, line(x[i], y[i], x[i+1], y[i+1])...)
	}
	// Connect the last vertex to the first
	ps = append(ps, line(x[len(x)-1], y[len(y)-1], x[0], y[0])...)

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
