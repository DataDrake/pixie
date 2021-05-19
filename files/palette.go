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
	"fmt"
	"github.com/DataDrake/pixie/files/encoding"
	"image/color"
	"os"
	"text/tabwriter"
	"time"
)

// Palette is a color picker for a Palette of colors
type Palette struct {
	Name     string          `json:"name"`
	Author   string          `json:"author"`
	Date     time.Time       `json:"date"`
	Revision int             `json:"revision"`
	Colors   encoding.Colors `json:"colors"`
	fg       int
	bg       int
	changed  bool
}

// LoadPalette reads a Palette from a JSON file and unmarshals it
func LoadPalette(path string) (p *Palette, err error) {
	p = &Palette{}
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(p)
	return
}

// HasChanged indicates that this palette has undergone changes
func (p Palette) HasChanged() bool {
	return p.changed
}

// Update clears the changed state
func (p *Palette) Update() error {
	p.changed = false
	return nil
}

// FG returns the index and color of the foreground
func (p Palette) FG() (index int, value color.Color) {
	return p.fg, p.Colors[p.fg]
}

// BG returns the index and color of the background
func (p Palette) BG() (index int, value color.Color) {
	return p.bg, p.Colors[p.bg]
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

// Describe privides a simple description of the Palette with its Metadata
func (p *Palette) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", p.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", p.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", p.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", p.Revision)
	fmt.Fprintf(tw, "Number of Colors\t: %d\n", len(p.Colors))
	tw.Flush()
}
