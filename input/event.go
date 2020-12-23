package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyEvent struct {
	Key   ebiten.Key
	Mod   Modifier
	State State
}

type Modifier int

const (
	None Modifier = 0
	Alt           = 1
	Ctrl          = 2
)

type State int

const (
	Pressed State = iota
	Released
)
