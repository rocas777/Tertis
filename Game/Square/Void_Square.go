package Square

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type VoidSquare struct {
}

func (v VoidSquare) Occupied() bool {
	return false
}

func (v VoidSquare) Color() color.Color {
	return colornames.Grey
}
