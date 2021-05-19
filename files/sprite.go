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

package files

import (
	"encoding/json"
	"github.com/DataDrake/pixie/files/encoding"
	"image"
	"image/color"
)

// Sprite holds the data for a single sprite
type Sprite struct {
	img      *image.Paletted
	colors   *Palette
	modified bool
	changed  bool
}

// Clone creates a deep copy of this Sprite
func (s *Sprite) Clone() (next *Sprite) {
	next = &Sprite{
		img: &image.Paletted{
			Stride:  s.img.Stride,
			Rect:    s.img.Rect,
			Palette: s.img.Palette,
		},
		colors:   s.colors,
		modified: true,
		changed:  true,
	}
	next.img.Pix = s.img.Pix
	return
}

// SetFG sets the pixel at the specified location to the FG color
func (s *Sprite) SetFG(x, y int) {
	fg, _ := s.colors.FG()
	s.img.SetColorIndex(x, y, uint8(fg))
	s.modified = true
	s.changed = true
}

// SetBG sets the pixel at the specified location to the BG color
func (s *Sprite) SetBG(x, y int) {
	bg, _ := s.colors.BG()
	s.img.SetColorIndex(x, y, uint8(bg))
	s.modified = true
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

// Image returns the internal image as a pointer
func (s *Sprite) Image() *image.Paletted {
	return s.img
}

// SetPalette changes the color palette for this sprite
func (s *Sprite) SetPalette(colors *Palette) {
	s.colors = colors
	s.img.Palette = color.Palette(s.colors.Colors)
	s.changed = true
}

// IsModified reports if this Sprite has edited
func (s *Sprite) IsModified() bool {
	return s.modified
}

// HasChanged report if the Sprite has been edited or if the color palette has changed
func (s *Sprite) HasChanged() bool {
	return s.changed || s.colors.HasChanged()
}

// Update clears any change flags used in rendering
func (s *Sprite) Update() {
	s.changed = false
}

// MarshalJSON is a custom marshaler for the Sprite type
func (s Sprite) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(encoding.NewSprite(s.img))
}

// UnmarshalJSON is a custom unmarshaler for the Sprite type
func (s *Sprite) UnmarshalJSON(b []byte) (err error) {
	var j encoding.Sprite
	if err = json.Unmarshal(b, &j); err != nil {
		return
	}
	(*s).img, err = j.Image()
	(*s).modified = false
	(*s).changed = true
	return
}
