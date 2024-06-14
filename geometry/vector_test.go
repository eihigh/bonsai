package geometry

func Example() {
	// Original code:
	/*
		func curve(uv vec2) vec2 {
			uv = (uv - 0.5) * 2
			uv *= 1.1
			uv.x *= (1 + pow((abs(uv.y)/14), 2))
			uv.y *= (1 + pow((abs(uv.x)/12), 2))
			uv = uv*0.5 + 0.5
			uv = uv*0.92 + 0.04

			return uv
		}

		func Fragment(dst vec4, src vec2, color vec4) vec4 {
			origin, size := imageSrcRegionOnTexture()
			q := (src - origin) / size
			uv := q
			uv = curve(uv)

			var col vec3
			col.r = imageSrc0At(vec2(uv.x+0.0001, uv.y+0.0001)*size+origin).x + 0.0025
			col.g = imageSrc0At(vec2(uv.x+0.0000, uv.y-0.0002)*size+origin).y + 0.0025
			col.b = imageSrc0At(vec2(uv.x-0.0002, uv.y+0.0000)*size+origin).z + 0.0025
			col.r += 0.04 * imageSrc0At((0.75*vec2(0.025, -0.027)+vec2(uv.x+0.001, uv.y+0.001))*size+origin).x
			col.g += 0.025 * imageSrc0At((0.75*vec2(-0.022, -0.02)+vec2(uv.x+0.000, uv.y-0.002))*size+origin).y
			col.b += 0.04 * imageSrc0At((0.75*vec2(-0.02, -0.018)+vec2(uv.x-0.002, uv.y+0.000))*size+origin).z

			col = clamp(col*0.6+0.4*col*col, 0, 1)

			vig := (40.0 * uv.x * uv.y * (1 - uv.x) * (1 - uv.y))
			col *= vec3(pow(vig, 0.3))
			col *= vec3(0.95, 1.05, 0.95)
			col *= 2.4

			scans := clamp(0.35+0.35*sin(uv.y*size.y*1.5), 0, 1)
			s := pow(scans, 3.7)
			col *= vec3(0.45 + 0.1*s)

			if uv.x < 0.0 || uv.x > 1.0 || uv.y < 0 || uv.y > 1 {
				col *= 0
			}

			col *= (1.0 - 0.25*vec3(clamp((mod(src.x, 2)-1)*2, 0, 1)))

			return vec4(col, 1) * 1.05
		}
	*/

	// Converted code:
	// dst := Vec4{}
	src := Vec2{}
	// color := Vec4{}

	origin, size := imageSrcRegionOnTexture()
	q := src.Sub(origin).Div(size)
	uv := q
	uv = curve(uv)

	col := Vec3{}
	col.X = imageSrc0At(uv.Add(Xy(0.0001, 0.0001)).Mul(size).Add(origin)).X + 0.0025
	col.Y = imageSrc0At(uv.Add(Xy(0, -0.0002).Mul(size).Add(origin))).Y + 0.0025
	col.Z = imageSrc0At(uv.Add(Xy(-0.0002, 0).Mul(size).Add(origin))).Z + 0.0025
	col.X += 0.04 * imageSrc0At(Xy(0.75*0.025, -0.027).Add(uv.Add(Xy(0.001, 0.001)).Mul(size)).Add(origin)).X
	col.Y += 0.025 * imageSrc0At(Xy(0.75*-0.022, -0.02).Add(uv.Add(Xy(0, -0.002)).Mul(size)).Add(origin)).Y
	col.Z += 0.04 * imageSrc0At(Xy(0.75*-0.02, -0.018).Add(uv.Add(Xy(-0.002, 0)).Mul(size)).Add(origin)).Z

	col = col.Scale(0.6).Add(col.Mul(col).Scale(0.4)).ClampAll(0, 1)

	vig := 40.0 * uv.X * uv.Y * (1 - uv.X) * (1 - uv.Y)
	col = col.Scale(vig.Pow(0.3))
	col = col.Mul(Xyz(0.95, 1.05, 0.95))
	col = col.Scale(2.4)

	scans := 0.35 + 0.35*(uv.Y*size.Y*1.5).Sin().Clamp(0, 1)
	s := scans.Pow(3.7)
	col = col.Scale(0.45 + 0.1*s)

	if uv.X < 0.0 || uv.X > 1.0 || uv.Y < 0 || uv.Y > 1 {
		col = col.Scale(0)
	}

	col = col.Scale(src.X.Mod(2).Sub(1).Mul(2).Clamp(0, 1)*-0.25 + 1)
	// col = col.Scale(src.X.Mod(2).Remap(0, 2, 1, 0.75))

	// GLSLに親しんだ人が、converted codeを読んで驚く点：
	// 1. ほとんどの演算子がメソッドに置き換えられている
	// 2. 命名規則が、キャメルケースからスネークケースに変わっている
	// 3. ベクトルの要素にアクセスするために、.X, .Y, .Z, .Wが使われている
	// 4. ベクトルの要素にアクセスするために、.Xyz(), .Xy(), .Xyzw()が使われている
	// 5. 一部の関数名が、略語から完全な単語に変わっている
	// 6. 四則演算の優先順位を明確にするために、括弧が使われている

}

func imageSrcRegionOnTexture() (Vec2, Vec2) {
	return Vec2{}, Vec2{}
}

func imageSrc0At(v Vec2) Vec4 {
	return v.Vec4(0, 0)
}

func curve(uv Vec2) Vec2 {
	uv = uv.SubAll(0.5).MulAll(2.2)
	uv.X *= uv.Y.Div(14).Abs().Pow(2).Add(1)
	uv.Y *= uv.X.Div(12).Abs().Pow(2).Add(1)
	uv = uv.Scale(0.5).AddAll(0.5)
	uv = uv.Scale(0.92).AddAll(0.04)

	return uv
}
