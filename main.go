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

package main

import (
	"fmt"
	"github.com/DataDrake/pixie/assets"
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

// Pixie is the global Pixie object
type Pixie struct {
	editor   *ui.Editor
	preview  *ui.Preview
	toolbar  *ui.Toolbar
	sprites  *ui.Selector
	colors   *ui.Palette
	cPreview *ui.ColorPreview
	grid     []*ui.Box
	last     time.Time
}

// Quit is called to terminate the program
func Quit(input.KeyEvent) {
	os.Exit(0)
}

// QuitEvent listens for CTRL+Q to be pressed
var QuitEvent = input.KeyEvent{
	Key:   ebiten.KeyQ,
	Mod:   input.Ctrl,
	State: input.Pressed,
}

// NewPixie creates a new Pixie object and populates the GUI
func NewPixie() *Pixie {
	input.Register(QuitEvent, Quit)

	pl, err := encoding.LoadPalette(assets.DefaultPalette())
	if err != nil {
		panic(err)
	}
	pl.Describe()
	println()
	palette := color.Palette(pl.Colors)

	ss, err := encoding.LoadSpriteSet(assets.DefaultSprites())
	if err != nil {
		panic(err)
	}
	ss.Describe()
	ss.ChangePalette(&palette)

	tbss, err := encoding.LoadSpriteSet(assets.DefaultToolbarIcons())
	if err != nil {
		panic(err)
	}
	tbss.Describe()
	tbss.ChangePalette(&palette)

	s := ui.NewSprite(16, 16, &palette)
	s.Swap(ss.Sprites[0].Convert())

	return &Pixie{
		editor:   ui.NewEditor(88, 8, s),
		preview:  ui.NewPreview(30, 226, s, &palette),
		toolbar:  ui.NewToolbar(8, 8, &tbss, &palette),
		sprites:  ui.NewSelector(8, 58, s, &palette),
		colors:   ui.NewPalette(348, 58, &palette),
		cPreview: ui.NewColorPreview(371, 222, palette[1], palette[0]),
		last:     time.Now(),
	}
}

// Update checks for all updates in the input and the internal state of Pixie
func (p *Pixie) Update() error {
	p.last = time.Now()
	input.Update()
	return p.editor.Update()
}

// Draw renders Pixie to a screen
func (p *Pixie) Draw(screen *ebiten.Image) {
	p.editor.Draw(screen)
	p.preview.Draw(screen)
	p.toolbar.Draw(screen)
	p.sprites.Draw(screen)
	p.colors.Draw(screen)
	p.cPreview.Draw(screen)
	elapsed := time.Now().Sub(p.last)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Time: %0.3fms", elapsed.Seconds()*1000), 8, 300)
}

// Layout specifies the layout dimensions of Pixie
func (p *Pixie) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pixie")
	println(ebiten.KeyMax)
	if err := ebiten.RunGame(NewPixie()); err != nil {
		log.Fatal(err)
	}
}
