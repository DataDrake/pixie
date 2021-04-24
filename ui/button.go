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
	"github.com/DataDrake/pixie/model"
	"github.com/DataDrake/pixie/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Button is a square button for performing actions
type Button struct {
	x, y     int
	size     int
	scale    int
	visible  bool
	src      *model.Sprite
	img      *ebiten.Image
	callback ButtonHandler
}

// ButtonHandler is a callback to be used when a mouse button is clicked
type ButtonHandler func(btn ebiten.MouseButton)

// NewButton creates a new empty Sprite of the specified size, scale (multiplier), and ButtonHandler
func NewButton(img *model.Sprite, scale int, callback ButtonHandler) *Button {
	size := img.Bounds().Dx()
	b := &Button{
		size:     size,
		scale:    scale,
		visible:  true,
		src:      img,
		img:      ebiten.NewImage(size, size),
		callback: callback,
	}
	_, bg := img.Palette().BG()
	b.img.Fill(bg)
	return b
}

// Bounds returns a rectangle indicating the visible boundaries of the Sprite
func (b *Button) Bounds() image.Rectangle {
	return image.Rect(b.x, b.y, b.x+(b.size*b.scale), b.y+(b.size*b.scale))
}

// Draw renders the Sprite to a screen
func (b *Button) Draw(screen *ebiten.Image) {
	if !b.visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(b.scale), float64(b.scale))
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, op)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (b *Button) SetPosition(x, y int) {
	b.x, b.y = x, y
}

// SetSize changes the width and height of the Sprite
func (b *Button) SetSize(sx, sy int) {
	if sx != sy {
		panic(fmt.Sprintf("Sprites must have square dimensions, %d != %d", sx, sy))
	}
	b.size = sx / b.scale
}

// PreferredSize calculates the desired width and height of the Sprite
func (b *Button) PreferredSize() (int, int) {
	return b.size * b.scale, b.size * b.scale
}

// SetVisible changes the visibility of the Sprite
func (b *Button) SetVisible(visible bool) {
	b.visible = visible
}

// Update detects a mouse click inside a sprite, changing the color according to the button pressed
func (b *Button) Update() error {
	if b.src.HasChanged() {
		b.src.Update()
		b.img = ebiten.NewImageFromImage(b.src.Image())
	}
	if !b.visible {
		return nil
	}
	cx, cy := ebiten.CursorPosition()
	if util.In(b.Bounds(), cx, cy) {
		cx, cy = (cx-b.x)/b.scale, (cy-b.y)/b.scale
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.callback(ebiten.MouseButtonLeft)
		} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			b.callback(ebiten.MouseButtonRight)
		} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
			b.callback(ebiten.MouseButtonMiddle)
		}
	}
	return nil
}
