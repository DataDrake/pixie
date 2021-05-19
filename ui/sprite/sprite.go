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

package sprite

import (
	"fmt"
	"github.com/DataDrake/pixie/files"
	"github.com/DataDrake/pixie/ui"
	"github.com/DataDrake/pixie/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Sprite represents a square sprite of pixels
type Sprite struct {
	x, y     int
	size     int
	scale    int
	visible  bool
	selected ui.Selected
	writable bool
	src      *files.Sprite
	img      *ebiten.Image
}

// NewSprite creates a new empty Sprite of the specified size, scale (multiplier), and color Palette
func NewSprite(img *files.Sprite, writable bool, scale int) *Sprite {
	size := img.Bounds().Dx()
	s := &Sprite{
		size:     size,
		scale:    scale,
		visible:  true,
		writable: writable,
		src:      img,
		img:      ebiten.NewImage(size, size),
	}
	_, bg := img.Palette().BG()
	s.img.Fill(bg)
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
func (s *Sprite) Swap(src *files.Sprite) (prev *files.Sprite) {
	prev, s.src = s.src, src
	s.img = ebiten.NewImageFromImage(s.src.Image())
	return
}

// SetVisible changes the visibility of the Sprite
func (s *Sprite) SetVisible(visible bool) {
	s.visible = visible
}

// Update detects a mouse click inside a sprite, changing the color according to the button pressed
func (s *Sprite) Update() error {
	s.selected = ui.UnSelected
	cx, cy := ebiten.CursorPosition()
	if util.In(s.Bounds(), cx, cy) {
		cx, cy = (cx-s.x)/s.scale, (cy-s.y)/s.scale
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if s.writable {
				s.src.SetFG(cx, cy)
			}
			s.selected = ui.LeftSelect
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			if s.writable {
				s.src.SetBG(cx, cy)
			}
			s.selected = ui.RightSelect
		}
	}
	if s.src.HasChanged() {
		s.img = ebiten.NewImageFromImage(s.src.Image())
	}
	return nil
}
