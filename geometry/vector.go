package geometry

import (
	"cmp"
	"math"
)

// 意思を持って float64 の代わりに Scalar を使う
// そうすれば幸せになると思う、多分

type Scalar float64

type Vec2 struct {
	X, Y Scalar
}

type Vec3 struct {
	X, Y, Z Scalar
}

type Vec4 struct {
	X, Y, Z, W Scalar
}

// ========================================
// Constructors
// ========================================

func Xy[T ~float64](x, y T) Vec2         { return Vec2{Scalar(x), Scalar(y)} }
func Xyz[T ~float64](x, y, z T) Vec3     { return Vec3{Scalar(x), Scalar(y), Scalar(z)} }
func Xyzw[T ~float64](x, y, z, w T) Vec4 { return Vec4{Scalar(x), Scalar(y), Scalar(z), Scalar(w)} }

func All2[T ~float64](v T) Vec2 { return Vec2{Scalar(v), Scalar(v)} }
func All3[T ~float64](v T) Vec3 { return Vec3{Scalar(v), Scalar(v), Scalar(v)} }
func All4[T ~float64](v T) Vec4 { return Vec4{Scalar(v), Scalar(v), Scalar(v), Scalar(v)} }

// ========================================
// Conversions
// ========================================

func (a Scalar) Float64() float64     { return float64(a) }
func (a Vec2) Xy() (x, y float64)     { return float64(a.X), float64(a.Y) }
func (a Vec3) Xyz() (x, y, z float64) { return float64(a.X), float64(a.Y), float64(a.Z) }
func (a Vec4) Xyzw() (x, y, z, w float64) {
	return float64(a.X), float64(a.Y), float64(a.Z), float64(a.W)
}

func (a Vec2) Scalars() (x, y Scalar)       { return a.X, a.Y }
func (a Vec3) Scalars() (x, y, z Scalar)    { return a.X, a.Y, a.Z }
func (a Vec4) Scalars() (x, y, z, w Scalar) { return a.X, a.Y, a.Z, a.W }

func (a Scalar) All2() Vec2 { return Vec2{a, a} }
func (a Scalar) All3() Vec3 { return Vec3{a, a, a} }
func (a Scalar) All4() Vec4 { return Vec4{a, a, a, a} }

func (a Scalar) Vec2(y Scalar) Vec2       { return Vec2{a, y} }
func (a Scalar) Vec3(y, z Scalar) Vec3    { return Vec3{a, y, z} }
func (a Scalar) Vec4(y, z, w Scalar) Vec4 { return Vec4{a, y, z, w} }

func (a Vec2) Vec3(z Scalar) Vec3    { return Vec3{a.X, a.Y, z} }
func (a Vec2) Vec4(z, w Scalar) Vec4 { return Vec4{a.X, a.Y, z, w} }

func (a Vec3) Vec2() Vec2         { return Vec2{a.X, a.Y} }
func (a Vec3) Vec4(w Scalar) Vec4 { return Vec4{a.X, a.Y, a.Z, w} }

func (a Vec4) Vec2() Vec2 { return Vec2{a.X, a.Y} }
func (a Vec4) Vec3() Vec3 { return Vec3{a.X, a.Y, a.Z} }

func (a Scalar) Array() [1]Scalar { return [1]Scalar{a} }
func (a Vec2) Array() [2]Scalar   { return [2]Scalar{a.X, a.Y} }
func (a Vec3) Array() [3]Scalar   { return [3]Scalar{a.X, a.Y, a.Z} }
func (a Vec4) Array() [4]Scalar   { return [4]Scalar{a.X, a.Y, a.Z, a.W} }

func (a Scalar) Slice() []Scalar { return []Scalar{a} }
func (a Vec2) Slice() []Scalar   { return []Scalar{a.X, a.Y} }
func (a Vec3) Slice() []Scalar   { return []Scalar{a.X, a.Y, a.Z} }
func (a Vec4) Slice() []Scalar   { return []Scalar{a.X, a.Y, a.Z, a.W} }

// ========================================
// Arithmetic
// ========================================

func (a Scalar) Neg() Scalar { return -a }
func (a Vec2) Neg() Vec2     { return Vec2{-a.X, -a.Y} }
func (a Vec3) Neg() Vec3     { return Vec3{-a.X, -a.Y, -a.Z} }
func (a Vec4) Neg() Vec4     { return Vec4{-a.X, -a.Y, -a.Z, -a.W} }

func (a Scalar) Add(b Scalar) Scalar { return a + b }
func (a Vec2) Add(b Vec2) Vec2       { return Vec2{a.X + b.X, a.Y + b.Y} }
func (a Vec3) Add(b Vec3) Vec3       { return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a Vec4) Add(b Vec4) Vec4       { return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W} }

