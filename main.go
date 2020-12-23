package main

import (
	"fmt"
	"github.com/DataDrake/pixie/encoding"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
	"time"
)

const (
	screenWidth  = 380
	screenHeight = 360
)

type Game struct {
	editor  *Box
	preview *Box
	grid    []*Box
	last    time.Time
}

func NewGame() *Game {
	pl, err := encoding.LoadPalette("examples/palette.json")
	if err != nil {
		panic(err)
	}
	pl.Describe()
	println()
	palette := color.Palette(pl.Colors)

	s := NewSprite(16, 16, &palette)
	editor := NewBox(s)
	editor.SetBorder(color.Gray{0x77})
	editor.SetMargin(1)
	editor.SetPadding(1)
	editor.SetPosition(88, 8)

	ss, err := encoding.LoadSpriteSet("examples/sprite_set.json")
	if err != nil {
		panic(err)
	}
	ss.Describe()
	ss.ChangePalette(&palette)
	s.Swap(ss.Sprites[0].Convert())

	p := NewSprite(16, 2, &palette)
	p.img = s.img
	preview := NewBox(p)
	preview.SetBorder(color.Gray{0x77})
	preview.SetMargin(1)
	preview.SetPadding(1)
	preview.SetPosition(30, 226)

	var grid []*Box
	for j := 0; j < 4; j++ {
		for i := 0; i < 2; i++ {
			sp := NewSprite(16, 1, &palette)
			sp.img = s.img
			sb := NewBox(sp)
			sb.SetBorder(color.Gray{0x77})
			sb.SetMargin(1)
			sb.SetPadding(1)
			sb.SetPosition(8+j*20, 8+i*20)
			grid = append(grid, sb)
		}
	}
	for j := 0; j < 4; j++ {
		for i := 0; i < 8; i++ {
			sp := NewSprite(16, 1, &palette)
			sp.img = s.img
			sb := NewBox(sp)
			sb.SetBorder(color.Gray{0x77})
			sb.SetMargin(1)
			sb.SetPadding(1)
			sb.SetPosition(8+j*20, 58+i*20)
			grid = append(grid, sb)
		}
	}
	return &Game{
		editor:  editor,
		preview: preview,
		grid:    grid,
		last:    time.Now(),
	}
}

func (g *Game) Update() error {
	g.last = time.Now()
	return g.editor.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.editor.Draw(screen)
	g.preview.Draw(screen)
	for _, sb := range g.grid {
		sb.Draw(screen)
	}
	elapsed := time.Now().Sub(g.last)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Time: %0.3fms", elapsed.Seconds()*1000), 300, 8)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pixie")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
