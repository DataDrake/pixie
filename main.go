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
	"github.com/DataDrake/pixie/editor"
	"github.com/DataDrake/pixie/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"os"
	"time"
)

const (
	screenWidth  = 448
	screenHeight = 360
)

// Pixie is the global Pixie object
type Pixie struct {
	editors  [editor.Max]editor.Editor
	current  editor.Kind
	last     time.Time
	quitting bool
}

// QuitEvent listens for CTRL+Q to be pressed
var QuitEvent = input.KeyEvent{
	Key:   ebiten.KeyQ,
	Mod:   input.Ctrl,
	State: input.Pressed,
}

// SaveEvent listens for CTRL+S to be pressed
var SaveEvent = input.KeyEvent{
	Key:   ebiten.KeyS,
	Mod:   input.Ctrl,
	State: input.Pressed,
}

// NewPixie creates a new Pixie object and populates the GUI
func NewPixie() *Pixie {
	// Build Editors
	var editors [editor.Max]editor.Editor
	editors[editor.SpriteKind] = editor.NewSprite()
	// Build Pixie
	p := &Pixie{
		editors: editors,
		current: editor.SpriteKind,
		last:    time.Now(),
	}
	// Register the quit event
	input.Register(QuitEvent, func(_ input.KeyEvent) {
		p.quitting = true
		p.current = editor.SpriteKind
	})
	// Register the save event
	input.Register(SaveEvent, func(_ input.KeyEvent) {
		if err := p.Save(); err != nil {
			log.Println(err)
		}
	})
	return p
}

// Save tells the current editor to save the thing(s) it is editing
func (p *Pixie) Save() error {
	return p.editors[p.current].Save()
}

// Update checks for all updates in the input and the internal state of Pixie
func (p *Pixie) Update() error {
	p.last = time.Now()
	input.Update()
	if p.quitting {
		done, err := p.editors[p.current].Exit()
		if err != nil {
			return err
		}
		if done {
			p.current++
		}
		if p.current == editor.Max {
			os.Exit(0)
		}
		return nil
	}
	return p.editors[p.current].Update()
}

// Draw renders Pixie to a screen
func (p *Pixie) Draw(screen *ebiten.Image) {
	p.editors[p.current].Draw(screen)
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
	if err := ebiten.RunGame(NewPixie()); err != nil {
		log.Fatal(err)
	}
}
