package Position

type IPosition interface {
	X() uint8
	Y() uint8
	Left() IPosition
	Right() IPosition
	Down() IPosition
}

func Equal(p1 IPosition, p2 IPosition) bool {
	return p1.X() == p2.X() && p1.Y() == p2.Y()
}
