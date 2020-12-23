package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Selector struct {
	grid []*Box
}

func NewSelector(x, y int, s *Sprite, palette *color.Palette) *Selector {
	var sel Selector
	for j := 0; j < 4; j++ {
		for i := 0; i < 8; i++ {
			sp := NewSprite(16, 1, palette)
			sp.img = s.img
			sb := NewBox(sp)
			sb.SetBorder(color.Gray{0x77})
			sb.SetMargin(1)
			sb.SetPadding(1)
			sb.SetPosition(x+j*20, y+i*20)
			sel.grid = append(sel.grid, sb)
		}
	}
	return &sel
}

func (s *Selector) Bounds() image.Rectangle {
	return s.grid[0].Bounds()
}

func (s *Selector) Draw(screen *ebiten.Image) {
	for _, b := range s.grid {
		b.Draw(screen)
	}
}

func (s *Selector) SetPosition(x, y int) {
	panic("position not implemented")
}

func (s *Selector) SetSize(sx, sy int) {
	panic("size not implement")
}

func (s *Selector) PreferredSize() (sx, sy int) {
	return s.grid[0].PreferredSize()
}

func (s *Selector) SetVisible(visible bool) {
	panic("visible not implemented")
}

func (s *Selector) Update() error {
	return nil
}
