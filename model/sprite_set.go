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
	set, err := LoadSpriteSet(assets.DefaultSprites(), os.O_RDONLY)
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
			next, err := LoadSpriteSet(os.Args[1], os.O_RDWR)
			if err != nil {
				log.Fatal(err)
			}
			set.Describe()
			println()
			SetSprites(next)
		}
	}
}

var spriteSet *SpriteSet

// GetSprites retrieves a list of all of the sprints in the current SpriteSet
func GetSprites() *SpriteSet {
	return spriteSet
}

// GetSprite retrieves a sprite from the current SpriteSet
func GetSprite(index int) *Sprite {
	return spriteSet.Sprites[index]
}

// SetSprites swaps out the current SpriteSet with a different one
func SetSprites(set *SpriteSet) {
	if spriteSet != nil {
		spriteSet.Close()
	}
	spriteSet = set
	spriteSet.changed = true
	spriteSet.SetPalette(GetPalette())
}

// SaveSprites saves out the currently open SpriteSet
func SaveSprites() error {
	return spriteSet.Save()
}

// SpriteSet represents one or more sprites belonging to a single set
type SpriteSet struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Sprites  []*Sprite `json:"sprites"`
	changed  bool
	file     *os.File
}

// LoadSpriteSet reads in a SpriteSet for a JSON file and decodes it
func LoadSpriteSet(path string, flag int) (ss *SpriteSet, err error) {
	ss = &SpriteSet{}
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return
	}
	ss.file = f
	dec := json.NewDecoder(f)
	err = dec.Decode(ss)
	return
}

// Clone creates a deep copy of an existing SpriteSet for a new file
func (ss SpriteSet) Clone(file *os.File) (next *SpriteSet) {
	next = &SpriteSet{
		Name:     ss.Name,
		Author:   ss.Author,
		Date:     time.Now(),
		Revision: 1,
		changed:  true,
		file:     file,
	}
	for _, s := range ss.Sprites {
		next.Sprites = append(next.Sprites, s.Clone())
	}
	return
}

// Update marks the SpriteSet as no longer being changed
func (ss *SpriteSet) Update() error {
	ss.changed = false
	for i, s := range ss.Sprites {
		s.Update()
		ss.Sprites[i] = s
	}
	return nil
}

// Save writes out a SpriteSet to the file it was read from
func (ss *SpriteSet) Save() error {
	if !ss.IsModified() {
		return nil
	}
	// Seek to the beginning of the file
	_, err := ss.file.Seek(0, 0)
	if err != nil {
		return err
	}
	// clear the contents of the file
	err = ss.file.Truncate(0)
	if err != nil {
		return err
	}
	// write out the new contents to the file
	enc := json.NewEncoder(ss.file)
	enc.SetIndent("", "\t")
	if err := enc.Encode(ss); err != nil {
		return err
	}
	for i, s := range ss.Sprites {
		s.modified = false
		ss.Sprites[i] = s
	}
	return nil
}

// IsModified checks if any fo the sprites have been modified
func (ss *SpriteSet) IsModified() bool {
	for _, s := range ss.Sprites {
		if s.IsModified() {
			return true
		}
	}
	return false
}

// Close lets go of the file for this sprite set
func (ss *SpriteSet) Close() {
	ss.file.Close()
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
