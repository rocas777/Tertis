package Game

import (
	"testing"
	"tetris/Game/Piece"
	"tetris/Game/Position"
	"tetris/Game/Square"
)

func Test_Game_left_1(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{&MockPosition{1, 1, game.Width(), game.Height()}, &MockPosition{1, 2, game.Width(), game.Height()}, &MockPosition{1, 3, game.Width(), game.Height()}, &MockPosition{1, 4, game.Width(), game.Height()}}
	game.AddPiece(MockVoidSquare{}, positions1)

	game.Left()
	positions2 := game.CurrentPiece().Positions()

	x1 := positions2[0].X()
	x2 := positions2[1].X()
	x3 := positions2[2].X()
	x4 := positions2[3].X()
	y1 := positions2[0].Y()
	y2 := positions2[1].Y()
	y3 := positions2[2].Y()
	y4 := positions2[3].Y()

	if x1 != 0 {
		t.Errorf("x_1 = %d; want 0", x1)
	}
	if x2 != 0 {
		t.Errorf("x_2 = %d; want 0", x2)
	}
	if x3 != 0 {
		t.Errorf("x_3 = %d; want 0", x3)
	}
	if x4 != 0 {
		t.Errorf("x_4 = %d; want 0", x4)
	}

	if y1 != 1 {
		t.Errorf("y_1 = %d; want 0", y1)
	}
	if y2 != 2 {
		t.Errorf("y_2 = %d; want 0", y2)
	}
	if y3 != 3 {
		t.Errorf("y_3 = %d; want 0", y3)
	}
	if y4 != 4 {
		t.Errorf("y_4 = %d; want 0", y4)
	}
}

func Test_Game_left_2(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{
		&MockPosition{1, 0, game.Width(), game.Height()},
		&MockPosition{2, 0, game.Width(), game.Height()},
		&MockPosition{2, 1, game.Width(), game.Height()},
		&MockPosition{3, 1, game.Width(), game.Height()},
	}
	game.AddPiece(MockVoidSquare{}, positions1)

	game.Left()
	positions2 := game.CurrentPiece().Positions()

	x1 := positions2[0].X()
	x2 := positions2[1].X()
	x3 := positions2[2].X()
	x4 := positions2[3].X()
	y1 := positions2[0].Y()
	y2 := positions2[1].Y()
	y3 := positions2[2].Y()
	y4 := positions2[3].Y()

	if x1 != 0 {
		t.Errorf("x_1 = %d; want 0", x1)
	}
	if x2 != 1 {
		t.Errorf("x_2 = %d; want 0", x2)
	}
	if x3 != 1 {
		t.Errorf("x_3 = %d; want 0", x3)
	}
	if x4 != 2 {
		t.Errorf("x_4 = %d; want 0", x4)
	}

	if y1 != 0 {
		t.Errorf("y_1 = %d; want 0", y1)
	}
	if y2 != 0 {
		t.Errorf("y_2 = %d; want 0", y2)
	}
	if y3 != 1 {
		t.Errorf("y_3 = %d; want 0", y3)
	}
	if y4 != 1 {
		t.Errorf("y_4 = %d; want 0", y4)
	}
}

func Test_Game_left_3(t *testing.T) {
	game.Clear()
	positions1 := [4]Position.IPosition{
		&MockPosition{1, 0, game.Width(), game.Height()},
		&MockPosition{2, 0, game.Width(), game.Height()},
		&MockPosition{2, 1, game.Width(), game.Height()},
		&MockPosition{3, 1, game.Width(), game.Height()},
	}
	game.AddPiece(MockColourSquare{}, positions1)

	board1 := game.Board()

	if board1[0][1].Occupied() != true {
		t.Errorf("Spot (0,0) must be occupied")
	}
	if board1[0][2].Occupied() != true {
		t.Errorf("Spot (1,0) must be occupied")
	}
	if board1[1][2].Occupied() != true {
		t.Errorf("Spot (1,1) must be occupied")
	}
	if board1[1][3].Occupied() != true {
		t.Errorf("Spot (2,1) must be occupied")
	}

	game.Left()
	board2 := game.Board()

	if board2[0][0].Occupied() != true {
		t.Errorf("Spot (0,0) must be occupied")
	}
	if board2[0][1].Occupied() != true {
		t.Errorf("Spot (1,0) must be occupied")
	}
	if board2[1][1].Occupied() != true {
		t.Errorf("Spot (1,1) must be occupied")
	}
	if board2[1][2].Occupied() != true {
		t.Errorf("Spot (2,1) must be occupied")
	}

	a := countNumberOfOccupied(t, board2)
	if a != 4 {
		t.Errorf("Only 4 occupied pieces must be used")
	}
}

func Test_Game_left_4(t *testing.T) {
	localGame := NewGame(10, 10)
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewOPositions(10, 10))
	r11 := localGame.Left()
	r12 := localGame.Left()
	r13 := localGame.Left()
	r14 := localGame.Left()
	if !r11 || !r12 || !r13 || !r14 {
		t.Errorf("Should be able to move")
	}
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewTPositions(10, 10))

	if !localGame.Left() {
		t.Errorf("Should be able to get one left")
	}
	localGame.Left()
	r2 := localGame.Left()
	if r2 {
		t.Errorf("Should not be able to move")
	}
}

func Test_Game_left_5(t *testing.T) {
	localGame := NewGame(10, 4)
	localGame.AddPiece(Square.YellowSquare{}, Piece.NewTPositions(10, 4))
	r11 := localGame.Left()
	r12 := localGame.Left()
	r13 := localGame.Left()
	if !r11 || !r12 || !r13 {
		t.Errorf("Should be able to move")
	}
	r2 := localGame.Left()
	if r2 {
		t.Errorf("Should not be able to move")
	}
}
