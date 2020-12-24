package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Selector struct {
	grid *Grid
}

func NewSelector(x, y int, s *Sprite, palette *color.Palette) *Selector {
	grid := NewGrid(8, 4)
	for j := 0; j < 34; j++ {
		sp := NewSprite(16, 1, palette)
		sp.img = s.img
		sb := NewBox(sp)
		sb.SetBorder(color.Gray{0x77})
		sb.SetMargin(1)
		sb.SetPadding(1)
		grid.Append(sb)
	}
	grid.SetPosition(x, y)
	return &Selector{
		grid: grid,
	}
}

func (s *Selector) Bounds() image.Rectangle {
	return s.grid.Bounds()
}

func (s *Selector) Draw(screen *ebiten.Image) {
	s.grid.Draw(screen)
}

func (s *Selector) SetPosition(x, y int) {
	s.grid.SetPosition(x, y)
}

func (s *Selector) SetSize(sx, sy int) {
	s.grid.SetSize(sx, sy)
}

func (s *Selector) PreferredSize() (sx, sy int) {
	return s.grid.PreferredSize()
}

func (s *Selector) SetVisible(visible bool) {
	s.grid.SetVisible(visible)
}

func (s *Selector) Update() error {
	return s.grid.Update()
}
