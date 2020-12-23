package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyListener func(ke KeyEvent)

var keyListeners map[KeyEvent][]KeyListener
var keys = []ebiten.Key{ebiten.KeyAlt, ebiten.KeyControl}

func init() {
	keyListeners = make(map[KeyEvent][]KeyListener)
}

func Register(event KeyEvent, listener KeyListener) {
	keyListeners[event] = append(keyListeners[event], listener)
	keys = append(keys, event.Key)
}

func Update() {
	var curr keyState
	var mod Modifier
	// Get the status of the keys that currently pressed
	for _, k := range keys {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyAlt:
				mod |= Alt
			case ebiten.KeyControl:
				mod |= Ctrl
			default:
				curr = append(curr, k)
			}
		}
	}
	// Generate KeyEvents for changes in state
	for _, event := range genEvents(last, curr, mod) {
		for _, listener := range keyListeners[event] {
			listener(event)
		}
	}
	// Update state for next frame
	last = curr
}