func (a Vec2) AddAll(b Scalar) Vec2 { return Vec2{a.X + b, a.Y + b} }
func (a Vec3) AddAll(b Scalar) Vec3 { return Vec3{a.X + b, a.Y + b, a.Z + b} }
func (a Vec4) AddAll(b Scalar) Vec4 { return Vec4{a.X + b, a.Y + b, a.Z + b, a.W + b} }

func (a Scalar) Sub(b Scalar) Scalar { return a - b }
func (a Vec2) Sub(b Vec2) Vec2       { return Vec2{a.X - b.X, a.Y - b.Y} }
func (a Vec3) Sub(b Vec3) Vec3       { return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z} }
func (a Vec4) Sub(b Vec4) Vec4       { return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W} }

func (a Vec2) SubAll(b Scalar) Vec2 { return Vec2{a.X - b, a.Y - b} }
func (a Vec3) SubAll(b Scalar) Vec3 { return Vec3{a.X - b, a.Y - b, a.Z - b} }
func (a Vec4) SubAll(b Scalar) Vec4 { return Vec4{a.X - b, a.Y - b, a.Z - b, a.W - b} }

func (a Scalar) Mul(b Scalar) Scalar { return a * b }
func (a Vec2) Mul(b Vec2) Vec2       { return Vec2{a.X * b.X, a.Y * b.Y} }
func (a Vec3) Mul(b Vec3) Vec3       { return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z} }
func (a Vec4) Mul(b Vec4) Vec4       { return Vec4{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W} }

func (a Vec2) MulAll(b Scalar) Vec2 { return Vec2{a.X * b, a.Y * b} }
func (a Vec3) MulAll(b Scalar) Vec3 { return Vec3{a.X * b, a.Y * b, a.Z * b} }
func (a Vec4) MulAll(b Scalar) Vec4 { return Vec4{a.X * b, a.Y * b, a.Z * b, a.W * b} }

// Scale: alias for MulAll

func (a Vec2) Scale(b Scalar) Vec2 { return a.MulAll(b) }
func (a Vec3) Scale(b Scalar) Vec3 { return a.MulAll(b) }
func (a Vec4) Scale(b Scalar) Vec4 { return a.MulAll(b) }

func (a Scalar) Div(b Scalar) Scalar { return a / b }
func (a Vec2) Div(b Vec2) Vec2       { return Vec2{a.X / b.X, a.Y / b.Y} }
func (a Vec3) Div(b Vec3) Vec3       { return Vec3{a.X / b.X, a.Y / b.Y, a.Z / b.Z} }
func (a Vec4) Div(b Vec4) Vec4       { return Vec4{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W} }

func (a Vec2) DivAll(b Scalar) Vec2 { return Vec2{a.X / b, a.Y / b} }
func (a Vec3) DivAll(b Scalar) Vec3 { return Vec3{a.X / b, a.Y / b, a.Z / b} }
func (a Vec4) DivAll(b Scalar) Vec4 { return Vec4{a.X / b, a.Y / b, a.Z / b, a.W / b} }

// ========================================
// Functions from GLSL
// ========================================

func (a Scalar) Sin() Scalar { return Scalar(math.Sin(float64(a))) }
func (a Vec2) Sin() Vec2     { return Vec2{a.X.Sin(), a.Y.Sin()} }
func (a Vec3) Sin() Vec3     { return Vec3{a.X.Sin(), a.Y.Sin(), a.Z.Sin()} }
func (a Vec4) Sin() Vec4     { return Vec4{a.X.Sin(), a.Y.Sin(), a.Z.Sin(), a.W.Sin()} }

func (a Scalar) Cos() Scalar { return Scalar(math.Cos(float64(a))) }
func (a Vec2) Cos() Vec2     { return Vec2{a.X.Cos(), a.Y.Cos()} }
func (a Vec3) Cos() Vec3     { return Vec3{a.X.Cos(), a.Y.Cos(), a.Z.Cos()} }
func (a Vec4) Cos() Vec4     { return Vec4{a.X.Cos(), a.Y.Cos(), a.Z.Cos(), a.W.Cos()} }

func (a Scalar) Tan() Scalar { return Scalar(math.Tan(float64(a))) }
func (a Vec2) Tan() Vec2     { return Vec2{a.X.Tan(), a.Y.Tan()} }
func (a Vec3) Tan() Vec3     { return Vec3{a.X.Tan(), a.Y.Tan(), a.Z.Tan()} }
func (a Vec4) Tan() Vec4     { return Vec4{a.X.Tan(), a.Y.Tan(), a.Z.Tan(), a.W.Tan()} }

