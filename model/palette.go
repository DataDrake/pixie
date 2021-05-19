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
	"github.com/DataDrake/pixie/assets"
	"github.com/DataDrake/pixie/files"
	"log"
)

var current *files.Palette

func init() {
	palette, err := files.LoadPalette(assets.DefaultPalette())
	if err != nil {
		log.Fatal(err)
	}
	palette.Describe()
	println()
	SetPalette(palette)
}

// SetPalette swaps the current Palette with a different one
func SetPalette(p *files.Palette) {
	if len(p.Colors) < 2 {
		log.Fatal("Palettes must have at least two colors")
	}
	p.SetFG(1)
	p.SetBG(0)
	current = p
}

// GetPalette returns the current Palette
func GetPalette() *files.Palette {
	return current
}
