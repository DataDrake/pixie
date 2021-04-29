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
	"github.com/DataDrake/pixie/editor/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

// Kind defines the different types of Editor
type Kind int

const (
	// SpriteKind is an editor of Sprites and Sprite Sets
	SpriteKind Kind = iota
	// Max is the total number of types of Editor
	Max
)

// Editor is an interactive editor for some kind of graphics
type Editor interface {
	// Draw renders this Editor to a screen
	Draw(screen *ebiten.Image)
	// Update changes the internal state of the Editor and/or its Children
	Update() error
	// Save writes out and unsaved changes to the thing being edited
	Save() error
	// Exit deals with things like saving unsaved work
	Exit() (done bool, err error)
}

// All builds the available Editors
func All() (editors [Max]Editor) {
	editors[SpriteKind] = sprite.NewEditor()
	return
}
