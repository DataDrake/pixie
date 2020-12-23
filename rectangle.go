package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Rectangle struct {
	x, y    int
	sx, sy  int
	visible bool
	img     *ebiten.Image
}

func NewRectangle(sx, sy int) *Rectangle {
	r := &Rectangle{
		visible: true,
		img:     ebiten.NewImage(1, 1),
	}
	r.SetSize(sx, sy)
	return r
}

func (r *Rectangle) Bounds() image.Rectangle {
	return r.img.Bounds()
}

func (r *Rectangle) Draw(screen *ebiten.Image) {
	if r.visible {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(r.sx), float64(r.sy))
		op.GeoM.Translate(float64(r.x), float64(r.y))
		screen.DrawImage(r.img, op)
	}
}

func (r *Rectangle) SetColor(c color.Color) {
	r.img.Fill(c)
}

func (r *Rectangle) SetPosition(x, y int) {
	r.x, r.y = x, y
}

func (r *Rectangle) SetSize(sx, sy int) {
	if sx <= 0 || sy <= 0 {
		r.SetVisible(false)
		r.sx, r.sy = 0, 0
		return
	}
	r.sx, r.sy = sx, sy
}

func (r *Rectangle) PreferredSize() (int, int) {
	return r.sx, r.sy
}

func (r *Rectangle) SetVisible(visible bool) {
	r.visible = visible
}

func (r *Rectangle) Update() error {
	return nil
}
