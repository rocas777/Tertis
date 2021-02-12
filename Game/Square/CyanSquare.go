package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type CyanSquare struct {
}

func (v CyanSquare) Occupied() bool {
	return true
}

func (v CyanSquare) Color() color.Color {
	return colornames.Cyan
}
