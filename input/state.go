package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type keyState []ebiten.Key

var last keyState

func change(prev, curr keyState, mod Modifier, state State) (events []KeyEvent) {
	for _, cKey := range curr {
		found := false
		for _, pKey := range prev {
			if cKey == pKey {
				found = true
				break
			}
		}
		if !found {
			e := KeyEvent{
				Key:   cKey,
				Mod:   mod,
				State: state,
			}
			events = append(events, e)
		}
	}
	return
}

func genEvents(prev, curr keyState, mod Modifier) (events []KeyEvent) {
	// pressed keys
	events = change(prev, curr, mod, Pressed)
	// released keys
	events = append(events, change(curr, prev, mod, Released)...)
	return
}
