package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Box struct {
	x, y    int
	margin  int
	padding int
	visible bool
	border  *Rectangle
	child   Drawable
}

func NewBox(child Drawable) *Box {
	b := &Box{
		visible: true,
		border:  NewRectangle(child.PreferredSize()),
		child:   child,
	}
	b.SetSize(child.PreferredSize())
	return b
}

func (b *Box) Bounds() image.Rectangle {
	sx, sy := b.PreferredSize()
	return image.Rect(b.x, b.y, b.x+sx, b.y+sy)
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
	b.SetSize(b.child.PreferredSize())
}

func (b *Box) SetPadding(padding int) {
	if padding < 0 {
		padding = 0
	}
	b.padding = padding
	b.SetPosition(b.x, b.y)
	b.SetSize(b.child.PreferredSize())
}

func (b *Box) SetPosition(x, y int) {
	b.x, b.y = x, y
	b.border.SetPosition(x+b.margin, y+b.margin)
	b.child.SetPosition(x+b.margin+b.padding, y+b.margin+b.padding)
}

func (b *Box) SetSize(sx, sy int) {
	if sx <= 0 || sy <= 0 {
		b.SetVisible(false)
		b.child.SetSize(0, 0)
		return
	}
	b.border.SetSize((2*b.padding)+sx, (2*b.padding)+sy)
	b.child.SetSize(sx, sy)
}

func (b *Box) PreferredSize() (int, int) {
	cx, cy := b.child.PreferredSize()
	return (2 * (b.margin + b.padding)) + cx, (2 * (b.margin + b.padding)) + cy
}

func (b *Box) SetVisible(visible bool) {
	b.visible = visible
	b.border.SetVisible(visible)
	b.child.SetVisible(visible)
}

func (b *Box) Update() error {
	return b.child.Update()
}
