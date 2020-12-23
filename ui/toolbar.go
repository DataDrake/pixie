package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Toolbar struct {
	grid []*Box
}

func NewToolbar(x, y int, s *Sprite, palette *color.Palette) *Toolbar {
	var tb Toolbar
	for j := 0; j < 4; j++ {
		for i := 0; i < 2; i++ {
			sp := NewSprite(16, 1, palette)
			sp.img = s.img
			sb := NewBox(sp)
			sb.SetBorder(color.Gray{0x77})
			sb.SetMargin(1)
			sb.SetPadding(1)
			sb.SetPosition(x+j*20, y+i*20)
			tb.grid = append(tb.grid, sb)
		}
	}
	return &tb
}

func (t *Toolbar) Bounds() image.Rectangle {
	return t.grid[0].Bounds()
}

func (t *Toolbar) Draw(screen *ebiten.Image) {
	for _, b := range t.grid {
		b.Draw(screen)
	}
}

func (t *Toolbar) SetPosition(x, y int) {
	panic("position not implemented")
}

func (t *Toolbar) SetSize(sx, sy int) {
	panic("size not implement")
}

func (t *Toolbar) PreferredSize() (sx, sy int) {
	return t.grid[0].PreferredSize()
}

func (t *Toolbar) SetVisible(visible bool) {
	panic("visible not implemented")
}

func (t *Toolbar) Update() error {
	return nil
}
