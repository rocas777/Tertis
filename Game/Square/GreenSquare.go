package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type GreenSquare struct {
}

func (v GreenSquare) Occupied() bool {
	return true
}

func (v GreenSquare) Color() color.Color {
	return colornames.Green
}
