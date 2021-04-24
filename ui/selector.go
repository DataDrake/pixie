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

// Selector is a 2D grid of Sprites
type Selector struct {
	grid *Grid
}

// NewSelector creates a Selector for the provided SpriteSet with the specified Palette
func NewSelector(x, y int) *Selector {
	grid := NewGrid(8, 4)
	for _, s := range model.GetSprites() {
		sp := NewSprite(s, false, 1)
		sb := NewBox(sp)
		sb.SetBorder(color.Gray{0x77})
		sb.SetMargin(1)
		sb.SetPadding(1)
		grid.Append(sb)
	}
	grid.SetPosition(x, y)
	return &Selector{
		grid: grid,
	}
}

// Bounds returns a rectangle indicatign the visible boundaries of the Selector
func (s *Selector) Bounds() image.Rectangle {
	return s.grid.Bounds()
}

// Draw renders the selector to the screen
func (s *Selector) Draw(screen *ebiten.Image) {
	s.grid.Draw(screen)
}

// SetPosition changes the XP coordinate of the upper-left pixel
func (s *Selector) SetPosition(x, y int) {
	s.grid.SetPosition(x, y)
}

// SetSize changes the width and height of the Selector
func (s *Selector) SetSize(sx, sy int) {
	s.grid.SetSize(sx, sy)
}

// PreferredSize calculates the desired width and height of the Selector
func (s *Selector) PreferredSize() (sx, sy int) {
	return s.grid.PreferredSize()
}

// SetVisible changes the visibility of the Selector
func (s *Selector) SetVisible(visible bool) {
	s.grid.SetVisible(visible)
}

// Update checks if any Sprites are selected
func (s *Selector) Update() error {
	return s.grid.Update()
}