func (a Scalar) Asin() Scalar { return Scalar(math.Asin(float64(a))) }
func (a Vec2) Asin() Vec2     { return Vec2{a.X.Asin(), a.Y.Asin()} }
func (a Vec3) Asin() Vec3     { return Vec3{a.X.Asin(), a.Y.Asin(), a.Z.Asin()} }
func (a Vec4) Asin() Vec4     { return Vec4{a.X.Asin(), a.Y.Asin(), a.Z.Asin(), a.W.Asin()} }

func (a Scalar) Acos() Scalar { return Scalar(math.Acos(float64(a))) }
func (a Vec2) Acos() Vec2     { return Vec2{a.X.Acos(), a.Y.Acos()} }
func (a Vec3) Acos() Vec3     { return Vec3{a.X.Acos(), a.Y.Acos(), a.Z.Acos()} }
func (a Vec4) Acos() Vec4     { return Vec4{a.X.Acos(), a.Y.Acos(), a.Z.Acos(), a.W.Acos()} }

func (a Scalar) Atan() Scalar { return Scalar(math.Atan(float64(a))) }
func (a Vec2) Atan() Vec2     { return Vec2{a.X.Atan(), a.Y.Atan()} }
func (a Vec3) Atan() Vec3     { return Vec3{a.X.Atan(), a.Y.Atan(), a.Z.Atan()} }
func (a Vec4) Atan() Vec4     { return Vec4{a.X.Atan(), a.Y.Atan(), a.Z.Atan(), a.W.Atan()} }

func (a Scalar) Atan2(b Scalar) Scalar { return Scalar(math.Atan2(float64(a), float64(b))) }
func (a Vec2) Atan2(b Vec2) Vec2       { return Vec2{a.X.Atan2(b.X), a.Y.Atan2(b.Y)} }
func (a Vec3) Atan2(b Vec3) Vec3       { return Vec3{a.X.Atan2(b.X), a.Y.Atan2(b.Y), a.Z.Atan2(b.Z)} }
func (a Vec4) Atan2(b Vec4) Vec4 {
	return Vec4{a.X.Atan2(b.X), a.Y.Atan2(b.Y), a.Z.Atan2(b.Z), a.W.Atan2(b.W)}
}

func (a Vec2) Atan2All(b Scalar) Vec2 { return Vec2{a.X.Atan2(b), a.Y.Atan2(b)} }
func (a Vec3) Atan2All(b Scalar) Vec3 { return Vec3{a.X.Atan2(b), a.Y.Atan2(b), a.Z.Atan2(b)} }
func (a Vec4) Atan2All(b Scalar) Vec4 {
	return Vec4{a.X.Atan2(b), a.Y.Atan2(b), a.Z.Atan2(b), a.W.Atan2(b)}
}

func (a Scalar) Sinh() Scalar { return Scalar(math.Sinh(float64(a))) }
func (a Vec2) Sinh() Vec2     { return Vec2{a.X.Sinh(), a.Y.Sinh()} }
func (a Vec3) Sinh() Vec3     { return Vec3{a.X.Sinh(), a.Y.Sinh(), a.Z.Sinh()} }
func (a Vec4) Sinh() Vec4     { return Vec4{a.X.Sinh(), a.Y.Sinh(), a.Z.Sinh(), a.W.Sinh()} }

func (a Scalar) Cosh() Scalar { return Scalar(math.Cosh(float64(a))) }
func (a Vec2) Cosh() Vec2     { return Vec2{a.X.Cosh(), a.Y.Cosh()} }
func (a Vec3) Cosh() Vec3     { return Vec3{a.X.Cosh(), a.Y.Cosh(), a.Z.Cosh()} }
func (a Vec4) Cosh() Vec4     { return Vec4{a.X.Cosh(), a.Y.Cosh(), a.Z.Cosh(), a.W.Cosh()} }

func (a Scalar) Tanh() Scalar { return Scalar(math.Tanh(float64(a))) }
func (a Vec2) Tanh() Vec2     { return Vec2{a.X.Tanh(), a.Y.Tanh()} }
func (a Vec3) Tanh() Vec3     { return Vec3{a.X.Tanh(), a.Y.Tanh(), a.Z.Tanh()} }
func (a Vec4) Tanh() Vec4     { return Vec4{a.X.Tanh(), a.Y.Tanh(), a.Z.Tanh(), a.W.Tanh()} }

func (a Scalar) Asinh() Scalar { return Scalar(math.Asinh(float64(a))) }
func (a Vec2) Asinh() Vec2     { return Vec2{a.X.Asinh(), a.Y.Asinh()} }
func (a Vec3) Asinh() Vec3     { return Vec3{a.X.Asinh(), a.Y.Asinh(), a.Z.Asinh()} }
func (a Vec4) Asinh() Vec4     { return Vec4{a.X.Asinh(), a.Y.Asinh(), a.Z.Asinh(), a.W.Asinh()} }

