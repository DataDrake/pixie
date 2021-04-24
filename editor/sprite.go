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

package editor

import (
	"github.com/DataDrake/pixie/assets"
	"github.com/DataDrake/pixie/model"
	"github.com/DataDrake/pixie/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"os"
	"path/filepath"
)

// Sprite is an editor for drawing Sprites and managing Sprite Sets
type Sprite struct {
	editor    *ui.Editor
	preview   *ui.Preview
	toolbar   *ui.Toolbar
	spriteBar *ui.Toolbar
	sprites   *ui.Selector
	colors    *ui.Palette
	cPreview  *ui.ColorPreview
}

// asset calculates a path relative to the Sprite Editor assets
func asset(path string) string {
	return assets.UI(filepath.Join("sprite", path))
}

// NewSprite creates a new Sprite Editor and populates its GUI
func NewSprite() *Sprite {

	editorColors, err := model.LoadPalette(assets.UI("palette.json"))
	if err != nil {
		log.Fatal(err)
	}
	editorColors.Describe()

	editorIcons, err := model.LoadSpriteSet(asset("editor_toolbar.json"), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	editorIcons.Describe()
	editorIcons.SetPalette(editorColors)

	spriteIcons, err := model.LoadSpriteSet(asset("sprite_toolbar.json"), os.O_RDONLY)
	if err != nil {
		log.Fatal(err)
	}
	spriteIcons.Describe()
	spriteIcons.SetPalette(editorColors)

	tb := ui.NewToolbar(8, 8)

	tb.Append(ui.NewButton(editorIcons.Sprites[0], 1, func(_ ebiten.MouseButton) {})) // New
	tb.Append(ui.NewButton(editorIcons.Sprites[1], 1, func(_ ebiten.MouseButton) {})) // Open
	tb.Append(ui.NewButton(editorIcons.Sprites[2], 1, func(btn ebiten.MouseButton) {
		if btn == ebiten.MouseButtonLeft {
			model.SaveSprites()
		}
	})) // Save
	tb.Append(ui.NewButton(editorIcons.Sprites[3], 1, func(_ ebiten.MouseButton) {})) // SaveAs
	tb.Append(ui.NewButton(editorIcons.Sprites[4], 1, func(_ ebiten.MouseButton) {})) // Export

	sb := ui.NewToolbar(360, 8)

	sb.Append(ui.NewButton(spriteIcons.Sprites[0], 1, func(_ ebiten.MouseButton) {})) // Add
	sb.Append(ui.NewButton(spriteIcons.Sprites[1], 1, func(_ ebiten.MouseButton) {})) // Duplicate
	sb.Append(ui.NewButton(spriteIcons.Sprites[2], 1, func(_ ebiten.MouseButton) {})) // Clear
	sb.Append(ui.NewButton(spriteIcons.Sprites[3], 1, func(_ ebiten.MouseButton) {})) // Remove
	sb.Append(ui.NewButton(spriteIcons.Sprites[4], 1, func(_ ebiten.MouseButton) {})) // First
	sb.Append(ui.NewButton(spriteIcons.Sprites[5], 1, func(_ ebiten.MouseButton) {})) // Left
	sb.Append(ui.NewButton(spriteIcons.Sprites[6], 1, func(_ ebiten.MouseButton) {})) // Right
	sb.Append(ui.NewButton(spriteIcons.Sprites[7], 1, func(_ ebiten.MouseButton) {})) // Last

	return &Sprite{
		editor:    ui.NewEditor(94, 8),
		preview:   ui.NewPreview(382, 226),
		toolbar:   tb,
		spriteBar: sb,
		sprites:   ui.NewSelector(360, 58),
		colors:    ui.NewPalette(8, 58),
		cPreview:  ui.NewColorPreview(30, 222),
	}
}

// Update checks for all updates in the internal state of the Sprite Editor
func (s *Sprite) Update() error {
	// Update View
	if err := s.editor.Update(); err != nil {
		return err
	}
	if err := s.preview.Update(); err != nil {
		return err
	}
	if err := s.toolbar.Update(); err != nil {
		return err
	}
	if err := s.spriteBar.Update(); err != nil {
		return err
	}
	if err := s.sprites.Update(); err != nil {
		return err
	}
	if err := s.colors.Update(); err != nil {
		return err
	}
	if err := s.cPreview.Update(); err != nil {
		return err
	}
	if err := model.GetPalette().Update(); err != nil {
		return err
	}
	if err := model.GetSprites().Update(); err != nil {
		return err
	}
	return nil
}

// Save writes out any changes to the current SpriteSet
func (s *Sprite) Save() error {
	return model.SaveSprites()
}

// Exit checks for unsaved state and starts the process of shutting down
func (s *Sprite) Exit() (done bool, err error) {
	done = true
	return
}

// Draw renders Sprite Editor to a screen
func (s *Sprite) Draw(screen *ebiten.Image) {
	s.editor.Draw(screen)
	s.preview.Draw(screen)
	s.toolbar.Draw(screen)
	s.spriteBar.Draw(screen)
	s.sprites.Draw(screen)
	s.colors.Draw(screen)
	s.cPreview.Draw(screen)
}
