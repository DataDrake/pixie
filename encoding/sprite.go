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
	"encoding/hex"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type spriteJSON struct {
	Size   int      `json:"size"`
	Pixels []string `json:"pixels"`
}

// Sprite is a square image made up of Palettizes Colors
type Sprite struct {
	img image.Paletted
}

// LoadSprite reads a single sprite from a JSON file and decodes it
func LoadSprite(path string) (s Sprite, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&s)
	return
}

// Convert translates a Sprite to an Ebiten Image for rendering
func (s *Sprite) Convert() *ebiten.Image {
	return ebiten.NewImageFromImage(&s.img)
}

// MarshalJSON is a custom marshaler for the Sprite type
//
// Sprites are encoded with a size (NxN) and an array of strings, where each string is a row of pixels, and each pixel is a 2 character hex value
func (s Sprite) MarshalJSON() (bs []byte, err error) {
	j := spriteJSON{
		Size: s.img.Stride,
	}
	var row []byte
	for i := 0; i < len(s.img.Pix); i += s.img.Stride {
		row = []byte(s.img.Pix[i*s.img.Stride : (i+1)*s.img.Stride])
		j.Pixels = append(j.Pixels, hex.EncodeToString(row))
	}
	return json.Marshal(j)
}

// UnmarshalJSON is a custom unmarshaler for the Sprite type
func (s *Sprite) UnmarshalJSON(b []byte) (err error) {
	var j spriteJSON
	if err = json.Unmarshal(b, &j); err != nil {
		return
	}
	var pix []byte
	for _, row := range j.Pixels {
		if pix, err = hex.DecodeString(row); err != nil {
			return
		}
		s.img.Pix = append(s.img.Pix, []uint8(pix)...)
	}
	s.img.Stride = j.Size
	s.img.Rect = image.Rect(0, 0, j.Size, j.Size)
	return
}
