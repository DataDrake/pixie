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
	"github.com/DataDrake/pixie/files"
	"github.com/DataDrake/pixie/model"
	"github.com/DataDrake/pixie/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Palette is a color picker for a Palette of colors
type Palette struct {
	x, y   int
	colors *files.Palette
	grid   *ui.Grid
}

// NewPalette creates a Palette at the specified location for the specified colors
func NewPalette(x, y int) *Palette {
	grid := ui.NewGrid(8, 4)
	grid.SetPosition(x, y)
	return &Palette{
		x:      x,
		y:      y,
		colors: model.GetPalette(),
		grid:   grid,
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
	if p.colors.HasChanged() {
		p.grid.Clear()
		for _, c := range p.colors.Colors {
			sw := NewSwatch(16, c)
			sb := ui.NewBox(sw)
			sb.SetBorder(color.Gray{0x77})
			sb.SetPadding(1)
			sb.SetMargin(1)
			p.grid.Append(sb)
		}
		p.grid.SetPosition(p.x, p.y)
	}
	if err := p.grid.Update(); err != nil {
		return err
	}
	for i, child := range p.grid.Children() {
		b := child.(*ui.Box)
		s := b.Child().(*Swatch)
		switch s.selected {
		case ui.LeftSelect:
			p.colors.SetFG(i)
			break
		case ui.RightSelect:
			p.colors.SetBG(i)
			break
		default:
			continue
		}
	}
	return nil
}
