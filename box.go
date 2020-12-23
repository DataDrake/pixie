package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Box struct {
	x, y    int
	sx, sy  int
	margin  int
	padding int
	visible bool
	border  *Rectangle
	child   Drawable
}

func NewBox(x, y, sx, sy int, child Drawable) *Box {
	b := &Box{
		margin:  0,
		padding: 0,
		visible: true,
		border:  NewRectangle(x, y, sx, sy),
		child:   child,
	}
	b.SetPosition(x, y)
	b.SetSize(sx, sy)
	return b
}

func (b *Box) Bounds() image.Rectangle {
	return image.Rect(b.x, b.y, b.x+b.sx, b.y+b.sy)
}

func (b *Box) Draw(screen *ebiten.Image) {
	b.border.Draw(screen)
	b.child.Draw(screen)
}

func (b *Box) SetBorder(c color.Color) {
	b.border.SetColor(c)
}

func (b *Box) SetMargin(margin int) {
	if margin < 0 {
		margin = 0
	}
	b.margin = margin
	b.SetPosition(b.x, b.y)
	b.SetSize(b.sx, b.sy)
}

func (b *Box) SetPadding(padding int) {
	if padding < 0 {
		padding = 0
	}
	b.padding = padding
	b.SetPosition(b.x, b.y)
	b.SetSize(b.sx, b.sy)
}

func (b *Box) SetPosition(x, y int) {
	b.x, b.y = x, y
	b.border.SetPosition(x+b.margin, y+b.margin)
	b.child.SetPosition(x+b.margin+b.padding, y+b.margin+b.padding)
}

func (b *Box) SetSize(sx, sy int) {
	if sx <= 0 || sy <= 0 {
		b.SetVisible(false)
		b.sx, b.sy = 0, 0
		b.child.SetSize(0, 0)
		return
	}
	b.sx, b.sy = sx, sy
	b.border.SetSize(sx-(2*b.margin), sy-(2*b.margin))
	b.child.SetSize(sx-(2*b.margin)-(2*b.padding), sy-(2*b.margin)-(2*b.padding))
}

func (b *Box) SetVisible(visible bool) {
	b.visible = visible
	b.border.SetVisible(visible)
	b.child.SetVisible(visible)
}

func (b *Box) Update() error {
	return b.child.Update()
}
