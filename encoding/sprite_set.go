package encoding

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"text/tabwriter"
	"time"
)

type SpriteSet struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Sprites  []*Sprite `json:"sprites"`
	palette  *color.Palette
}

func LoadSpriteSet(path string) (ss SpriteSet, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&ss)
	return
}

func (ss *SpriteSet) ChangePalette(p *color.Palette) {
	ss.palette = p
	for _, s := range ss.Sprites {
		s.img.Palette = *p
	}
}

func (ss *SpriteSet) Palette() *color.Palette {
	return ss.palette
}

func (ss *SpriteSet) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", ss.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", ss.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", ss.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", ss.Revision)
	fmt.Fprintf(tw, "Number of Sprites\t: %d\n", len(ss.Sprites))
	tw.Flush()
}
