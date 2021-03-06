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

package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// KeyListener is a function to call in response to a KeyEvent
type KeyListener func(ke KeyEvent)

var keyListeners map[KeyEvent][]KeyListener

func init() {
	keyListeners = make(map[KeyEvent][]KeyListener)
}

// Register associates a KeyListener with a specific KeyEvent
func Register(event KeyEvent, listener KeyListener) {
	keyListeners[event] = append(keyListeners[event], listener)
}

// Update reads all of the currently pressed keys and compares them against previously pressed keys to generate KeyEvents
func Update() {
	curr := make(keyState)
	var mod Modifier
	// Get the status of the keys that currently pressed
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		mod |= Alt
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		mod |= Ctrl
	}
	for event := range keyListeners {
		if ebiten.IsKeyPressed(event.Key) {
			curr[event.Key] = true
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
