package encoding

import (
    "encoding/json"
    "image/color"
)

type Colors color.Palette

type colorsJSON []colorJSON

func (cj colorsJSON) Palette() Colors {
    var cs Colors
    for _, c := range cj {
        cs = append(cs, c.Color())
    }
    return cs
}

func (cs Colors) MarshalJSON() ([]byte, error) {
    var j colorsJSON
    for _, c := range cs {
        j = append(j, newColor(c))
    }
    return json.Marshal(j)
}

func (cs *Colors) UnmarshalJSON(bs []byte) error {
    var j colorsJSON
    if err := json.Unmarshal(bs, &j); err != nil {
        return err
    }
    *cs = j.Palette()
    return nil
}
