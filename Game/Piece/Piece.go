package Piece

import (
	"math"
	"tetris/Game/Position"
)

type Piece struct {
	positions [4]Position.IPosition
}

func NewPiece(positions [4]Position.IPosition) IPiece {
	return &Piece{positions}
}

func (p *Piece) IsNull() bool {
	return false
}

func (p *Piece) SetPositions(positions [4]Position.IPosition) {
	p.positions = positions
}

func (p Piece) Positions() [4]Position.IPosition {
	return p.positions
}

func NewTPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-3) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(0+addition, 1, w, h),
		Position.NewPosition(1+addition, 1, w, h),
		Position.NewPosition(2+addition, 1, w, h),
	}
}
func NewOPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-2) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(0+addition, 0, w, h),
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(0+addition, 1, w, h),
		Position.NewPosition(1+addition, 1, w, h),
	}
}

func NewJPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-3) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(0+addition, 0, w, h),
		Position.NewPosition(0+addition, 1, w, h),
		Position.NewPosition(1+addition, 1, w, h),
		Position.NewPosition(2+addition, 1, w, h),
	}
}

func NewLPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-3) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(0+addition, 0, w, h),
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(2+addition, 0, w, h),
		Position.NewPosition(0+addition, 1, w, h),
	}
}
func NewSPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-3) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(2+addition, 0, w, h),
		Position.NewPosition(0+addition, 1, w, h),
		Position.NewPosition(1+addition, 1, w, h),
	}
}
func NewZPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-3) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(0+addition, 0, w, h),
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(1+addition, 1, w, h),
		Position.NewPosition(2+addition, 1, w, h),
	}
}
func NewIPositions(w, h uint8) [4]Position.IPosition {
	addition := uint8(math.Floor(float64(w-4) / 2))
	return [4]Position.IPosition{
		Position.NewPosition(0+addition, 0, w, h),
		Position.NewPosition(1+addition, 0, w, h),
		Position.NewPosition(2+addition, 0, w, h),
		Position.NewPosition(3+addition, 0, w, h),
	}
}
