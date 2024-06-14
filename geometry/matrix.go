package geometry

import "math"

// TODO: Mat3以外

var (
	Ident3 = Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
)

// Mat3 is a 3x3 matrix.
type Mat3 [3 * 3]float64

// 上記Matに対して、平行移動、拡大縮小、回転、スキューのメソッドを提供する：

// Translate returns a matrix that translates by (x, y).
func (m Mat3) Translate(x, y float64) Mat3 {
	n := m
	n[6] += x
	n[7] += y
	return n
}

// Scale returns a matrix that scales by (x, y).
func (m Mat3) Scale(x, y float64) Mat3 {
	n := m
	n[0] *= x
	n[1] *= x
	n[2] *= x
	n[3] *= y
	n[4] *= y
	n[5] *= y
	return n
}

// Rotate returns a matrix that rotates by θ.
func (m Mat3) Rotate(θ float64) Mat3 {
	n := m
	if θ == 0 {
		return n
	}
	s, c := math.Sincos(θ)
	n[0], n[1] = c*n[0]-s*n[3], s*n[0]+c*n[3]
	n[3], n[4] = c*n[3]-s*n[4], s*n[3]+c*n[4]
	return n
}
