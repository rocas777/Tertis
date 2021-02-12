package Piece

import "tetris/Game/Position"

type NullPiece struct {
}

func (n NullPiece) IsNull() bool {
	return true
}

func (n NullPiece) Positions() [4]Position.IPosition {
	panic("I Should have done something...")
}

func (n NullPiece) SetPositions([4]Position.IPosition) {
	panic("I Should have done something...")
}
