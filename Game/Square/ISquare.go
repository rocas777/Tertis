package Square

import "image/color"

type ISquare interface {
	Occupied() bool
	Color() color.Color
}
