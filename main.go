package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"tetris/Game"
	"tetris/TetrisApp"
)

func main() {
	game := Game.NewGame(10, 25)
	app := TetrisApp.NewTetrisApp(game)

	images := make([]image.Image, 1)
	_, file, _ := ebitenutil.NewImageFromFile("icon.png", ebiten.FilterLinear)
	images[0] = file
	if file != nil {
		ebiten.SetWindowIcon(images)
	}

	if err := ebiten.RunGame(&app); err != nil {
		if (err == TetrisApp.GameOver{}) {
			return
		}
	}
}
