// Package ars provides math for the art.
package ars

import "math"

const (
	Tau = math.Pi * 2
)

type IntFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float64 | ~float32
}

type Vec2g[T IntFloat] struct {
	X, Y T
}

type Vec3g[T IntFloat] struct {
	X, Y, Z T
}

type Vec4g[T IntFloat] struct {
	X, Y, Z, W T
}

type Vec2 = Vec2g[float64]
type Vec2i = Vec2g[int]
type Vec3 = Vec3g[float64]
type Vec3i = Vec3g[int]
type Vec4 = Vec4g[float64]
type Vec4i = Vec4g[int]

func (v Vec2g[T]) Vec2() Vec2g[T] { return v }
func (v Vec2g[T]) Vec3() Vec3g[T] { return Vec3g[T]{v.X, v.Y, 0} }
func (v Vec2g[T]) Vec4() Vec4g[T] { return Vec4g[T]{v.X, v.Y, 0, 0} }

func (v Vec3g[T]) Vec2() Vec2g[T] { return Vec2g[T]{v.X, v.Y} }
func (v Vec3g[T]) Vec3() Vec3g[T] { return v }
func (v Vec3g[T]) Vec4() Vec4g[T] { return Vec4g[T]{v.X, v.Y, v.Z, 0} }

func (v Vec4g[T]) Vec2() Vec2g[T] { return Vec2g[T]{v.X, v.Y} }
func (v Vec4g[T]) Vec3() Vec3g[T] { return Vec3g[T]{v.X, v.Y, v.Z} }
func (v Vec4g[T]) Vec4() Vec4g[T] { return v }

func (v Vec2g[T]) ExtendZ(z T) Vec3g[T]     { return Vec3g[T]{v.X, v.Y, z} }
func (v Vec2g[T]) ExtendZW(z, w T) Vec4g[T] { return Vec4g[T]{v.X, v.Y, z, w} }
func (v Vec3g[T]) ExtendW(w T) Vec4g[T]     { return Vec4g[T]{v.X, v.Y, v.Z, w} }

func (v Vec2g[T]) Add(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{v.X + w.X, v.Y + w.Y} }

func (v Vec2g[T]) From(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{v.X - w.X, v.Y - w.Y} }

func (v Vec2g[T]) To(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{w.X - v.X, w.Y - v.Y} }

func (v Vec2g[T]) Scale(k T) Vec2g[T] { return Vec2g[T]{v.X * k, v.Y * k} }

func (v Vec2g[T]) Dot(w Vec2g[T]) T { return v.X*w.X + v.Y*w.Y }

// EwMul calculates entrywise multiplication.
func (v Vec2g[T]) EwMul(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{v.X * w.X, v.Y * w.Y} }

func (v Vec2g[T]) ScaleTo(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{w.X / v.X, w.Y / v.Y} }

func (v Vec2g[T]) ScaleFrom(w Vec2g[T]) Vec2g[T] { return Vec2g[T]{v.X / w.X, v.Y / w.Y} }
