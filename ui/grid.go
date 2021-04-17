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

// Grid is a 2D matrix of Drawables
type Grid struct {
	rows, cols int
	x, y       int
	sx, sy     int
	children   []Drawable
}

// NewGrid creates a Grid of the specified rows and columns
func NewGrid(rows, cols int) *Grid {
	return &Grid{
		rows: rows,
		cols: cols,
	}
}

// Append adds a Child to the Grid
func (g *Grid) Append(child Drawable) {
	g.children = append(g.children, child)
	g.SetPosition(g.x, g.y)
}

// Bounds returns a rectangle indicating the boundaries of the Grid if completely filled
func (g *Grid) Bounds() image.Rectangle {
	return image.Rect(g.x, g.y, g.x+g.sx, g.y+g.sy)
}

// Draw renders the Grid to a screen
//
// Children are drawn row by row, from left to right, until no more rows or children are available
func (g *Grid) Draw(screen *ebiten.Image) {
	for _, child := range g.children {
		child.Draw(screen)
	}
}

// SetPosition changes the XY coordinate of the upper-left pixel
//
// Additionally, each child's position is modified, setting any extra children to invisible
func (g *Grid) SetPosition(x, y int) {
	var maxW, maxH int
	for _, child := range g.children {
		cx, cy := child.PreferredSize()
		if cx > maxW {
			maxW = cx
		}
		if cy > maxH {
			maxH = cy
		}
	}
	g.sx, g.sy = maxW*g.cols, maxH*g.rows
	var row, col int
	for _, child := range g.children {
		if row >= g.rows {
			child.SetVisible(false)
			continue
		}
		child.SetPosition(col*maxW+x, row*maxH+y)
		col++
		if col == g.cols {
			col, row = 0, row+1
		}
	}
}

// SetSize sets the width and heght of the Grid (NOT IMPLEMENTED)
func (g *Grid) SetSize(sx, sy int) {
	panic("set size not implemented")
}

// PreferredSize indicated the width and height of the Grid if completely filled
func (g *Grid) PreferredSize() (int, int) {
	return g.sx, g.sy
}

// SetVisible changes the visibility of the Grid and its Children
func (g *Grid) SetVisible(visible bool) {
	for _, child := range g.children {
		child.SetVisible(visible)
	}
	// When visible, hide extra Children
	if visible {
		g.SetPosition(g.x, g.y)
	}
}

// Update calls Update for each child, one at a time
func (g *Grid) Update() error {
	for _, child := range g.children {
		if err := child.Update(); err != nil {
			return err
		}
	}
	return nil
}

// Clear removes all the children
func (g *Grid) Clear() {
	g.children = make([]Drawable, 0)
}
