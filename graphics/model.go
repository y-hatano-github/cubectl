package graphics

import (
	"math"
	"sort"
)

// 画面サイズやモデル情報
type Model struct {
	Width  int
	Height int

	Vertices []*Vertex
	Faces    []Face

	center Point
	scale  float64
}

// 頂点
type Vertex struct {
	X int
	Y int
	Z int

	// 回転後の座標
	RX int
	RY int
	RZ int

	// スクリーン上の座標
	SX int
	SY int
}

// 面
type Face struct {
	V []*Vertex
	// 奥行き（面の平均Z）
	Z float64
}

func NewModel(width, height int) Model {
	return Model{
		Width:  width,
		Height: height,
	}
}

type VertexData [][3]int
type FaceData [][]int

// モデル設定
func (m *Model) Set(vd VertexData, fd FaceData) {
	for _, v := range vd {
		m.Vertices = append(m.Vertices, &Vertex{
			X: v[0],
			Y: v[1],
			Z: v[2],
		})
	}

	for _, f := range fd {
		var vs []*Vertex
		for _, idx := range f {
			vs = append(vs, m.Vertices[idx])
		}
		m.Faces = append(m.Faces, Face{V: vs})
	}

	m.scale = float64(m.Width) * 0.4 / 2
	m.center = Point{m.Width / 2, m.Height / 2}
}

// 描画用の座標列を取得
// t: Y軸回転（ヨー）
// p: X軸回転（ピッチ）
// z: 拡大率
// left, top: 描画オフセット
func (m *Model) GetShape(t, p, z float64, left, top int) [][]Point {
	m.update(t, p, z)

	var ps [][]Point
	for _, f := range m.Faces {
		var px []int
		var py []int
		for _, v := range f.V {
			px = append(px, v.SX+left)
			py = append(py, v.SY+top)
		}
		ps = append(ps, polygon(px, py))
	}
	return ps
}

// 頂点のスクリーン座標・面の奥行きを更新
func (m *Model) update(t, p, z float64) {
	ct := math.Cos(t) // yaw (Y軸)
	st := math.Sin(t)
	cp := math.Cos(p) // pitch (X軸)
	sp := math.Sin(p)

	s := m.scale * z

	// 頂点の回転＆投影
	for _, v := range m.Vertices {
		x := float64(v.X)
		y := float64(v.Y)
		zv := float64(v.Z)

		// --- Y軸回転（ヨー） ---
		// x' = x*cosT + z*sinT
		// y' = y
		// z' = -x*sinT + z*cosT
		x1 := x*ct + zv*st
		y1 := y
		z1 := -x*st + zv*ct

		// --- X軸回転（ピッチ） ---
		// x'' = x'
		// y'' = y'*cosP - z'*sinP
		// z'' = y'*sinP + z'*cosP
		x2 := x1
		y2 := y1*cp - z1*sp
		z2 := y1*sp + z1*cp

		v.RX = int(x2)
		v.RY = int(y2)
		v.RZ = int(z2)

		// スクリーン座標（平行投影）
		// 文字の横長補正をここで行う（x方向だけ2倍）
		v.SX = m.center.X + int(s*x2*2.0)
		v.SY = m.center.Y - int(s*y2)
	}

	// 面の奥行き更新（平均Z）
	for i := range m.Faces {
		f := &m.Faces[i]
		sumZ := 0.0
		for _, v := range f.V {
			sumZ += float64(v.RZ)
		}
		f.Z = sumZ / float64(len(f.V))
	}

	// 奥行きでソート：奥（Z小）→手前（Z大）に描画したいので
	// 「手前が後から描画される」ように昇順ではなく降順にする
	sort.SliceStable(m.Faces, func(i, j int) bool {
		return m.Faces[i].Z < m.Faces[j].Z
	})
}