func (a Scalar) Acosh() Scalar { return Scalar(math.Acosh(float64(a))) }
func (a Vec2) Acosh() Vec2     { return Vec2{a.X.Acosh(), a.Y.Acosh()} }
func (a Vec3) Acosh() Vec3     { return Vec3{a.X.Acosh(), a.Y.Acosh(), a.Z.Acosh()} }
func (a Vec4) Acosh() Vec4     { return Vec4{a.X.Acosh(), a.Y.Acosh(), a.Z.Acosh(), a.W.Acosh()} }

func (a Scalar) Atanh() Scalar { return Scalar(math.Atanh(float64(a))) }
func (a Vec2) Atanh() Vec2     { return Vec2{a.X.Atanh(), a.Y.Atanh()} }
func (a Vec3) Atanh() Vec3     { return Vec3{a.X.Atanh(), a.Y.Atanh(), a.Z.Atanh()} }
func (a Vec4) Atanh() Vec4     { return Vec4{a.X.Atanh(), a.Y.Atanh(), a.Z.Atanh(), a.W.Atanh()} }

func (a Scalar) Pow(b Scalar) Scalar { return Scalar(math.Pow(float64(a), float64(b))) }
func (a Vec2) Pow(b Vec2) Vec2       { return Vec2{a.X.Pow(b.X), a.Y.Pow(b.Y)} }
func (a Vec3) Pow(b Vec3) Vec3       { return Vec3{a.X.Pow(b.X), a.Y.Pow(b.Y), a.Z.Pow(b.Z)} }
func (a Vec4) Pow(b Vec4) Vec4       { return Vec4{a.X.Pow(b.X), a.Y.Pow(b.Y), a.Z.Pow(b.Z), a.W.Pow(b.W)} }

func (a Vec2) PowAll(b Scalar) Vec2 { return Vec2{a.X.Pow(b), a.Y.Pow(b)} }
func (a Vec3) PowAll(b Scalar) Vec3 { return Vec3{a.X.Pow(b), a.Y.Pow(b), a.Z.Pow(b)} }
func (a Vec4) PowAll(b Scalar) Vec4 { return Vec4{a.X.Pow(b), a.Y.Pow(b), a.Z.Pow(b), a.W.Pow(b)} }

func (a Scalar) Exp() Scalar { return Scalar(math.Exp(float64(a))) }
func (a Vec2) Exp() Vec2     { return Vec2{a.X.Exp(), a.Y.Exp()} }
func (a Vec3) Exp() Vec3     { return Vec3{a.X.Exp(), a.Y.Exp(), a.Z.Exp()} }
func (a Vec4) Exp() Vec4     { return Vec4{a.X.Exp(), a.Y.Exp(), a.Z.Exp(), a.W.Exp()} }

func (a Scalar) Log() Scalar { return Scalar(math.Log(float64(a))) }
func (a Vec2) Log() Vec2     { return Vec2{a.X.Log(), a.Y.Log()} }
func (a Vec3) Log() Vec3     { return Vec3{a.X.Log(), a.Y.Log(), a.Z.Log()} }
func (a Vec4) Log() Vec4     { return Vec4{a.X.Log(), a.Y.Log(), a.Z.Log(), a.W.Log()} }

func (a Scalar) Exp2() Scalar { return Scalar(math.Exp2(float64(a))) }
func (a Vec2) Exp2() Vec2     { return Vec2{a.X.Exp2(), a.Y.Exp2()} }
func (a Vec3) Exp2() Vec3     { return Vec3{a.X.Exp2(), a.Y.Exp2(), a.Z.Exp2()} }
func (a Vec4) Exp2() Vec4     { return Vec4{a.X.Exp2(), a.Y.Exp2(), a.Z.Exp2(), a.W.Exp2()} }

func (a Scalar) Log2() Scalar { return Scalar(math.Log2(float64(a))) }
func (a Vec2) Log2() Vec2     { return Vec2{a.X.Log2(), a.Y.Log2()} }
func (a Vec3) Log2() Vec3     { return Vec3{a.X.Log2(), a.Y.Log2(), a.Z.Log2()} }
func (a Vec4) Log2() Vec4     { return Vec4{a.X.Log2(), a.Y.Log2(), a.Z.Log2(), a.W.Log2()} }

func (a Scalar) Sqrt() Scalar { return Scalar(math.Sqrt(float64(a))) }
func (a Vec2) Sqrt() Vec2     { return Vec2{a.X.Sqrt(), a.Y.Sqrt()} }
func (a Vec3) Sqrt() Vec3     { return Vec3{a.X.Sqrt(), a.Y.Sqrt(), a.Z.Sqrt()} }
func (a Vec4) Sqrt() Vec4     { return Vec4{a.X.Sqrt(), a.Y.Sqrt(), a.Z.Sqrt(), a.W.Sqrt()} }

