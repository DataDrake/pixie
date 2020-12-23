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
    return color.RGBA {
        R: uint8(cj[0]),
        G: uint8(cj[1]),
        B: uint8(cj[2]),
        A: uint8(cj[3]),
    }
}

func (c colorJSON) MarshalJSON() ([]byte, error) {
    s := hex.EncodeToString([]byte(c))
    return json.Marshal(s)
}

func (c *colorJSON) UnmarshalJSON(bs []byte) error {
    var s string
    if err := json.Unmarshal(bs, &s); err != nil {
        return err
    }
    raw, err := hex.DecodeString(s)
    if err != nil {
        return err
    }
    *c = colorJSON(raw)
    return nil
}
