package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	rkeyabord "github.com/hajimehoshi/ebiten/v2/examples/resources/images/keyboard"
)

type Keyboard struct {
	img     *ebiten.Image
	pressed []ebiten.Key
}

func NewKeyboard() *Keyboard {
	img, _, err := image.Decode(bytes.NewReader(rkeyabord.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}
	return &Keyboard{
		img: ebiten.NewImageFromImage(img),
	}
}

func (k *Keyboard) Update() error {
	k.pressed = getPressed()
	return nil
}

func (k *Keyboard) Draw(screen *ebiten.Image) {
	const (
		offsetX = 24
		offsetY = 40
	)

	// Draw the base (grayed) keyboard image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(offsetX, offsetY)
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	screen.DrawImage(k.img, op)

	// Draw the highlighted keys.
	op = &ebiten.DrawImageOptions{}
	for _, p := range k.pressed {
		op.GeoM.Reset()
		r, ok := keyboard.KeyRect(p)
		if !ok {
			continue
		}
		op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
		op.GeoM.Translate(offsetX, offsetY)
		screen.DrawImage(k.img.SubImage(r).(*ebiten.Image), op)
	}

	keyStrs := []string{}
	for _, p := range k.pressed {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
}