func (a Scalar) InverseSqrt() Scalar { return Scalar(1) / a.Sqrt() }
func (a Vec2) InverseSqrt() Vec2     { return Vec2{a.X.InverseSqrt(), a.Y.InverseSqrt()} }
func (a Vec3) InverseSqrt() Vec3 {
	return Vec3{a.X.InverseSqrt(), a.Y.InverseSqrt(), a.Z.InverseSqrt()}
}
func (a Vec4) InverseSqrt() Vec4 {
	return Vec4{a.X.InverseSqrt(), a.Y.InverseSqrt(), a.Z.InverseSqrt(), a.W.InverseSqrt()}
}

func (a Scalar) Abs() Scalar { return Scalar(math.Abs(float64(a))) }
func (a Vec2) Abs() Vec2     { return Vec2{a.X.Abs(), a.Y.Abs()} }
func (a Vec3) Abs() Vec3     { return Vec3{a.X.Abs(), a.Y.Abs(), a.Z.Abs()} }
func (a Vec4) Abs() Vec4     { return Vec4{a.X.Abs(), a.Y.Abs(), a.Z.Abs(), a.W.Abs()} }

func sign[T ~float64](x T) T {
	return T(cmp.Compare(x, 0))
}

func (a Scalar) Sign() Scalar { return Scalar(sign(a)) }
func (a Vec2) Sign() Vec2     { return Vec2{a.X.Sign(), a.Y.Sign()} }
func (a Vec3) Sign() Vec3     { return Vec3{a.X.Sign(), a.Y.Sign(), a.Z.Sign()} }
func (a Vec4) Sign() Vec4     { return Vec4{a.X.Sign(), a.Y.Sign(), a.Z.Sign(), a.W.Sign()} }

func (a Scalar) Floor() Scalar { return Scalar(math.Floor(float64(a))) }
func (a Vec2) Floor() Vec2     { return Vec2{a.X.Floor(), a.Y.Floor()} }
func (a Vec3) Floor() Vec3     { return Vec3{a.X.Floor(), a.Y.Floor(), a.Z.Floor()} }
func (a Vec4) Floor() Vec4     { return Vec4{a.X.Floor(), a.Y.Floor(), a.Z.Floor(), a.W.Floor()} }

func (a Scalar) Ceil() Scalar { return Scalar(math.Ceil(float64(a))) }
func (a Vec2) Ceil() Vec2     { return Vec2{a.X.Ceil(), a.Y.Ceil()} }
func (a Vec3) Ceil() Vec3     { return Vec3{a.X.Ceil(), a.Y.Ceil(), a.Z.Ceil()} }
func (a Vec4) Ceil() Vec4     { return Vec4{a.X.Ceil(), a.Y.Ceil(), a.Z.Ceil(), a.W.Ceil()} }

func (a Scalar) Trunc() Scalar { return Scalar(math.Trunc(float64(a))) }
func (a Vec2) Trunc() Vec2     { return Vec2{a.X.Trunc(), a.Y.Trunc()} }
func (a Vec3) Trunc() Vec3     { return Vec3{a.X.Trunc(), a.Y.Trunc(), a.Z.Trunc()} }
func (a Vec4) Trunc() Vec4     { return Vec4{a.X.Trunc(), a.Y.Trunc(), a.Z.Trunc(), a.W.Trunc()} }

func (a Scalar) Fract() Scalar { return a - a.Floor() }
func (a Vec2) Fract() Vec2     { return Vec2{a.X.Fract(), a.Y.Fract()} }
func (a Vec3) Fract() Vec3     { return Vec3{a.X.Fract(), a.Y.Fract(), a.Z.Fract()} }
func (a Vec4) Fract() Vec4     { return Vec4{a.X.Fract(), a.Y.Fract(), a.Z.Fract(), a.W.Fract()} }

func (a Scalar) Mod(b Scalar) Scalar { return Scalar(math.Mod(float64(a), float64(b))) }
func (a Vec2) Mod(b Vec2) Vec2       { return Vec2{a.X.Mod(b.X), a.Y.Mod(b.Y)} }
func (a Vec3) Mod(b Vec3) Vec3       { return Vec3{a.X.Mod(b.X), a.Y.Mod(b.Y), a.Z.Mod(b.Z)} }
func (a Vec4) Mod(b Vec4) Vec4       { return Vec4{a.X.Mod(b.X), a.Y.Mod(b.Y), a.Z.Mod(b.Z), a.W.Mod(b.W)} }

