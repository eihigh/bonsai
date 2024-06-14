package geometry

import "math"

const (
	Tau     = math.Pi * 2
	Rad2Deg = 360 / Tau
	Deg2Rad = Tau / 360
)

type Rect struct {
	Min, Max Vec2
}

func Xyxy(x0, y0, x1, y1 float64) Rect { return Rect{Xy(x0, y0), Xy(x1, y1)} }
func Xywh(x, y, w, h float64) Rect     { return Rect{Xy(x, y), Xy(x+w, y+h)} }
