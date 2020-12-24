package ui

import (
	"github.com/DataDrake/pixie/encoding"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Toolbar struct {
	grid *Grid
}

func NewToolbar(x, y int, ss *encoding.SpriteSet, palette *color.Palette) *Toolbar {
	grid := NewGrid(2, 4)
	for _, s := range ss.Sprites {
		img := s.Convert()
		sz, _ := img.Size()
		sp := NewSprite(sz, 1, palette)
		sp.img = s.Convert()
		sb := NewBox(sp)
		sb.SetBorder(color.Gray{0x77})
		sb.SetMargin(1)
		sb.SetPadding(1)
		grid.Append(sb)
	}
	grid.SetPosition(x, y)
	return &Toolbar{
		grid: grid,
	}
}

func (t *Toolbar) Bounds() image.Rectangle {
	return t.grid.Bounds()
}

func (t *Toolbar) Draw(screen *ebiten.Image) {
	t.grid.Draw(screen)
}

func (t *Toolbar) SetPosition(x, y int) {
	t.grid.SetPosition(x, y)
}

func (t *Toolbar) SetSize(sx, sy int) {
	t.grid.SetSize(sx, sy)
}

func (t *Toolbar) PreferredSize() (sx, sy int) {
	return t.grid.PreferredSize()
}

func (t *Toolbar) SetVisible(visible bool) {
	t.grid.SetVisible(visible)
}

func (t *Toolbar) Update() error {
	return t.grid.Update()
}
