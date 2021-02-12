package Game

import (
	"testing"
	"tetris/Game/Piece"
)

func Test_NewTPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewTPositions(game.Width(), game.Height()))
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][5].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}

func Test_NewOPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewOPositions(game.Width(), game.Height()))
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][5].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][5].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}

func Test_NewJPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewJPositions(game.Width(), game.Height()))
	if !game.Board()[0][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][5].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}
func Test_NewLPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewLPositions(game.Width(), game.Height()))
	if !game.Board()[0][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][5].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][3].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}

func Test_NewSPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewSPositions(game.Width(), game.Height()))
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][5].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][4].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}

func Test_NewZPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewZPositions(game.Width(), game.Height()))
	if !game.Board()[0][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[1][5].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}

func Test_NewIPiece(t *testing.T) {
	game.Clear()
	game.AddPiece(MockColourSquare{}, Piece.NewIPositions(game.Width(), game.Height()))
	if !game.Board()[0][3].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][4].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][5].Occupied() {
		t.Errorf("")
	}
	if !game.Board()[0][6].Occupied() {
		t.Errorf("")
	}
	if countNumberOfOccupied(t, game.Board()) != 4 {
		t.Errorf("Should only exist 4 occupied squares")
	}
}
