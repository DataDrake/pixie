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
	"github.com/DataDrake/pixie/assets"
	"github.com/DataDrake/pixie/files"
	"github.com/DataDrake/pixie/model"
	"github.com/DataDrake/pixie/ui/color"
	"github.com/DataDrake/pixie/ui/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"path/filepath"
)

// Editor is for drawing Sprites and managing Sprite Sets
type Editor struct {
	file     *FileBar
	colors   *color.Palette
	swatches *color.Preview
	editor   *sprite.Editor
	preview  *sprite.Preview
	tools    *ToolBar
	sprites  *sprite.Selector
}

// asset calculates a path relative to the Sprite Editor assets
func asset(path string) string {
	return assets.UI(filepath.Join("sprite", path))
}

// NewEditor creates a new Sprite Editor and populates its GUI
func NewEditor() *Editor {

	colors, err := files.LoadPalette(assets.UI("palette.json"))
	if err != nil {
		log.Fatal(err)
	}
	colors.Describe()

	return &Editor{
		file:     NewFileBar(8, 8, colors),
		colors:   color.NewPalette(8, 58),
		swatches: color.NewPreview(30, 222),
		editor:   sprite.NewEditor(94, 8),
		tools:    NewToolBar(360, 8, colors),
		sprites:  sprite.NewSelector(360, 58),
		preview:  sprite.NewPreview(382, 226),
	}
}

// Update checks for all updates in the internal state of the Sprite Editor
func (e *Editor) Update() error {
	// Update View
	if err := e.file.Update(); err != nil {
		return err
	}
	if err := e.colors.Update(); err != nil {
		return err
	}
	if err := e.swatches.Update(); err != nil {
		return err
	}
	if err := e.editor.Update(); err != nil {
		return err
	}
	if err := e.tools.Update(); err != nil {
		return err
	}
	if err := e.sprites.Update(); err != nil {
		return err
	}
	if err := e.preview.Update(); err != nil {
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
func (e *Editor) Save() error {
	return model.SaveSprites()
}

// Exit checks for unsaved state and starts the process of shutting down
func (e *Editor) Exit() (done bool, err error) {
	done = true
	return
}

// Draw renders Sprite Editor to a screen
func (e *Editor) Draw(screen *ebiten.Image) {
	e.file.Draw(screen)
	e.colors.Draw(screen)
	e.swatches.Draw(screen)
	e.editor.Draw(screen)
	e.tools.Draw(screen)
	e.sprites.Draw(screen)
	e.preview.Draw(screen)
}
