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
	"image/color"
	"log"
)

// Palette is a color picker for a Palette of colors
type Palette struct {
	colors  color.Palette
	fg      int
	bg      int
	changed bool
}

// NewPalette creates a Palette from existing colors
func NewPalette(colors color.Palette) *Palette {
	if len(colors) < 2 {
		log.Fatal("Palettes must have at least two colors")
	}
	return &Palette{
		colors:  colors,
		fg:      1,
		bg:      0,
		changed: true,
	}
}

// Clone creates a deep copy of this Palette
func (p *Palette) Clone() (clone *Palette) {
	clone = &Palette{
		colors:  p.colors,
		fg:      p.fg,
		bg:      p.bg,
		changed: true,
	}
	return
}

// HasChanged indicates that this palette has undergone changes
func (p *Palette) HasChanged() bool {
	return p.changed
}

// Update clears the changed state
func (p *Palette) Update() error {
	p.changed = false
	return nil
}

// FG returns the index and color of the foreground
func (p *Palette) FG() (index int, value color.Color) {
	return p.fg, p.colors[p.fg]
}

// BG returns the index and color of the background
func (p *Palette) BG() (index int, value color.Color) {
	return p.bg, p.colors[p.bg]
}

// Colors returns the color.Palette of this Palette
func (p *Palette) Colors() color.Palette {
	return p.colors
}

// SetFG changes the FG color
func (p *Palette) SetFG(index int) {
	if p.fg != index {
		p.fg = index
		p.changed = true
	}
}

// SetBG changes the BG color
func (p *Palette) SetBG(index int) {
	if p.bg != index {
		p.bg = index
		p.changed = true
	}
}

// SetColors changes the color.Palette of this Palette
func (p *Palette) SetColors(colors color.Palette) {
	p.colors = colors
	if p.fg >= len(p.colors) {
		p.SetFG(1)
	}
	if p.bg >= len(p.colors) {
		p.SetBG(0)
	}
	p.changed = true
}
