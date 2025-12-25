package graphics

// 座標
type Point struct {
	X int
	Y int
}

// 面を構成する頂点列から線分群に変換
func polygon(x, y []int) []Point {
	ps := []Point{}

	// 各頂点を結ぶ
	for i := 0; i < len(x)-1; i++ {
		ps = append(ps, line(x[i], y[i], x[i+1], y[i+1])...)
	}
	// 最後と最初を結ぶ
	ps = append(ps, line(x[len(x)-1], y[len(y)-1], x[0], y[0])...)

	return ps
}

// Bresenham の直線アルゴリズム
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

	// 最初の点はループ内で追加する
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
