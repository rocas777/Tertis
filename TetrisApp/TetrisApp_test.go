package TetrisApp

import (
	"testing"
	"tetris/Game"
)

func TestNewTetrisApp(t *testing.T) {
	app := NewTetrisApp(Game.NewGame(10, 10))
	if app.leftMargin != 335 {
		t.Errorf("Left margin shoud be 335,instead found %f", app.leftMargin)
	}
	if app.size != 124 {
		t.Errorf("Size shoud be 124,instead found %f", app.size)
	}
}