func (a Vec2) ModAll(b Scalar) Vec2 { return Vec2{a.X.Mod(b), a.Y.Mod(b)} }
func (a Vec3) ModAll(b Scalar) Vec3 { return Vec3{a.X.Mod(b), a.Y.Mod(b), a.Z.Mod(b)} }
func (a Vec4) ModAll(b Scalar) Vec4 { return Vec4{a.X.Mod(b), a.Y.Mod(b), a.Z.Mod(b), a.W.Mod(b)} }

func (a Scalar) Min(b Scalar) Scalar { return Scalar(math.Min(float64(a), float64(b))) }
func (a Vec2) Min(b Vec2) Vec2       { return Vec2{a.X.Min(b.X), a.Y.Min(b.Y)} }
func (a Vec3) Min(b Vec3) Vec3       { return Vec3{a.X.Min(b.X), a.Y.Min(b.Y), a.Z.Min(b.Z)} }
func (a Vec4) Min(b Vec4) Vec4       { return Vec4{a.X.Min(b.X), a.Y.Min(b.Y), a.Z.Min(b.Z), a.W.Min(b.W)} }

func (a Vec2) MinAll(b Scalar) Vec2 { return Vec2{a.X.Min(b), a.Y.Min(b)} }
func (a Vec3) MinAll(b Scalar) Vec3 { return Vec3{a.X.Min(b), a.Y.Min(b), a.Z.Min(b)} }
func (a Vec4) MinAll(b Scalar) Vec4 { return Vec4{a.X.Min(b), a.Y.Min(b), a.Z.Min(b), a.W.Min(b)} }

func (a Scalar) Max(b Scalar) Scalar { return Scalar(math.Max(float64(a), float64(b))) }
func (a Vec2) Max(b Vec2) Vec2       { return Vec2{a.X.Max(b.X), a.Y.Max(b.Y)} }
func (a Vec3) Max(b Vec3) Vec3       { return Vec3{a.X.Max(b.X), a.Y.Max(b.Y), a.Z.Max(b.Z)} }
func (a Vec4) Max(b Vec4) Vec4       { return Vec4{a.X.Max(b.X), a.Y.Max(b.Y), a.Z.Max(b.Z), a.W.Max(b.W)} }

func (a Vec2) MaxAll(b Scalar) Vec2 { return Vec2{a.X.Max(b), a.Y.Max(b)} }
func (a Vec3) MaxAll(b Scalar) Vec3 { return Vec3{a.X.Max(b), a.Y.Max(b), a.Z.Max(b)} }
func (a Vec4) MaxAll(b Scalar) Vec4 { return Vec4{a.X.Max(b), a.Y.Max(b), a.Z.Max(b), a.W.Max(b)} }

func clamp[T ~float64](x, minval, maxval T) T {
	return min(maxval, max(minval, x))
}

func (a Scalar) Clamp(min, max Scalar) Scalar { return Scalar(clamp(a, min, max)) }
func (a Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2{a.X.Clamp(min.X, max.X), a.Y.Clamp(min.Y, max.Y)}
}
func (a Vec3) Clamp(min, max Vec3) Vec3 {
	return Vec3{a.X.Clamp(min.X, max.X), a.Y.Clamp(min.Y, max.Y), a.Z.Clamp(min.Z, max.Z)}
}
func (a Vec4) Clamp(min, max Vec4) Vec4 {
	return Vec4{a.X.Clamp(min.X, max.X), a.Y.Clamp(min.Y, max.Y), a.Z.Clamp(min.Z, max.Z), a.W.Clamp(min.W, max.W)}
}

func (a Vec2) ClampAll(min, max Scalar) Vec2 {
	return Vec2{a.X.Clamp(min, max), a.Y.Clamp(min, max)}
}
func (a Vec3) ClampAll(min, max Scalar) Vec3 {
	return Vec3{a.X.Clamp(min, max), a.Y.Clamp(min, max), a.Z.Clamp(min, max)}
}
func (a Vec4) ClampAll(min, max Scalar) Vec4 {
	return Vec4{a.X.Clamp(min, max), a.Y.Clamp(min, max), a.Z.Clamp(min, max), a.W.Clamp(min, max)}
}

func mix[T ~float64](x, y, a T) T {
	return x*(1-a) + y*a
}

func (a Scalar) Mix(b, t Scalar) Scalar { return Scalar(mix(a, b, t)) }
func (a Vec2) Mix(b, t Vec2) Vec2       { return Vec2{a.X.Mix(b.X, t.X), a.Y.Mix(b.Y, t.Y)} }
func (a Vec3) Mix(b, t Vec3) Vec3 {
	return Vec3{a.X.Mix(b.X, t.X), a.Y.Mix(b.Y, t.Y), a.Z.Mix(b.Z, t.Z)}
}
func (a Vec4) Mix(b, t Vec4) Vec4 {
	return Vec4{a.X.Mix(b.X, t.X), a.Y.Mix(b.Y, t.Y), a.Z.Mix(b.Z, t.Z), a.W.Mix(b.W, t.W)}
}

