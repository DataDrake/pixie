package encoding

import (
    "encoding/hex"
    "encoding/json"
    "github.com/hajimehoshi/ebiten/v2"
    "image"
    "image/color"
    "os"
)

type Sprite struct {
    img image.Paletted
}

type spriteJSON struct {
    Size    int      `json:"size"`
    Pixels  []string `json:"pixels"`
}

func LoadSprite(path string) (s Sprite, err error) {
    f, err := os.Open(path)
    if err != nil {
        return
    }
    defer f.Close()
    dec := json.NewDecoder(f)
    err = dec.Decode(&s)
    s.img.Palette = []color.Color {
        color.RGBA{0x00,0x00,0x00,0xFF},
        color.RGBA{0xFF,0x00,0x00,0xFF},
    }
    return
}

func (s *Sprite) Convert() *ebiten.Image {
    return ebiten.NewImageFromImage(&s.img)
}

func (s Sprite) MarshalJSON() (bs []byte, err error) {
    j := spriteJSON {
        Size: s.img.Stride,
    }
    var row []byte
    for i := 0; i < len(s.img.Pix); i += s.img.Stride {
        row = []byte(s.img.Pix[i*s.img.Stride:(i+1)*s.img.Stride])
        j.Pixels = append(j.Pixels, hex.EncodeToString(row))
    }
    return json.Marshal(j)
}

func (s *Sprite) UnmarshalJSON(b []byte) (err error) {
    var j spriteJSON
    if err = json.Unmarshal(b, &j); err != nil {
        return
    }
    var pix []byte
    for _, row := range j.Pixels {
        if pix, err = hex.DecodeString(row); err != nil {
            return
        }
        s.img.Pix = append(s.img.Pix, []uint8(pix)...)
    }
    s.img.Stride = j.Size
    s.img.Rect   = image.Rect(0, 0, j.Size, j.Size)
    return
}
