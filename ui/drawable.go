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
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Drawable is an object which can be rendered to a screen and may have its own state
type Drawable interface {
	// Bounds returns a rectangle indicating the visible boundaries of the Drawable
	Bounds() image.Rectangle
	// Draw renders this Drawable to a screen
	Draw(screen *ebiten.Image)
	// SetPosition sets the XY coordinate of the upper-left pixel
	SetPosition(x, y int)
	// SetSize changes the width and height of the Drawable
	SetSize(sx, sy int)
	// PreferredSize calculates the desired width and height of the Drawable
	PreferredSize() (sx, sy int)
	// SetVisible changes the visible of the Drawable
	SetVisible(visible bool)
	// Update changes the internal state of the Drawable and/or its Children
	Update() error
}
