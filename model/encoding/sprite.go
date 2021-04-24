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
	"image"
)

// Sprite is the JSON representation of a Pixie sprite
type Sprite struct {
	Size   int      `json:"size"`
	Pixels []string `json:"pixels"`
}

// NewSprite creates a Sprite from a paletted image
// Sprites are encoded with a size (NxN) and an array of strings, where each string is a row of pixels, and each pixel is a 2 character hex value
func NewSprite(img *image.Paletted) (s Sprite) {
	s.Size = img.Stride
	var row []byte
	for i := 0; i < len(img.Pix); i += img.Stride {
		row = []byte(img.Pix[i*img.Stride : (i+1)*img.Stride])
		s.Pixels = append(s.Pixels, hex.EncodeToString(row))
	}
	return
}

// Image decodes this sprite as a paletted image
func (s Sprite) Image() (img *image.Paletted, err error) {
	img = &image.Paletted{
		Stride: s.Size,
		Rect:   image.Rect(0, 0, s.Size, s.Size),
	}
	var pix []byte
	for _, row := range s.Pixels {
		if pix, err = hex.DecodeString(row); err != nil {
			return
		}
		img.Pix = append(img.Pix, []uint8(pix)...)
	}
	return
}
