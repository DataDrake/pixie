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

// Toolbar is a grid of toolbar icons for carrying out specific actions
type Toolbar struct {
	grid *Grid
}

// NewToolbar creates a Toolbar from a SpriteSet and a Palette
func NewToolbar(x, y int, ss model.SpriteSet) *Toolbar {
	grid := NewGrid(2, 4)
	for _, s := range ss {
		sp := NewSprite(s, false, 1)
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

// Bounds returns a rectangle indicating the visible boundaries of the Toolbar
func (t *Toolbar) Bounds() image.Rectangle {
	return t.grid.Bounds()
}

// Draw renders the Toolbar to a screen
func (t *Toolbar) Draw(screen *ebiten.Image) {
	t.grid.Draw(screen)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (t *Toolbar) SetPosition(x, y int) {
	t.grid.SetPosition(x, y)
}

// SetSize changes the desired width and height
func (t *Toolbar) SetSize(sx, sy int) {
	t.grid.SetSize(sx, sy)
}

// PreferredSize calculates the desired width and height of the Toolbar
func (t *Toolbar) PreferredSize() (sx, sy int) {
	return t.grid.PreferredSize()
}

// SetVisible changes the visibility of the Toolbar
func (t *Toolbar) SetVisible(visible bool) {
	t.grid.SetVisible(visible)
}

// Update checks if a toolbar button is pressed (NOT IMPLEMENTED)
func (t *Toolbar) Update() error {
	return t.grid.Update()
}
