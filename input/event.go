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

// KeyEvent represents the combinator of a normal key, modifier keys, and a State, used to represent an Event to scan for
type KeyEvent struct {
	Key   ebiten.Key
	Mod   Modifier
	State State
}

// Modifier represents zero or more modifier keys
type Modifier int

const (
	// None indicates no modifiers
	None Modifier = 0
	// Alt indicates Left Alt is pressed
	Alt = 1
	// Ctrl indicated that either Ctrl key is pressed
	Ctrl = 2
)

// State indicates the state of a Key
type State int

const (
	// Pressed indicates a Key was just pressed
	Pressed State = iota
	// Released indicated a Key was just released
	Released
)
