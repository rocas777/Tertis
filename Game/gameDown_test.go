package Game

import (
	"testing"
	"tetris/Game/Piece"
	"tetris/Game/Position"
	"tetris/Game/Square"
)

func Test_Game_down_1(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{
		&MockPosition{0, game.Height() - 1, game.Width(), game.Height()},
		&MockPosition{1, game.Height() - 1, game.Width(), game.Height()},
		&MockPosition{2, game.Height() - 1, game.Width(), game.Height()},
		&MockPosition{3, game.Height() - 1, game.Width(), game.Height()},
	}
	game.AddPiece(MockColourSquare{}, positions1)
	game.Down()

	board := game.Board()
	positions1 = game.CurrentPiece().Positions()

	if positions1[0].Y() != game.Height()-1 {
		t.Errorf("")
	}
	if positions1[1].Y() != game.Height()-1 {
		t.Errorf("")
	}
	if positions1[2].Y() != game.Height()-1 {
		t.Errorf("")
	}
	if positions1[3].Y() != game.Height()-1 {
		t.Errorf("")
	}

	if !board[game.Height()-1][0].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-1][1].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-1][2].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-1][3].Occupied() {
		t.Errorf("")
	}
}

func Test_Game_down_2(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{
		&MockPosition{0, game.Height() - 2, game.Width(), game.Height()},
		&MockPosition{1, game.Height() - 3, game.Width(), game.Height()},
		&MockPosition{2, game.Height() - 4, game.Width(), game.Height()},
		&MockPosition{3, game.Height() - 5, game.Width(), game.Height()},
	}
	game.AddPiece(MockColourSquare{}, positions1)
	game.Down()

	board := game.Board()
	positions1 = game.CurrentPiece().Positions()

	if positions1[0].Y() != game.Height()-1 {
		t.Errorf("")
	}
	if positions1[1].Y() != game.Height()-2 {
		t.Errorf("")
	}
	if positions1[2].Y() != game.Height()-3 {
		t.Errorf("")
	}
	if positions1[3].Y() != game.Height()-4 {
		t.Errorf("")
	}

	if !board[game.Height()-1][0].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-2][1].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-3][2].Occupied() {
		t.Errorf("")
	}
	if !board[game.Height()-4][3].Occupied() {
		t.Errorf("")
	}
	if board[game.Height()-5][3].Occupied() {
		t.Errorf("")
	}
}

func Test_Game_down_3(t *testing.T) {
	localGame := NewGame(10, 10)
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewTPositions(10, 10))
	if !localGame.Down() {
		t.Errorf("Should be able to get one down")
	}
	localGame.Down()
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewTPositions(10, 10))
	if localGame.Down() {
		t.Errorf("Should not be able to get one down")
	}
}

func Test_Game_down_4(t *testing.T) {
	localGame := NewGame(10, 4)
	localGame.Next()
	localGame.Next()
	if localGame.Down() {
		t.Errorf("Should not be able to get one down")
	}
}
