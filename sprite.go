package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Sprite struct {
	x, y    int
	size    int
	scale   int
	visible bool
	img     *ebiten.Image
	colors  *color.Palette
}

func NewSprite(size, scale int, colors *color.Palette) *Sprite {
	s := &Sprite{
		size:    size,
		scale:   scale,
		visible: true,
		img:     ebiten.NewImage(size, size),
		colors:  colors,
	}
	s.img.Fill((*s.colors)[0])
	return s
}

func (s *Sprite) Bounds() image.Rectangle {
	return image.Rect(s.x, s.y, s.x+(s.size*s.scale), s.y+(s.size*s.scale))
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(s.scale), float64(s.scale))
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(s.img, op)
}

func (s *Sprite) SetPosition(x, y int) {
	s.x, s.y = x, y
}

func (s *Sprite) SetSize(sx, sy int) {
	if sx != sy {
		panic(fmt.Sprintf("Sprites must have square dimensions, %d != %d", sx, sy))
	}
	s.size = sx / s.scale
}

func (s *Sprite) PreferredSize() (int, int) {
	return s.size * s.scale, s.size * s.scale
}

func (s *Sprite) Swap(img *ebiten.Image) (prev *ebiten.Image) {
	prev, s.img = s.img, img
	return
}

func (s *Sprite) SetVisible(visible bool) {
	s.visible = visible
}

func (s *Sprite) Update() error {
	cx, cy := ebiten.CursorPosition()
	if !In(s.Bounds(), cx, cy) {
		return nil
	}
	cx, cy = (cx-s.x)/s.scale, (cy-s.y)/s.scale
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.img.Set(cx, cy, (*s.colors)[1])
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.img.Set(cx, cy, (*s.colors)[0])
	}
	return nil
}
