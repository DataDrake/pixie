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

// ToolBar is a toolbar for file operation related to SpriteSets
type ToolBar struct {
	bar *ui.Toolbar
}

// NewToolBar creates a new ToolBar for the Sprite Editor
func NewToolBar(x, y int, colors *model.Palette) *ToolBar {

	icons, err := model.LoadSpriteSet(asset("sprite_toolbar.json"), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	icons.Describe()
	icons.SetPalette(colors)

	sb := ui.NewToolbar(x, y)
	sb.Append(ui.NewButton(icons.Sprites[0], 1, func(_ ebiten.MouseButton) {})) // Add
	sb.Append(ui.NewButton(icons.Sprites[1], 1, func(_ ebiten.MouseButton) {})) // Duplicate
	sb.Append(ui.NewButton(icons.Sprites[2], 1, func(_ ebiten.MouseButton) {})) // Clear
	sb.Append(ui.NewButton(icons.Sprites[3], 1, func(_ ebiten.MouseButton) {})) // Remove
	sb.Append(ui.NewButton(icons.Sprites[4], 1, func(_ ebiten.MouseButton) {})) // First
	sb.Append(ui.NewButton(icons.Sprites[5], 1, func(_ ebiten.MouseButton) {})) // Left
	sb.Append(ui.NewButton(icons.Sprites[6], 1, func(_ ebiten.MouseButton) {})) // Right
	sb.Append(ui.NewButton(icons.Sprites[7], 1, func(_ ebiten.MouseButton) {})) // Last

	return &ToolBar{
		bar: sb,
	}
}

// Update checks for all updates in the internal toolbar
func (b *ToolBar) Update() error {
	return b.bar.Update()
}

// Draw renders ToolBar to a screen
func (b *ToolBar) Draw(screen *ebiten.Image) {
	b.bar.Draw(screen)
}
