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
	"encoding/json"
	"fmt"
	"github.com/DataDrake/pixie/assets"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

func init() {
	set, err := LoadSpriteSet(assets.DefaultSprites())
	if err != nil {
		log.Fatal(err)
	}
	set.Describe()
	println()
	SetSprites(set)
}

var spriteSet *SpriteSet

// GetSprites retrieves a list of all of the sprints in the current SpriteSet
func GetSprites() []*Sprite {
	return spriteSet.Sprites
}

// GetSprite retrieves a sprite from the current SpriteSet
func GetSprite(index int) *Sprite {
	return spriteSet.Sprites[index]
}

// SetSprites swaps out the current SpriteSet with a different one
func SetSprites(set *SpriteSet) {
	spriteSet = set
	spriteSet.SetPalette(GetPalette())
}

// SpriteSet represents one or more sprites belonging to a single set
type SpriteSet struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Sprites  []*Sprite `json:"sprites"`
	modified bool
}

// LoadSpriteSet reads in a SpriteSet for a JSON file and decodes it
func LoadSpriteSet(path string) (ss *SpriteSet, err error) {
	ss = &SpriteSet{
		modified: false,
	}
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(ss)
	return
}

// SetPalette sets the color Palette for this SpriteSet
func (ss *SpriteSet) SetPalette(p *Palette) {
	for i, s := range ss.Sprites {
		s.SetPalette(p)
		ss.Sprites[i] = s
	}
}

// Describe summarizes a SpriteSet according to its metadata
func (ss *SpriteSet) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", ss.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", ss.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", ss.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", ss.Revision)
	fmt.Fprintf(tw, "Number of Sprites\t: %d\n", len(ss.Sprites))
	tw.Flush()
}
