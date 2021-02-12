package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type YellowSquare struct {
}

func (v YellowSquare) Occupied() bool {
	return true
}

func (v YellowSquare) Color() color.Color {
	return colornames.Yellow
}
