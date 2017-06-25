package main

import (
	"github.com/nsf/termbox-go"
)

var colorMap = map[int]termbox.Attribute{
	2:    termbox.ColorWhite,
	4:    termbox.ColorWhite,
	8:    termbox.ColorWhite,
	16:   termbox.ColorWhite,
	32:   termbox.ColorWhite,
	64:   termbox.ColorGreen,
	128:  termbox.ColorMagenta,
	256:  termbox.ColorMagenta,
	512:  termbox.ColorMagenta,
	1024: termbox.ColorCyan,
	2048: termbox.ColorRed,
}

func getColor(value int) termbox.Attribute {
	if color, ok := colorMap[value]; ok {
		return color
	}

	return termbox.ColorRed
}
