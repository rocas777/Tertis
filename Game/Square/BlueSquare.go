package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type BlueSquare struct {
}

func (v BlueSquare) Occupied() bool {
	return true
}

func (v BlueSquare) Color() color.Color {
	return colornames.Blue
}