func (a Vec2) MixAll(b Vec2, t Scalar) Vec2 { return Vec2{a.X.Mix(b.X, t), a.Y.Mix(b.Y, t)} }
func (a Vec3) MixAll(b Vec3, t Scalar) Vec3 {
	return Vec3{a.X.Mix(b.X, t), a.Y.Mix(b.Y, t), a.Z.Mix(b.Z, t)}
}
func (a Vec4) MixAll(b Vec4, t Scalar) Vec4 {
	return Vec4{a.X.Mix(b.X, t), a.Y.Mix(b.Y, t), a.Z.Mix(b.Z, t), a.W.Mix(b.W, t)}
}

func step[T ~float64](edge, x T) T {
	if x < edge {
		return 0
	}
	return 1
}

func (a Scalar) Step(edge Scalar) Scalar { return Scalar(step(edge, a)) }
func (a Vec2) Step(edge Vec2) Vec2       { return Vec2{a.X.Step(edge.X), a.Y.Step(edge.Y)} }
func (a Vec3) Step(edge Vec3) Vec3       { return Vec3{a.X.Step(edge.X), a.Y.Step(edge.Y), a.Z.Step(edge.Z)} }
func (a Vec4) Step(edge Vec4) Vec4 {
	return Vec4{a.X.Step(edge.X), a.Y.Step(edge.Y), a.Z.Step(edge.Z), a.W.Step(edge.W)}
}

func (a Vec2) StepAll(edge Scalar) Vec2 { return Vec2{a.X.Step(edge), a.Y.Step(edge)} }
func (a Vec3) StepAll(edge Scalar) Vec3 { return Vec3{a.X.Step(edge), a.Y.Step(edge), a.Z.Step(edge)} }
func (a Vec4) StepAll(edge Scalar) Vec4 {
	return Vec4{a.X.Step(edge), a.Y.Step(edge), a.Z.Step(edge), a.W.Step(edge)}
}

func smoothstep[T ~float64](edge0, edge1, x T) T {
	t := clamp((x-edge0)/(edge1-edge0), 0, 1)
	return t * t * (3 - 2*t)
}

func (a Scalar) SmoothStep(edge0, edge1 Scalar) Scalar { return Scalar(smoothstep(edge0, edge1, a)) }
func (a Vec2) SmoothStep(edge0, edge1 Vec2) Vec2 {
	return Vec2{a.X.SmoothStep(edge0.X, edge1.X), a.Y.SmoothStep(edge0.Y, edge1.Y)}
}
func (a Vec3) SmoothStep(edge0, edge1 Vec3) Vec3 {
	return Vec3{a.X.SmoothStep(edge0.X, edge1.X), a.Y.SmoothStep(edge0.Y, edge1.Y), a.Z.SmoothStep(edge0.Z, edge1.Z)}
}
func (a Vec4) SmoothStep(edge0, edge1 Vec4) Vec4 {
	return Vec4{a.X.SmoothStep(edge0.X, edge1.X), a.Y.SmoothStep(edge0.Y, edge1.Y), a.Z.SmoothStep(edge0.Z, edge1.Z), a.W.SmoothStep(edge0.W, edge1.W)}
}

func (a Vec2) SmoothStepAll(edge0, edge1 Scalar) Vec2 {
	return Vec2{a.X.SmoothStep(edge0, edge1), a.Y.SmoothStep(edge0, edge1)}
}
func (a Vec3) SmoothStepAll(edge0, edge1 Scalar) Vec3 {
	return Vec3{a.X.SmoothStep(edge0, edge1), a.Y.SmoothStep(edge0, edge1), a.Z.SmoothStep(edge0, edge1)}
}
func (a Vec4) SmoothStepAll(edge0, edge1 Scalar) Vec4 {
	return Vec4{a.X.SmoothStep(edge0, edge1), a.Y.SmoothStep(edge0, edge1), a.Z.SmoothStep(edge0, edge1), a.W.SmoothStep(edge0, edge1)}
}

func (a Vec2) Length() Scalar { return Scalar(math.Sqrt(float64(a.LengthSq()))) }
func (a Vec3) Length() Scalar { return Scalar(math.Sqrt(float64(a.LengthSq()))) }
func (a Vec4) Length() Scalar { return Scalar(math.Sqrt(float64(a.LengthSq()))) }

