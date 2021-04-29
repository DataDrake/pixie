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

package color

import (
	"fmt"
	"github.com/DataDrake/pixie/ui"
	"github.com/DataDrake/pixie/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Swatch is a color swatch
type Swatch struct {
	x, y     int
	size     int
	visible  bool
	selected ui.Selected
	img      *ebiten.Image
	color    color.Color
}

// NewSwatch creates a Swatch of the specified color
func NewSwatch(size int, color color.Color) *Swatch {
	s := &Swatch{
		size:    size,
		visible: true,
		img:     ebiten.NewImage(1, 1),
		color:   color,
	}
	s.img.Fill(color)
	return s
}

// Bounds returns a rectangle indicating the visible bounds of the Swatch
func (s *Swatch) Bounds() image.Rectangle {
	return image.Rect(s.x, s.y, s.x+s.size, s.y+s.size)
}

// Draw renders a Swatch to a screen
func (s *Swatch) Draw(screen *ebiten.Image) {
	if !s.visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(s.size), float64(s.size))
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(s.img, op)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (s *Swatch) SetPosition(x, y int) {
	s.x, s.y = x, y
}

// SetSize changes the width and height of the swatch
func (s *Swatch) SetSize(sx, sy int) {
	if sx != sy {
		panic(fmt.Sprintf("Swatchs must have square dimensions, %d != %d", sx, sy))
	}
	s.size = sx
}

// PreferredSize calculates the desired width and height of the Swatch
func (s *Swatch) PreferredSize() (int, int) {
	return s.size, s.size
}

// Swap changes the color of the Swatch
func (s *Swatch) Swap(color color.Color) (prev color.Color) {
	prev, s.color = s.color, color
	s.img.Fill(color)
	return
}

// SetVisible changes the visibility of the Swatch
func (s *Swatch) SetVisible(visible bool) {
	s.visible = visible
}

// Update checks for a mouse click inside the swatch (NOT IMPLEMENTED)
func (s *Swatch) Update() error {
	cx, cy := ebiten.CursorPosition()
	if !util.In(s.Bounds(), cx, cy) {
		s.selected = ui.UnSelected
		return nil
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.selected = ui.LeftSelect
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.selected = ui.RightSelect
	}
	return nil
}
