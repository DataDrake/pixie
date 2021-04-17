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

// ColorPreview displays the currently selected FG and BG colors
type ColorPreview struct {
	x, y    int
	fg      *Box
	bg      *Box
	palette *model.Palette
}

// NewColorPreview creates a new ColorPreview with the specified size and colors
func NewColorPreview(x, y int, palette *model.Palette) *ColorPreview {
	prev := &ColorPreview{
		x:       x,
		y:       y,
		palette: palette,
	}

	_, fg := palette.FG()
	fgS := NewSwatch(16, fg)
	prev.fg = NewBox(fgS)
	prev.fg.SetBorder(color.Gray{0x77})
	prev.fg.SetPadding(1)

	_, bg := palette.BG()
	bgS := NewSwatch(16, bg)
	prev.bg = NewBox(bgS)
	prev.bg.SetBorder(color.Gray{0x77})
	prev.bg.SetPadding(1)

	prev.SetPosition(x, y)
	return prev
}

// Bounds returns a Rectangle defining the visible area occupied by the ColorPreview
func (p *ColorPreview) Bounds() image.Rectangle {
	cx, cy := p.PreferredSize()
	return image.Rect(p.x, p.y, p.x+cx, p.y+cy)
}

// Draw renders the ColorPreview to a screen
func (p *ColorPreview) Draw(screen *ebiten.Image) {
	p.bg.Draw(screen)
	p.fg.Draw(screen)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (p *ColorPreview) SetPosition(x, y int) {
	cx, cy := p.fg.PreferredSize()
	p.bg.SetPosition(x+cx, y+cy)
	p.fg.SetPosition(x, y)
}

// SetSize changes the width and height of the ColorPreview (NOT IMPLEMENTED)
func (p *ColorPreview) SetSize(sx, sy int) {
	panic("set size not implemented")
}

// PreferredSize calculated the desired height and width of the ColorPreview
func (p *ColorPreview) PreferredSize() (sx, sy int) {
	cx, cy := p.fg.PreferredSize()
	return 2 * cx, 2 * cy
}

// SetVisible changes the visibility of the ColorPreview
func (p *ColorPreview) SetVisible(visible bool) {
	p.fg.SetVisible(visible)
	p.bg.SetVisible(visible)
}

// Update calls Update for the FG and BG swatches
func (p *ColorPreview) Update() error {
	if p.palette.HasChanged() {
		_, fg := p.palette.FG()
		_, bg := p.palette.BG()
		p.fg.child.(*Swatch).Swap(fg)
		p.bg.child.(*Swatch).Swap(bg)
	}
	if err := p.fg.Update(); err != nil {
		return err
	}
	return p.bg.Update()
}