func (a Vec2) Distance(b Vec2) Scalar { return a.Sub(b).Length() }
func (a Vec3) Distance(b Vec3) Scalar { return a.Sub(b).Length() }
func (a Vec4) Distance(b Vec4) Scalar { return a.Sub(b).Length() }

func (a Vec2) Normalize() Vec2 { return a.DivAll(a.Length()) }
func (a Vec3) Normalize() Vec3 { return a.DivAll(a.Length()) }
func (a Vec4) Normalize() Vec4 { return a.DivAll(a.Length()) }

func (a Vec2) Dot(b Vec2) Scalar { return a.X*b.X + a.Y*b.Y }
func (a Vec3) Dot(b Vec3) Scalar { return a.X*b.X + a.Y*b.Y + a.Z*b.Z }
func (a Vec4) Dot(b Vec4) Scalar { return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W }

func (a Vec2) Cross(b Vec2) Scalar { return a.X*b.Y - a.Y*b.X }
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func (a Vec2) Reflect(n Vec2) Vec2 { return a.Sub(n.MulAll(2 * a.Dot(n))) }
func (a Vec3) Reflect(n Vec3) Vec3 { return a.Sub(n.MulAll(2 * a.Dot(n))) }
func (a Vec4) Reflect(n Vec4) Vec4 { return a.Sub(n.MulAll(2 * a.Dot(n))) }

// ========================================
// Misc
// ========================================

func (a Vec2) LengthSq() Scalar { return a.X*a.X + a.Y*a.Y }
func (a Vec3) LengthSq() Scalar { return a.X*a.X + a.Y*a.Y + a.Z*a.Z }
func (a Vec4) LengthSq() Scalar { return a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W }

func (a Vec2) Polar() Vec2 { return Xy(a.Length(), a.Y.Atan2(a.X)) }

func (a Vec2) Rect() Vec2 {
	s, c := math.Sincos(float64(a.Y))
	return Xy(a.X*Scalar(c), a.X*Scalar(s))
}

func (a Scalar) Remap(inMin, inMax, outMin, outMax Scalar) Scalar {
	return outMin + (a-inMin)*(outMax-outMin)/(inMax-inMin)
}
func (a Vec2) Remap(inMin, inMax, outMin, outMax Vec2) Vec2 {
	return Vec2{a.X.Remap(inMin.X, inMax.X, outMin.X, outMax.X), a.Y.Remap(inMin.Y, inMax.Y, outMin.Y, outMax.Y)}
}
func (a Vec3) Remap(inMin, inMax, outMin, outMax Vec3) Vec3 {
	return Vec3{a.X.Remap(inMin.X, inMax.X, outMin.X, outMax.X), a.Y.Remap(inMin.Y, inMax.Y, outMin.Y, outMax.Y), a.Z.Remap(inMin.Z, inMax.Z, outMin.Z, outMax.Z)}
}
func (a Vec4) Remap(inMin, inMax, outMin, outMax Vec4) Vec4 {
	return Vec4{a.X.Remap(inMin.X, inMax.X, outMin.X, outMax.X), a.Y.Remap(inMin.Y, inMax.Y, outMin.Y, outMax.Y), a.Z.Remap(inMin.Z, inMax.Z, outMin.Z, outMax.Z), a.W.Remap(inMin.W, inMax.W, outMin.W, outMax.W)}
}

func (a Vec2) RemapAll(inMin, inMax, outMin, outMax Scalar) Vec2 {
	return Vec2{a.X.Remap(inMin, inMax, outMin, outMax), a.Y.Remap(inMin, inMax, outMin, outMax)}
}
func (a Vec3) RemapAll(inMin, inMax, outMin, outMax Scalar) Vec3 {
	return Vec3{a.X.Remap(inMin, inMax, outMin, outMax), a.Y.Remap(inMin, inMax, outMin, outMax), a.Z.Remap(inMin, inMax, outMin, outMax)}
}
func (a Vec4) RemapAll(inMin, inMax, outMin, outMax Scalar) Vec4 {
	return Vec4{a.X.Remap(inMin, inMax, outMin, outMax), a.Y.Remap(inMin, inMax, outMin, outMax), a.Z.Remap(inMin, inMax, outMin, outMax), a.W.Remap(inMin, inMax, outMin, outMax)}
}

func (a Scalar) OneMinus() Scalar { return 1 - a }
func (a Vec2) OneMinus() Vec2     { return Vec2{1 - a.X, 1 - a.Y} }
func (a Vec3) OneMinus() Vec3     { return Vec3{1 - a.X, 1 - a.Y, 1 - a.Z} }
func (a Vec4) OneMinus() Vec4     { return Vec4{1 - a.X, 1 - a.Y, 1 - a.Z, 1 - a.W} }
