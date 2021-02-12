package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type RedSquare struct {
}

func (v RedSquare) Occupied() bool {
	return true
}

func (v RedSquare) Color() color.Color {
	return colornames.Red
}
