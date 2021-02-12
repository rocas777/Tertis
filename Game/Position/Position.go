package Position

type Position struct {
	x         uint8
	y         uint8
	maxWidth  uint8
	maxHeight uint8
}

func NewPosition(x, y, maxWidth, maxHeight uint8) IPosition {
	return Position{x, y, maxWidth, maxHeight}
}

func (t Position) Right() IPosition {
	if t.x < t.maxWidth-1 {
		return Position{t.x + 1, t.y, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t Position) Left() IPosition {
	if t.x > 0 {
		return Position{t.x - 1, t.y, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t Position) Down() IPosition {
	if t.y < t.maxHeight-1 {
		return Position{t.x, t.y + 1, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t Position) X() uint8 { return t.x }
func (t Position) Y() uint8 { return t.y }
