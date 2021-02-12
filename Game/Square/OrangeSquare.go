package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type OrangeSquare struct {
}

func (v OrangeSquare) Occupied() bool {
	return true
}

func (v OrangeSquare) Color() color.Color {
	return colornames.Orange
}
