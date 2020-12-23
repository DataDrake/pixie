package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Drawables are objects which can be rendered to a screen and may have their own state
type Drawable interface {
	Bounds() image.Rectangle
	Draw(screen *ebiten.Image)
	SetPosition(x, y int)
	SetSize(sx, sy int)
	PreferredSize() (sx, sy int)
	SetVisible(visible bool)
	Update() error
}
