package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Preview struct {
	box *Box
}

func NewPreview(x, y int, s *Sprite, palette *color.Palette) *Preview {
	var prev Preview
	p := NewSprite(16, 2, palette)
	p.img = s.img
	prev.box = NewBox(p)
	prev.box.SetBorder(color.Gray{0x77})
	prev.box.SetMargin(1)
	prev.box.SetPadding(1)
	prev.box.SetPosition(x, y)
	return &prev
}

func (p *Preview) Bounds() image.Rectangle {
	return p.box.Bounds()
}

func (p *Preview) Draw(screen *ebiten.Image) {
	p.box.Draw(screen)
}

func (p *Preview) SetPosition(x, y int) {
	p.box.SetPosition(x, y)
}

func (p *Preview) SetSize(sx, sy int) {
	p.box.SetSize(sx, sy)
}

func (p *Preview) PreferredSize() (sx, sy int) {
	return p.box.PreferredSize()
}

func (p *Preview) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

func (p *Preview) Update() error {
	return nil
}
