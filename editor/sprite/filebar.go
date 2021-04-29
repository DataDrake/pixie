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

package sprite

import (
	"github.com/DataDrake/pixie/model"
	"github.com/DataDrake/pixie/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"os"
)

// FileBar is a toolbar for file operation related to SpriteSets
type FileBar struct {
	bar *ui.Toolbar
}

// NewFileBar creates a new FileBar for the Sprite Editor
func NewFileBar(x, y int, colors *model.Palette) *FileBar {

	editorIcons, err := model.LoadSpriteSet(asset("editor_toolbar.json"), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	editorIcons.Describe()
	editorIcons.SetPalette(colors)

	tb := ui.NewToolbar(x, y)
	tb.Append(ui.NewButton(editorIcons.Sprites[0], 1, func(_ ebiten.MouseButton) {})) // New
	tb.Append(ui.NewButton(editorIcons.Sprites[1], 1, func(_ ebiten.MouseButton) {})) // Open
	tb.Append(ui.NewButton(editorIcons.Sprites[2], 1, func(btn ebiten.MouseButton) {
		if btn == ebiten.MouseButtonLeft {
			model.SaveSprites()
		}
	})) // Save
	tb.Append(ui.NewButton(editorIcons.Sprites[3], 1, func(_ ebiten.MouseButton) {})) // SaveAs
	tb.Append(ui.NewButton(editorIcons.Sprites[4], 1, func(_ ebiten.MouseButton) {})) // Export
	return &FileBar{
		bar: tb,
	}
}

// Update checks for all updates in the internal toolbar
func (b *FileBar) Update() error {
	return b.bar.Update()
}

// Draw renders FileBar to a screen
func (b *FileBar) Draw(screen *ebiten.Image) {
	b.bar.Draw(screen)
}
