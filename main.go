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
