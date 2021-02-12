package Game

import (
	"fmt"
	"math/rand"
	"tetris/Game/Piece"
	"tetris/Game/Position"
	"tetris/Game/Square"
	"time"
)

type Model struct {
	board [][]Square.ISquare
	piece Piece.IPiece
}

func (g *Model) Drop() {
	for g.Down() {
	}
}

func (g *Model) Next() bool {
	if g.piece.IsNull() {
		rand.Seed(time.Now().UnixNano())
		switch rand.Int() % 7 {
		case 0:
			g.AddPiece(Square.PurpleSquare{}, Piece.NewTPositions(g.Width(), g.Height()))
			break

		case 1:
			g.AddPiece(Square.YellowSquare{}, Piece.NewOPositions(g.Width(), g.Height()))
			break

		case 2:
			g.AddPiece(Square.BlueSquare{}, Piece.NewJPositions(g.Width(), g.Height()))
			break

		case 3:
			g.AddPiece(Square.OrangeSquare{}, Piece.NewLPositions(g.Width(), g.Height()))
			break

		case 4:
			g.AddPiece(Square.GreenSquare{}, Piece.NewSPositions(g.Width(), g.Height()))
			break

		case 5:
			g.AddPiece(Square.RedSquare{}, Piece.NewZPositions(g.Width(), g.Height()))
			break

		case 6:
			g.AddPiece(Square.CyanSquare{}, Piece.NewIPositions(g.Width(), g.Height()))
			break
		}
	}
	if !g.Down() {
		g.piece = Piece.NullPiece{}
		return false
	}
	return true
}

func NewGame(width uint8, height uint8) IGame {
	matrix := make([][]Square.ISquare, height)
	piece := Piece.NullPiece{}
	for i := 0; i < int(height); i++ {
		matrix[i] = make([]Square.ISquare, width)
		for x := 0; x < int(width); x++ {
			matrix[i][x] = Square.VoidSquare{}
		}
	}
	return &Model{
		matrix,
		piece,
	}
}

func printMatrix(matrix [][]Square.ISquare) {
	for _, l := range matrix {
		for _, s := range l {
			if s.Occupied() {
				fmt.Print(" 1")
			} else {
				fmt.Print(" 0")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (g *Model) Collapse() {
	for line := len(g.board) - 2; line >= 0; line-- {
		if lineHasPieces(&g.board[line]) {
			l := line
			for l < len(g.board)-1 && !lineHasPieces(&g.board[l+1]) {
				g.board[l+1] = g.board[l]
				g.board[l] = *clearLine(len(g.board[l]))
				l++
			}
		}
	}
}

func lineIsFull(line *[]Square.ISquare) bool {
	for _, elem := range *line {
		if !elem.Occupied() {
			return false
		}
	}
	return true
}

func lineHasPieces(line *[]Square.ISquare) bool {
	for _, elem := range *line {
		if elem.Occupied() {
			return true
		}
	}
	return false
}

func (g *Model) HasFullLine() bool {
	for _, line := range g.board {
		if lineIsFull(&line) {
			return true
		}
	}
	return false
}

func clearLine(size int) *[]Square.ISquare {
	out := make([]Square.ISquare, size)
	for x := 0; x < size; x++ {
		out[x] = Square.VoidSquare{}
	}
	return &out
}

func (g *Model) SplashLines() uint8 {
	sum := 0
	for line := 0; line < len(g.board); line++ {
		if lineIsFull(&g.board[line]) {
			g.board[line] = *clearLine(len(g.board[line]))
			sum++
		}
	}
	return uint8(sum)
}

func inArray(array [4]Position.IPosition, element Position.IPosition) bool {
	for _, e := range array {
		if Position.Equal(e, element) {
			return true
		}
	}
	return false
}

func (g *Model) Down() bool {
	for _, p := range g.piece.Positions() {
		if g.board[p.Down().Y()][p.X()].Occupied() && !inArray(g.piece.Positions(), p.Down()) {
			return false
		}
		if p.Down().Y() == p.Y() {
			return false
		}
	}
	square := g.board[g.piece.Positions()[0].Y()][g.piece.Positions()[0].X()]
	for _, p := range g.piece.Positions() {
		g.board[p.Y()][p.X()] = Square.VoidSquare{}
	}
	for _, p := range g.piece.Positions() {
		g.board[p.Down().Y()][p.X()] = square
	}
	out := g.piece.Positions()
	for i := 0; i < 4; i++ {
		out[i] = out[i].Down()
	}
	g.piece.SetPositions(out)
	return true
}

func (g *Model) Clear() {
	for i := 0; i < len(g.board); i++ {
		g.board[i] = *clearLine(len(g.board[i]))
	}
}

func (g *Model) Right() bool {
	for _, p := range g.piece.Positions() {
		if g.board[p.Y()][p.Right().X()].Occupied() && !inArray(g.piece.Positions(), p.Right()) {
			return false
		}
		if p.Right().X() == p.X() {
			return false
		}
	}
	square := g.board[g.piece.Positions()[0].Y()][g.piece.Positions()[0].X()]
	for _, p := range g.piece.Positions() {
		g.board[p.Y()][p.X()] = Square.VoidSquare{}
	}
	for _, p := range g.piece.Positions() {
		g.board[p.Y()][p.Right().X()] = square
	}
	out := g.piece.Positions()
	for i := 0; i < 4; i++ {
		out[i] = out[i].Right()
	}
	g.piece.SetPositions(out)
	return true
}

func (g *Model) Width() uint8 {
	return uint8(len(g.board[0]))
}

func (g *Model) Height() uint8 {
	return uint8(len(g.board))
}

func (g *Model) Left() bool {
	for _, p := range g.piece.Positions() {
		if g.board[p.Y()][p.Left().X()].Occupied() && !inArray(g.piece.Positions(), p.Left()) {
			return false
		}
		if p.Left().X() == p.X() {
			return false
		}
	}
	square := g.board[g.piece.Positions()[0].Y()][g.piece.Positions()[0].X()]
	for _, p := range g.piece.Positions() {
		g.board[p.Y()][p.X()] = Square.VoidSquare{}
	}
	for _, p := range g.piece.Positions() {
		g.board[p.Y()][p.Left().X()] = square
	}
	out := g.piece.Positions()
	for i := 0; i < 4; i++ {
		out[i] = out[i].Left()
	}
	g.piece.SetPositions(out)
	return true
}

func (g *Model) AddPiece(square Square.ISquare, positions [4]Position.IPosition) {
	g.piece = Piece.NewPiece(positions)
	for _, p := range positions {
		g.board[p.Y()][p.X()] = square
	}
}

func (g *Model) CurrentPiece() Piece.IPiece {
	return g.piece
}

func (g Model) Board() [][]Square.ISquare {
	return g.board
}
