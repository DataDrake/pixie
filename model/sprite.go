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

package model

import (
	"image"
)

// Sprite holds the data for a single sprite
type Sprite struct {
	img     image.Paletted
	colors  *Palette
	changed bool
}

// NewSprite creates a sprite from image data and a palette
func NewSprite(img image.Paletted, colors *Palette) *Sprite {
	s := &Sprite{
		img: img,
	}
	s.SetPalette(colors)
	return s
}

// Image returns a pointer to the contained image data
func (s *Sprite) Image() *image.Paletted {
	return &s.img
}

// SetFG sets the pixel at the specified location to the FG color
func (s *Sprite) SetFG(x, y int) {
	fg, _ := s.colors.FG()
	s.img.SetColorIndex(x, y, uint8(fg))
	s.changed = true
}

// SetBG sets the pixel at the specified location to the BG color
func (s *Sprite) SetBG(x, y int) {
	bg, _ := s.colors.BG()
	s.img.SetColorIndex(x, y, uint8(bg))
	s.changed = true
}

// Bounds returns the bounding box for this sprite
func (s *Sprite) Bounds() image.Rectangle {
	return s.img.Bounds()
}

// Palette returns the color palette for this sprite
func (s *Sprite) Palette() *Palette {
	return s.colors
}

// SetPalette changes the color palette for this sprite
func (s *Sprite) SetPalette(colors *Palette) {
	s.colors = colors
	s.img.Palette = s.colors.Colors()
	s.changed = true
}

// Update checks for any internal changes for this sprite
func (s *Sprite) Update() error {
	s.changed = false
	if s.colors.HasChanged() {
		s.changed = true
	}
	return nil
}

// HasChanged reports if this Sprite or its Palette have changed and a redraw is needed
func (s *Sprite) HasChanged() bool {
	return s.changed || s.colors.HasChanged()
}
