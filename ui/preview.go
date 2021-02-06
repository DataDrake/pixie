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

// Preview provides an enlarged view of a Sprite
type Preview struct {
	box *Box
}

// NewPreview creates a new Preview for the specified Palette
func NewPreview(x, y int, s *Sprite, palette *color.Palette) *Preview {
	var prev Preview
	p := NewSprite(16, 2, palette)
	p.img = s.img
	prev.box = NewBox(p)
	prev.box.SetBorder(color.Gray{0x77})
	prev.box.SetMargin(1)
	prev.box.SetPadding(1)
	prev.box.SetPosition(x, y)
	return &prev
}

// Bounds returns a rectangle indicating the visible boundaries of the Preview
func (p *Preview) Bounds() image.Rectangle {
	return p.box.Bounds()
}

// Draw renders the Preview to a screen
func (p *Preview) Draw(screen *ebiten.Image) {
	p.box.Draw(screen)
}

// SetPosition changes the XY coordinate of the upper-left pixel
func (p *Preview) SetPosition(x, y int) {
	p.box.SetPosition(x, y)
}

// SetSize changes the width and height of the Preview
func (p *Preview) SetSize(sx, sy int) {
	p.box.SetSize(sx, sy)
}

// PreferredSize calculates the desired width and height of the Preview
func (p *Preview) PreferredSize() (sx, sy int) {
	return p.box.PreferredSize()
}

// SetVisible changes the visibility of the Preview
func (p *Preview) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

// Update checks for mouse clicks in the Preview (NOT IMPLEMENTED)
func (p *Preview) Update() error {
	return nil
}
