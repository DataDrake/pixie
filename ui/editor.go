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
	"github.com/DataDrake/pixie/model"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Editor allows for pixel-by-pixel modification of a Sprite
type Editor struct {
	box *Box
}

// NewEditor creates an Editor at the desired location for the specified Sprite
func NewEditor(x, y int, s *model.Sprite) *Editor {
	var ed Editor
	sp := NewSprite(s, true, 16)
	ed.box = NewBox(sp)
	ed.box.SetBorder(color.Gray{0x77})
	ed.box.SetMargin(1)
	ed.box.SetPadding(1)
	ed.box.SetPosition(x, y)
	return &ed
}

// Bounds returns a Rectangle indicating the bounding box for the Editor
func (e *Editor) Bounds() image.Rectangle {
	return e.box.Bounds()
}

// Draw renders the Editor to a screen
func (e *Editor) Draw(screen *ebiten.Image) {
	e.box.Draw(screen)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (e *Editor) SetPosition(x, y int) {
	e.box.SetPosition(x, y)
}

// SetSize changes the width and height of the Editor
func (e *Editor) SetSize(sx, sy int) {
	e.box.SetSize(sx, sy)
}

// PreferredSize returns the desired width and height of the Editor
func (e *Editor) PreferredSize() (sx, sy int) {
	return e.box.PreferredSize()
}

// SetVisible changes the visibility of the Editor
func (e *Editor) SetVisible(visible bool) {
	e.box.SetVisible(visible)
}

// Update tells the internal Sprite to Update
func (e *Editor) Update() error {
	return e.box.Update()
}
