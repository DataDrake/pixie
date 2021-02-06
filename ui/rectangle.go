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
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Rectangle is a colored box
type Rectangle struct {
	x, y    int
	sx, sy  int
	visible bool
	img     *ebiten.Image
}

// NewRectangle creates a new Rectangle of the specified width and height
func NewRectangle(sx, sy int) *Rectangle {
	r := &Rectangle{
		visible: true,
		img:     ebiten.NewImage(1, 1),
	}
	r.SetSize(sx, sy)
	return r
}

// Bounds returns a rectangle indicating the visible area of the Rectangle
func (r *Rectangle) Bounds() image.Rectangle {
	return r.img.Bounds()
}

// Draw renders the Rectangle to a screen
func (r *Rectangle) Draw(screen *ebiten.Image) {
	if r.visible {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(r.sx), float64(r.sy))
		op.GeoM.Translate(float64(r.x), float64(r.y))
		screen.DrawImage(r.img, op)
	}
}

// SetColor changes the color of the Rectangle
func (r *Rectangle) SetColor(c color.Color) {
	r.img.Fill(c)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (r *Rectangle) SetPosition(x, y int) {
	r.x, r.y = x, y
}

// SetSize changes that width and height of the Rectangle
func (r *Rectangle) SetSize(sx, sy int) {
	if sx <= 0 || sy <= 0 {
		r.SetVisible(false)
		r.sx, r.sy = 0, 0
		return
	}
	r.sx, r.sy = sx, sy
}

// PreferredSize calculates the desired width and height of the Rectangle
func (r *Rectangle) PreferredSize() (int, int) {
	return r.sx, r.sy
}

// SetVisible changes the visibility of the Rectangle
func (r *Rectangle) SetVisible(visible bool) {
	r.visible = visible
}

// Update updates the state of the Rectangle (NOT IMPLEMENTED)
func (r *Rectangle) Update() error {
	return nil
}
