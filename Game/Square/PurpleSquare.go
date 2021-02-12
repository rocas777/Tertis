package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type PurpleSquare struct {
}

func (v PurpleSquare) Occupied() bool {
	return true
}

func (v PurpleSquare) Color() color.Color {
	return colornames.Purple
}
