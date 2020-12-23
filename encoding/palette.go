package encoding

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Palette struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Colors   Colors    `json:"colors"`
}

func LoadPalette(path string) (p Palette, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&p)
	return
}

func (p *Palette) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", p.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", p.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", p.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", p.Revision)
	fmt.Fprintf(tw, "Number of Colors\t: %d\n", len(p.Colors))
	tw.Flush()
}
