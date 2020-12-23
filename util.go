package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func getPressed() (keys []ebiten.Key) {
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			keys = append(keys, k)
		}
	}
	return
}

func In(r image.Rectangle, x, y int) bool {
	return image.Rect(x, y, x+1, y+1).In(r)
}
