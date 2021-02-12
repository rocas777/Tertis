package Game

import (
	"tetris/Game/Piece"
	"tetris/Game/Position"
	"tetris/Game/Square"
)

type IGame interface {
	Board() [][]Square.ISquare
	AddPiece(square Square.ISquare, positions [4]Position.IPosition)
	CurrentPiece() Piece.IPiece
	Left() bool
	Right() bool
	Down() bool
	Drop()
	Width() uint8
	Height() uint8
	Clear()
	HasFullLine() bool
	SplashLines() uint8
	Collapse()
	Next() bool
}
