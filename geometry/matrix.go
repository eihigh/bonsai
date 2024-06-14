package geometry

// TODO: Mat3以外

var (
	Ident3 = Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
)

const (
	r0c0 = 0
	r0c1 = 1
	r0c2 = 2
	r1c0 = 3
	r1c1 = 4
	r1c2 = 5
	r2c0 = 6
	r2c1 = 7
	r2c2 = 8
)

// Mat3 is a 3x3 matrix.
type Mat3 [3 * 3]float64

func (m Mat3) Mul(n Mat3) Mat3 {
	return Mat3{
		m[r0c0]*n[r0c0] + m[r0c1]*n[r1c0] + m[r0c2]*n[r2c0],
		m[r0c0]*n[r0c1] + m[r0c1]*n[r1c1] + m[r0c2]*n[r2c1],
		m[r0c0]*n[r0c2] + m[r0c1]*n[r1c2] + m[r0c2]*n[r2c2],

		m[r1c0]*n[r0c0] + m[r1c1]*n[r1c0] + m[r1c2]*n[r2c0],
		m[r1c0]*n[r0c1] + m[r1c1]*n[r1c1] + m[r1c2]*n[r2c1],
		m[r1c0]*n[r0c2] + m[r1c1]*n[r1c2] + m[r1c2]*n[r2c2],

		m[r2c0]*n[r0c0] + m[r2c1]*n[r1c0] + m[r2c2]*n[r2c0],
		m[r2c0]*n[r0c1] + m[r2c1]*n[r1c1] + m[r2c2]*n[r2c1],
		m[r2c0]*n[r0c2] + m[r2c1]*n[r1c2] + m[r2c2]*n[r2c2],
	}
}
