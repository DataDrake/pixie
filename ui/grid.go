package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Grid struct {
	rows, cols int
	x, y       int
	sx, sy     int
	children   []Drawable
}

func NewGrid(rows, cols int) *Grid {
	return &Grid{
		rows: rows,
		cols: cols,
	}
}

func (g *Grid) Append(child Drawable) {
	g.children = append(g.children, child)
	g.SetPosition(g.x, g.y)
}

func (g *Grid) Bounds() image.Rectangle {
	return image.Rect(g.x, g.y, g.x+g.sx, g.y+g.sy)
}

func (g *Grid) Draw(screen *ebiten.Image) {
	for _, child := range g.children {
		child.Draw(screen)
	}
}

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

func (g *Grid) SetSize(sx, sy int) {
	panic("set size not implemented")
}

func (g *Grid) PreferredSize() (int, int) {
	return g.sx, g.sy
}

func (g *Grid) SetVisible(visible bool) {
	for _, child := range g.children {
		child.SetVisible(visible)
	}
}

func (g *Grid) Update() error {
	for _, child := range g.children {
		if err := child.Update(); err != nil {
			return err
		}
	}
	return nil
}
