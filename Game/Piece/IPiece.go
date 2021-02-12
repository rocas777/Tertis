package Piece

import "tetris/Game/Position"

type IPiece interface {
	Positions() [4]Position.IPosition
	SetPositions([4]Position.IPosition)
	IsNull() bool
}
