package main

import (
	"image/color"
	"image/gif"
	"log"
	"os"
	"strconv"
)

var (
	ParrotColors     []color.Color
	DarkParrotColors []color.Color
	LightGopherBlue  color.Color
	DarkGopherBlue   color.Color
)

func init() {
	var err error

for _, s := range []string{
		"FF6B6B",
		"FF6BB5",
		"FF81FF",
		"FF81FF",
		"D081FF",
		"81ACFF",
		"81FFFF",
		"81FF81",
		"FFD081",
		"FF8181",
} {
		c, err := hexToColor(s)
		if err != nil {
			log.Fatal(err)
		}
		ParrotColors = append(ParrotColors, c)
		DarkParrotColors = append(DarkParrotColors, darken(c))
	}

	LightGopherBlue, err = hexToColor("8BD0FF")
	if err != nil {
		log.Fatal(err)
	}
	DarkGopherBlue, err = hexToColor("82C2EE")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var (
		lbi int
		dbi int
	)


	f, err := os.Open("dancing-gopher.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gopher, err := gif.DecodeAll(f)
	if err != nil {
		log.Fatal(err)
	}
	
	for i, frame := range gopher.Image {
		lbi = frame.Palette.Index(LightGopherBlue)
		dbi = frame.Palette.Index(DarkGopherBlue)

		frame.Palette[lbi] = ParrotColors[i%len(ParrotColors)]
		frame.Palette[dbi] = DarkParrotColors[i%len(DarkParrotColors)]
	}

	o, _ := os.OpenFile("party-gopher.gif", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer o.Close()
	gif.EncodeAll(o, gopher)
}

func hexToColor(hex string) (color.Color, error) {
	c := color.RGBA{0, 0, 0, 255}

	r, err := strconv.ParseInt(hex[0:2], 16, 16)
	if err != nil {
		return c, err
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 16)
	if err != nil {
		return c, err
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 16)
	if err != nil {
		return c, err
	}
