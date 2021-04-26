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

type keyState map[ebiten.Key]bool

var last = make(keyState)

func change(prev, curr keyState, mod Modifier, state State) (events []KeyEvent) {
	for key := range curr {
		if !last[key] {
			e := KeyEvent{
				Key:   key,
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
