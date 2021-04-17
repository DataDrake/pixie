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
	"github.com/DataDrake/pixie/encoding"
	"image"
)

// SpriteSet represents one or more sprites belonging to a single set
type SpriteSet []*Sprite

// NewSpriteSet creates a set from decoded image data and a Palette
func NewSpriteSet(ss encoding.SpriteSet, p *Palette) (set SpriteSet) {
	for _, s := range ss.Sprites {
		sp := NewSprite(image.Paletted(s), p)
		set = append(set, sp)
	}
	return
}

// SetPalette sets the color Palette for this SpriteSet
func (ss *SpriteSet) SetPalette(p *Palette) {
	for i, s := range *ss {
		s.SetPalette(p)
		(*ss)[i] = s
	}
}
