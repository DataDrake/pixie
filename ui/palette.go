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

// Palette is a color picker for a Palette of colors
type Palette struct {
	grid *Grid
}

// NewPalette creates a Palette at the specified location for the specified colors
func NewPalette(x, y int, palette *color.Palette) *Palette {
	grid := NewGrid(8, 4)
	for _, c := range *palette {
		sw := NewSwatch(16, c)
		sb := NewBox(sw)
		sb.SetBorder(color.Gray{0x77})
		sb.SetPadding(1)
		sb.SetMargin(1)
		grid.Append(sb)
	}
	grid.SetPosition(x, y)
	return &Palette{
		grid: grid,
	}
}

// Bounds returns a rectangle indicating the visible boundaries of the Palette
func (p *Palette) Bounds() image.Rectangle {
	return p.grid.Bounds()
}

// Draw renders the Palette to a screen
func (p *Palette) Draw(screen *ebiten.Image) {
	p.grid.Draw(screen)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (p *Palette) SetPosition(x, y int) {
	p.grid.SetPosition(x, y)
}

// SetSize changes the width and height
func (p *Palette) SetSize(sx, sy int) {
	p.grid.SetSize(sx, sy)
}

// PreferredSize calculates the desired width and height of the Palette
func (p *Palette) PreferredSize() (sx, sy int) {
	return p.grid.PreferredSize()
}

// SetVisible changes the visibility of the Palette
func (p *Palette) SetVisible(visible bool) {
	p.grid.SetVisible(visible)
}

// Update checks for any mouse clicks on the Palette
func (p *Palette) Update() error {
	return p.grid.Update()
}