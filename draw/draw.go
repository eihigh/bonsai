package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	. "github.com/eihigh/bonsai/geometry"
)

type Texture = ebiten.Image

type Vertex = ebiten.Vertex

type Pipeline struct {
	// TODO: 各種オプションを持たせる
}

type Helper struct {
	p      *Pipeline
	dsts   []*Texture
	srcs   []*Texture
	tf     func(Vertex) Vertex
	affine ebiten.GeoM
	color  color.Color
}

func NewHelper(p *Pipeline, dst *Texture, srcs ...*Texture) *Helper {
	return &Helper{p: p, dsts: []*Texture{dst}, srcs: srcs}
}

func (h *Helper) Draw() {
	// TODO
}

func (h *Helper) Translate(p Vec2) {
	if h.tf != nil {
		panic("do not use affine and tf at the same time")
	}
	h.affine.Translate(p.Xy())
}

// type Mesh struct {
// 	AllVertices iter.Seq[Vertex]
// 	AllIndices  iter.Seq[uint16]
// }
//
// type Drawer struct {
// 	// 便利な描画メソッドを実現するためによく使う値をとにかくぶち込む
//
// 	p    *Pipeline
// 	dsts []*Texture
// 	srcs []*Texture
// 	mesh Mesh // nil なら src の領域を元形状として使う
// 	mat3 Mat3
// 	clr  color.Color
// }
//
// func (p *Pipeline) Draw(dsts, srcs []*Texture, mesh Mesh) {
// 	vs := make([]Vertex, 4)
// 	iter.Append(vs, mesh.AllVertices)
// 	is := make([]uint16, 6)
// 	iter.Append(is, mesh.AllIndices)
// 	opts := &ebiten.DrawTrianglesOptions{}
// 	// TODO: p のオプションを opts に反映
// 	dsts[0].DrawTriangles(vs, is, srcs[0], opts)
// }
