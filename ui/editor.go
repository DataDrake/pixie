package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Editor struct {
	box *Box
}

func NewEditor(x, y int, sprite *Sprite) *Editor {
	var ed Editor
	ed.box = NewBox(sprite)
	ed.box.SetBorder(color.Gray{0x77})
	ed.box.SetMargin(1)
	ed.box.SetPadding(1)
	ed.box.SetPosition(x, y)
	return &ed
}

func (e *Editor) Bounds() image.Rectangle {
	return e.box.Bounds()
}

func (e *Editor) Draw(screen *ebiten.Image) {
	e.box.Draw(screen)
}

func (e *Editor) SetPosition(x, y int) {
	e.box.SetPosition(x, y)
}

func (e *Editor) SetSize(sx, sy int) {
	e.box.SetSize(sx, sy)
}

func (e *Editor) PreferredSize() (sx, sy int) {
	return e.box.PreferredSize()
}

func (e *Editor) SetVisible(visible bool) {
	e.box.SetVisible(visible)
}

func (e *Editor) Update() error {
	return e.box.Update()
}
