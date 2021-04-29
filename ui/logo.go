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

package ui

import (
	"github.com/DataDrake/pixie/assets"
	"github.com/DataDrake/pixie/model"
	"image"
	"log"
	"os"
)

var logos []image.Image

// GetLogo returns the Pixie logo in multiple sizes
func GetLogo() []image.Image {
	if len(logos) > 0 {
		return logos
	}
	colors, err := model.LoadPalette(assets.UI("palette.json"))
	if err != nil {
		log.Fatal(err)
	}
	sprites, err := model.LoadSpriteSet(assets.UI("logo.json"), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	sprites.SetPalette(colors)
	logo := sprites.Sprites[0].Image()
	logos = append(logos, logo)
	sizes := []int{32, 48, 64}
	for _, size := range sizes {
		logos = append(logos, resize(logo, size))
	}
	return logos
}

func resize(src *image.Paletted, stride int) (dst *image.Paletted) {
	dst = &image.Paletted{
		Stride:  stride,
		Rect:    image.Rect(0, 0, stride, stride),
		Palette: src.Palette,
	}
	scale := stride / src.Stride
	bounds := src.Bounds()
	idx := 0
	for y := 0; y < bounds.Dy(); y++ {
		var row []uint8
		for x := 0; x < bounds.Dx(); x++ {
			for s := 0; s < scale; s++ {
				row = append(row, src.Pix[idx])
			}
			idx++
		}
		for s := 0; s < scale; s++ {
			dst.Pix = append(dst.Pix, row...)
		}
	}
	return
}
