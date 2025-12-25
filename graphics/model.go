package graphics

import (
	"math"
	"sort"
)

// Screen size and model information
type Model struct {
	Width  int
	Height int

	Vertices []*Vertex
	Faces    []Face

	center Point
	scale  float64
}

// Vertex
type Vertex struct {
	X int
	Y int
	Z int

	// Coordinates after rotation
	RX int
	RY int
	RZ int

	// Coordinates on the screen
	SX int
	SY int
}

// Face
type Face struct {
	V []*Vertex
	// Depth (average Z of the face)
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

// Set model data
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

// Get polygon coordinates for drawing
// t: rotation around Y-axis (yaw)
// p: rotation around X-axis (pitch)
// z: scale factor
// left, top: drawing offset
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

// Update screen coordinates of vertices and depth of faces
func (m *Model) update(t, p, z float64) {
	ct := math.Cos(t) // yaw (Y-axis)
	st := math.Sin(t)
	cp := math.Cos(p) // pitch (X-axis)
	sp := math.Sin(p)

	s := m.scale * z

	// Rotate and project vertices
	for _, v := range m.Vertices {
		x := float64(v.X)
		y := float64(v.Y)
		zv := float64(v.Z)

		// --- Rotation around Y-axis (yaw) ---
		// x' = x*cosT + z*sinT
		// y' = y
		// z' = -x*sinT + z*cosT
		x1 := x*ct + zv*st
		y1 := y
		z1 := -x*st + zv*ct

		// --- Rotation around X-axis (pitch) ---
		// x'' = x'
		// y'' = y'*cosP - z'*sinP
		// z'' = y'*sinP + z'*cosP
		x2 := x1
		y2 := y1*cp - z1*sp
		z2 := y1*sp + z1*cp

		v.RX = int(x2)
		v.RY = int(y2)
		v.RZ = int(z2)

		// Screen coordinates (orthographic projection)
		// Apply horizontal stretch correction for text rendering (double X scale)
		v.SX = m.center.X + int(s*x2*2.0)
		v.SY = m.center.Y - int(s*y2)
	}

	// Update face depth (average Z)
	for i := range m.Faces {
		f := &m.Faces[i]
		sumZ := 0.0
		for _, v := range f.V {
			sumZ += float64(v.RZ)
		}
		f.Z = sumZ / float64(len(f.V))
	}

	// Sort faces by depth: far (small Z) → near (large Z)
	// Draw near faces later, so sort in descending order
	sort.SliceStable(m.Faces, func(i, j int) bool {
		return m.Faces[i].Z < m.Faces[j].Z
	})
}
