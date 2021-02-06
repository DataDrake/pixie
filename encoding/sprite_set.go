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

package encoding

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"text/tabwriter"
	"time"
)

// SpriteSet contains multiple sprites which belong together
type SpriteSet struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Sprites  []*Sprite `json:"sprites"`
	palette  *color.Palette
}

// LoadSpriteSet reads in a SpriteSet for a JSON file and decodes it
func LoadSpriteSet(path string) (ss SpriteSet, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&ss)
	return
}

// ChangePalette sets the color Palette for this SpriteSet
func (ss *SpriteSet) ChangePalette(p *color.Palette) {
	ss.palette = p
	for _, s := range ss.Sprites {
		s.img.Palette = *p
	}
}

// Palette is the color Palette used by this SpriteSet
func (ss *SpriteSet) Palette() *color.Palette {
	return ss.palette
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
