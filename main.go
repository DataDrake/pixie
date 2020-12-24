package main

import (
	"fmt"
	"github.com/DataDrake/pixie/encoding"
	"github.com/DataDrake/pixie/input"
	"github.com/DataDrake/pixie/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"time"
)

const (
	screenWidth  = 440
	screenHeight = 360
)

type Game struct {
	editor  *ui.Editor
	preview *ui.Preview
	toolbar *ui.Toolbar
	sprites *ui.Selector
	colors  *ui.Selector
	grid    []*ui.Box
	last    time.Time
}

func Quit(input.KeyEvent) {
	os.Exit(0)
}

var QuitEvent = input.KeyEvent{
	Key:   ebiten.KeyQ,
	Mod:   input.Ctrl,
	State: input.Pressed,
}

func NewGame() *Game {
	input.Register(QuitEvent, Quit)

	pl, err := encoding.LoadPalette("examples/palette.json")
	if err != nil {
		panic(err)
	}
	pl.Describe()
	println()
	palette := color.Palette(pl.Colors)

	ss, err := encoding.LoadSpriteSet("examples/sprite_set.json")
	if err != nil {
		panic(err)
	}
	ss.Describe()
	ss.ChangePalette(&palette)

	tbss, err := encoding.LoadSpriteSet("examples/sprite_toolbar.json")
	if err != nil {
		panic(err)
	}
	tbss.Describe()
	tbss.ChangePalette(&palette)

	s := ui.NewSprite(16, 16, &palette)
	s.Swap(ss.Sprites[0].Convert())

	return &Game{
		editor:  ui.NewEditor(88, 8, s),
		preview: ui.NewPreview(30, 226, s, &palette),
		toolbar: ui.NewToolbar(8, 8, &tbss, &palette),
		sprites: ui.NewSelector(8, 58, s, &palette),
		colors:  ui.NewSelector(348, 58, s, &palette),
		last:    time.Now(),
	}
}

func (g *Game) Update() error {
	g.last = time.Now()
	input.Update()
	return g.editor.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.editor.Draw(screen)
	g.preview.Draw(screen)
	g.toolbar.Draw(screen)
	g.sprites.Draw(screen)
	g.colors.Draw(screen)
	elapsed := time.Now().Sub(g.last)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Time: %0.3fms", elapsed.Seconds()*1000), 8, 300)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pixie")
	println(ebiten.KeyMax)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
