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

package encoding

import (
	"encoding/hex"
	"encoding/json"
	"image/color"
)

type colorJSON []byte

func newColor(c color.Color) colorJSON {
	var cj colorJSON
	r, g, b, a := c.RGBA()
	cj = append(cj, byte(r), byte(g), byte(b), byte(a))
	return cj
}

func (cj colorJSON) Color() color.Color {
	return color.RGBA{
		R: uint8(cj[0]),
		G: uint8(cj[1]),
		B: uint8(cj[2]),
		A: uint8(cj[3]),
	}
}

func (cj colorJSON) MarshalJSON() ([]byte, error) {
	s := hex.EncodeToString([]byte(cj))
	return json.Marshal(s)
}

func (cj *colorJSON) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	raw, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*cj = colorJSON(raw)
	return nil
}
