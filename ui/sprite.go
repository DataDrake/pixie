//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ui

import (
	"fmt"
	"github.com/DataDrake/pixie/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Sprite represents a square sprite of pixels
type Sprite struct {
	x, y    int
	size    int
	scale   int
	visible bool
	img     *ebiten.Image
	colors  *color.Palette
}

// NewSprite creates a new empty Sprite of the specified size, scale (multiplier), and color Palette
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

// Bounds returns a rectangle indicating the visible boundaries of the Sprite
func (s *Sprite) Bounds() image.Rectangle {
	return image.Rect(s.x, s.y, s.x+(s.size*s.scale), s.y+(s.size*s.scale))
}

// Draw renders the Sprite to a screen
func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(s.scale), float64(s.scale))
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(s.img, op)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (s *Sprite) SetPosition(x, y int) {
	s.x, s.y = x, y
}

// SetSize changes the width and height of the Sprite
func (s *Sprite) SetSize(sx, sy int) {
	if sx != sy {
		panic(fmt.Sprintf("Sprites must have square dimensions, %d != %d", sx, sy))
	}
	s.size = sx / s.scale
}

// PreferredSize calculates the desired width and height of the Sprite
func (s *Sprite) PreferredSize() (int, int) {
	return s.size * s.scale, s.size * s.scale
}

// Swap changes out the Image for this sprite with another
func (s *Sprite) Swap(img *ebiten.Image) (prev *ebiten.Image) {
	prev, s.img = s.img, img
	return
}

// SetVisible changes the visibility of the Sprite
func (s *Sprite) SetVisible(visible bool) {
	s.visible = visible
}

// Update detects a mouse click inside a sprite, changing the color according to the button pressed
func (s *Sprite) Update() error {
	cx, cy := ebiten.CursorPosition()
	if !util.In(s.Bounds(), cx, cy) {
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
