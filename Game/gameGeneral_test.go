package Game

import (
	"golang.org/x/image/colornames"
	"image/color"
	"testing"
	"tetris/Game/Piece"
	"tetris/Game/Position"
	"tetris/Game/Square"
)

type MockVoidSquare struct{}

func (m MockVoidSquare) Occupied() bool     { return false }
func (m MockVoidSquare) Color() color.Color { return colornames.Grey }

type MockPosition struct {
	x         uint8
	y         uint8
	maxWidth  uint8
	maxHeight uint8
}

func (t MockPosition) Right() Position.IPosition {
	if t.x < t.maxWidth-1 {
		return MockPosition{t.x + 1, t.y, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t MockPosition) Left() Position.IPosition {
	if t.x > 0 {
		return MockPosition{t.x - 1, t.y, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t MockPosition) Down() Position.IPosition {
	if t.y < t.maxHeight-1 {
		return MockPosition{t.x, t.y + 1, t.maxWidth, t.maxHeight}
	}
	return t
}
func (t MockPosition) X() uint8 { return t.x }
func (t MockPosition) Y() uint8 { return t.y }

type MockColourSquare struct{}

func (m MockColourSquare) Occupied() bool     { return true }
func (m MockColourSquare) Color() color.Color { return colornames.Blue }

var game = NewGame(10, 20)

func countNumberOfOccupied(t *testing.T, board [][]Square.ISquare) uint {
	var out uint = 0
	for _, line := range board {
		for _, square := range line {
			if square.Occupied() != false {
				out++
			} else if square.Color() != colornames.Grey {
				t.Errorf("Color must be none")
			}
		}
	}
	return out
}

func Test_Game_Initial(t *testing.T) {
	game.Clear()
	board := game.Board()
	if countNumberOfOccupied(t, board) != 0 {
		t.Errorf("Every board spot must be clear")
	}
}

func Test_Clear(t *testing.T) {
	localGame := NewGame(10, 10)
	positions1 := [4]Position.IPosition{&MockPosition{localGame.Width() - 1, 1, localGame.Width(), localGame.Height()}, &MockPosition{localGame.Width() - 1, 2, localGame.Width(), localGame.Height()}, &MockPosition{localGame.Width() - 1, 3, localGame.Width(), localGame.Height()}, &MockPosition{localGame.Width() - 1, 4, localGame.Width(), localGame.Height()}}
	localGame.AddPiece(MockColourSquare{}, positions1)

	if countNumberOfOccupied(t, localGame.Board()) == 0 {
		t.Errorf("Board should not be clear")
	}
	localGame.Clear()
	if countNumberOfOccupied(t, localGame.Board()) != 0 {
		t.Errorf("Every board spot must be clear")
	}

}

func Test_Check_Piece_Addition(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{&MockPosition{game.Width() - 1, 1, game.Width(), game.Height()}, &MockPosition{game.Width() - 1, 2, game.Width(), game.Height()}, &MockPosition{game.Width() - 1, 3, game.Width(), game.Height()}, &MockPosition{game.Width() - 1, 4, game.Width(), game.Height()}}
	game.AddPiece(MockVoidSquare{}, positions1)
	if game.CurrentPiece().Positions() != positions1 {
		t.Errorf("Vectors should be the same")
	}
}

func Test_Next_1(t *testing.T) {
	localGame := NewGame(10, 10)
	localGame.Next()
	lastPiece := localGame.CurrentPiece()
	if countNumberOfOccupied(t, localGame.Board()) != 4 {
		t.Errorf("Should have 4 squares")
	}
	localGame.Next()
	currentPiece := localGame.CurrentPiece()
	if lastPiece != currentPiece {
		t.Errorf("Pieces should be the same")
	}
}

func Test_Next_2(t *testing.T) {
	game.Clear()
	localBoard := NewGame(10, 10).Board()
	localBoard[4][5] = Square.YellowSquare{}
	localBoard[4][4] = Square.YellowSquare{}
	localBoard[4][3] = Square.YellowSquare{}
	localBoard[3][6] = Square.YellowSquare{}
	localGame := Model{
		board: localBoard,
		piece: Piece.NullPiece{},
	}
	if !localGame.Next() {
		t.Errorf("Next should be true")
	}

	if !localGame.Next() {
		t.Errorf("Next should be true")
	}

	if localGame.Next() {
		t.Errorf("Next should be false")
	}

}

func Test_InstantDrop(t *testing.T) {
	game.Clear()
	localGame := NewGame(10, 10)
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewOPositions(10, 10))
	localGame.Drop()
	if !localGame.Board()[8][4].Occupied() {
		t.Errorf("Should be occupied")
	}
	if !localGame.Board()[8][5].Occupied() {
		t.Errorf("Should be occupied")
	}
	if !localGame.Board()[9][4].Occupied() {
		t.Errorf("Should be occupied")
	}
	if !localGame.Board()[9][5].Occupied() {
		t.Errorf("Should be occupied")
	}
	if countNumberOfOccupied(t, localGame.Board()) != 4 {
		t.Errorf("Should only exist 4 Squares occupied")
	}
}
