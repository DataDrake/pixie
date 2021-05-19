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
	"os"
)

func init() {
	set, err := files.LoadSpriteSet(assets.DefaultSprites(), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	set.Describe()
	println()
	SetSprites(set)
	if len(os.Args) > 1 {
		_, err := os.Stat(os.Args[1])
		if os.IsNotExist(err) {
			f, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				log.Fatal(err)
			}
			next := spriteSet.Clone(f)
			next.Describe()
			println()
			SetSprites(next)
		} else if err != nil {
			log.Fatal(err)
		} else {
			next, err := files.LoadSpriteSet(os.Args[1], os.O_RDWR)
			if err != nil {
				log.Fatal(err)
			}
			set.Describe()
			println()
			SetSprites(next)
		}
	}
}

var spriteSet *files.SpriteSet

// GetSprites retrieves a list of all of the sprints in the current SpriteSet
func GetSprites() *files.SpriteSet {
	return spriteSet
}

// GetSprite retrieves a sprite from the current SpriteSet
func GetSprite(index int) *files.Sprite {
	return spriteSet.Sprites[index]
}

// SetSprites swaps out the current SpriteSet with a different one
func SetSprites(set *files.SpriteSet) {
	if spriteSet != nil {
		spriteSet.Close()
	}
	spriteSet = set
	spriteSet.SetPalette(GetPalette())
}

// SaveSprites saves out the currently open SpriteSet
func SaveSprites() error {
	return spriteSet.Save()
}
