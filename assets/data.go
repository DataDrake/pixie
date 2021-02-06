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

package assets

import (
	"path/filepath"
)

// DataDir is the shared data directory for Pixie's assets
var DataDir string

// DefaultPalette is the location of the default palette to load when opening Pixie
func DefaultPalette() string {
	return filepath.Join(DataDir, "defaults", "palette.json")
}

// DefaultSprites are the "starter" sprites to load when opening Pixie
func DefaultSprites() string {
	return filepath.Join(DataDir, "defaults", "sprites.json")
}

// DefaultToolbarIcons are the icons to use for the toolbar in the sprite editor
func DefaultToolbarIcons() string {
	return filepath.Join(DataDir, "defaults", "sprite_toolbar.json")
}
