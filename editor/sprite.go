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
	"path/filepath"
)

// Sprite is an editor for drawing Sprites and managing Sprite Sets
type Sprite struct {
	editor   *ui.Editor
	preview  *ui.Preview
	toolbar  *ui.Toolbar
	sprites  *ui.Selector
	colors   *ui.Palette
	cPreview *ui.ColorPreview
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

	editorIcons, err := model.LoadSpriteSet(asset("editor_toolbar.json"))
	if err != nil {
		log.Fatal(err)
	}
	editorIcons.Describe()
	editorIcons.SetPalette(editorColors)

	return &Sprite{
		editor:   ui.NewEditor(94, 8),
		preview:  ui.NewPreview(382, 226),
		toolbar:  ui.NewToolbar(8, 8, editorIcons),
		sprites:  ui.NewSelector(360, 58),
		colors:   ui.NewPalette(8, 58),
		cPreview: ui.NewColorPreview(30, 222),
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
	if err := s.sprites.Update(); err != nil {
		return err
	}
	if err := s.colors.Update(); err != nil {
		return err
	}
	if err := s.cPreview.Update(); err != nil {
		return err
	}
	return nil
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
	s.sprites.Draw(screen)
	s.colors.Draw(screen)
	s.cPreview.Draw(screen)
}
