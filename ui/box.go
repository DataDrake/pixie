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
	"image/color"
)

// Box is a rectangle which may have padding and/or a colored border, as well as a single Child
type Box struct {
	x, y    int
	margin  int
	padding int
	visible bool
	border  *Rectangle
	child   Drawable
}

// NewBox creates a Box containing the specified Child
func NewBox(child Drawable) *Box {
	b := &Box{
		visible: true,
		border:  NewRectangle(child.PreferredSize()),
		child:   child,
	}
	b.SetSize(child.PreferredSize())
	return b
}

// Bounds returnsa Rectangle indicating the visual boundaries of the Box
func (b *Box) Bounds() image.Rectangle {
	sx, sy := b.PreferredSize()
	return image.Rect(b.x, b.y, b.x+sx, b.y+sy)
}

// Draw renders a Box to a screen
func (b *Box) Draw(screen *ebiten.Image) {
	b.border.Draw(screen)
	b.child.Draw(screen)
}

// SetBorder changes the color of the border
func (b *Box) SetBorder(c color.Color) {
	b.border.SetColor(c)
}

// SetMargin changes the margin around the Box
func (b *Box) SetMargin(margin int) {
	if margin < 0 {
		margin = 0
	}
	b.margin = margin
	b.SetPosition(b.x, b.y)
	b.SetSize(b.child.PreferredSize())
}

// SetPadding changes the width of the Border
func (b *Box) SetPadding(padding int) {
	if padding < 0 {
		padding = 0
	}
	b.padding = padding
	b.SetPosition(b.x, b.y)
	b.SetSize(b.child.PreferredSize())
}

// SetPosition changes the XY corrdinate of the upper-left pixel of the Box
func (b *Box) SetPosition(x, y int) {
	b.x, b.y = x, y
	b.border.SetPosition(x+b.margin, y+b.margin)
	b.child.SetPosition(x+b.margin+b.padding, y+b.margin+b.padding)
}

// SetSize changes the width and height of the Box
func (b *Box) SetSize(sx, sy int) {
	if sx <= 0 || sy <= 0 {
		b.SetVisible(false)
		b.child.SetSize(0, 0)
		return
	}
	b.border.SetSize((2*b.padding)+sx, (2*b.padding)+sy)
	b.child.SetSize(sx, sy)
}

// PreferredSize calculates the width and height of the box including the margin and padding
func (b *Box) PreferredSize() (int, int) {
	cx, cy := b.child.PreferredSize()
	return (2 * (b.margin + b.padding)) + cx, (2 * (b.margin + b.padding)) + cy
}

// SetVisible changes the visibility of the Box
func (b *Box) SetVisible(visible bool) {
	b.visible = visible
	b.border.SetVisible(visible)
	b.child.SetVisible(visible)
}

// Update tells the child to update
func (b *Box) Update() error {
	return b.child.Update()
}

// Child give access to a Box's child
func (b *Box) Child() Drawable {
	return b.child
}
