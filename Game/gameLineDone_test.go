package Game

import (
	"testing"
	"tetris/Game/Position"
)

type MockPiece struct {
}

func (m MockPiece) IsNull() bool {
	return false
}

func (m MockPiece) Positions() [4]Position.IPosition {
	panic("implement me")
}

func (m MockPiece) SetPositions([4]Position.IPosition) {
	panic("implement me")
}

func Test_Game_SplashLine_1(t *testing.T) {
	game.Clear()
	board := game.Board()
	if game.HasFullLine() {
		t.Errorf("Board should be cleared")
	}
	for x := 0; x < len(board[2]); x++ {
		board[2][x] = MockColourSquare{}
		board[3][x] = MockColourSquare{}
		board[4][x] = MockColourSquare{}
	}
	localGame := Model{board, MockPiece{}}
	if number := countNumberOfOccupied(t, board); number != uint(len(board[0])*3) {
		t.Errorf("Board should have %d coloured squares", len(board[0])*3)
	}
	if localGame.HasFullLine() {
		numberOfLines := localGame.SplashLines()
		if numberOfLines != 3 {
			t.Errorf("The function should have splashed 3 lines")
		}
	}
	if number := countNumberOfOccupied(t, board); number != 0 {
		t.Errorf("Board should have 0 coloured squares")
	}
}

func Test_Game_SplashLine_2(t *testing.T) {
	game.Clear()
	board := game.Board()
	for x := 0; x < len(board[2]); x++ {
		board[2][x] = MockColourSquare{}
		board[3][x] = MockColourSquare{}
		board[4][x] = MockColourSquare{}
	}
	board[4][0] = MockVoidSquare{}
	localGame := Model{board, MockPiece{}}
	if number := countNumberOfOccupied(t, board); number != uint(len(board[0])*3-1) {
		t.Errorf("Board should have %d coloured squares", len(board[0])*3-1)
	}
	if localGame.HasFullLine() {
		numberOfLines := localGame.SplashLines()
		if numberOfLines != 2 {
			t.Errorf("The function should have splashed 2 lines")
		}
	}
	if number := countNumberOfOccupied(t, board); number != uint(len(board[0])-1) {
		t.Errorf("Board should have %d coloured squares", len(board[0])-1)
	}

	numberOfLines := localGame.SplashLines()
	if numberOfLines != 0 {
		t.Errorf("The function should have splashed 0 lines")
	}
	if number := countNumberOfOccupied(t, board); number != uint(len(board[0])-1) {
		t.Errorf("Board should have %d coloured squares", len(board[0])-1)
	}
}

func Test_Game_Collapse_2(t *testing.T) {
	game.Clear()
	board := game.Board()
	for x := 0; x < len(board[2]); x++ {
		board[game.Height()-1][x] = MockColourSquare{}
		board[game.Height()-2][x] = MockColourSquare{}
		board[game.Height()-4][x] = MockColourSquare{}
		board[game.Height()-6][x] = MockColourSquare{}
	}

	board[game.Height()-4][0] = MockVoidSquare{}
	board[game.Height()-6][1] = MockVoidSquare{}
	localGame := Model{board, MockPiece{}}

	localGame.Collapse()

	if board[game.Height()-3][0].Occupied() {
		t.Errorf("(0,%d) should be free", game.Height()-3)
	}
	if !board[game.Height()-3][1].Occupied() {
		t.Errorf("(1,%d) should be occupied", game.Height()-3)
	}

	if !board[game.Height()-4][0].Occupied() {
		t.Errorf("(0,%d) should be occupied", game.Height()-4)
	}
	if board[game.Height()-4][1].Occupied() {
		t.Errorf("(1,%d) should be free", game.Height()-4)
	}
}
