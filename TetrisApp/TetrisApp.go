package TetrisApp

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"tetris/Game"
	"time"
)

type GameOver struct {
}

func (o GameOver) Error() string {
	return "game has ended"
}

type TetrisApp struct {
	time             uint
	size             float64
	leftMargin       float64
	upMargin         float64
	downMargin       float64
	interPieceMargin float64
	model            Game.IGame
}

var screenW = 1920
var screenH = 1080

func NewTetrisApp(game Game.IGame) TetrisApp {
	interPieceMargin := 1
	upMargin := 40
	downMargin := 40
	h := game.Height() - 2
	w := game.Width()
	size := float64((screenH - downMargin - upMargin - interPieceMargin*int(h)) / int(h))     //size of a square
	leftMargin := (float64(screenW) - size*float64(w) - float64(interPieceMargin*int(w))) / 2 //left margin to start the piece draw

	//ebiten.SetFullscreen(true)
	ebiten.SetWindowResizable(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetVsyncEnabled(false)
	return TetrisApp{
		0,
		size,
		leftMargin,
		float64(upMargin),
		float64(downMargin),
		float64(interPieceMargin),
		game,
	}
}

var right = false
var left = false

func (g *TetrisApp) Update(*ebiten.Image) error {
	if !ebiten.IsFocused() {
		ebiten.MinimizeWindow()
		ebiten.SetRunnableOnUnfocused(false)
	} else if !ebiten.IsRunnableOnUnfocused() {
		ebiten.SetRunnableOnUnfocused(true)
	}
	g.time++
	if g.time == 20 {
		g.time = 0
		g.model.Next()
		if g.model.CurrentPiece().IsNull() {
			if g.model.SplashLines() > 0 {
				g.model.Collapse()
			}
			if !g.model.Next() {
				return GameOver{}
			}
			//prevent key from bring pressed between pieces
			//right = false
			//left = false
		}

	}
	switch true {
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		g.model.Drop()
		break

	case ebiten.IsKeyPressed(ebiten.KeyDown) && inpututil.KeyPressDuration(ebiten.KeyDown) > 0 && inpututil.KeyPressDuration(ebiten.KeyDown)%5 == 0:
		g.model.Down()
		break
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		g.model.Down()
		break

	case left && ebiten.IsKeyPressed(ebiten.KeyLeft) && inpututil.KeyPressDuration(ebiten.KeyLeft) > 20 && inpututil.KeyPressDuration(ebiten.KeyLeft)%5 == 0:
		g.model.Left()
		println(left)
		break
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		g.model.Left()
		left = true
		break

	case right && ebiten.IsKeyPressed(ebiten.KeyRight) && inpututil.KeyPressDuration(ebiten.KeyRight) > 20 && inpututil.KeyPressDuration(ebiten.KeyRight)%5 == 0:
		g.model.Right()
		break
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		g.model.Right()
		right = true
		break
	}

	return nil
}

func (g *TetrisApp) Draw(screen *ebiten.Image) {
	time.Sleep(time.Duration(1000000))
	setScreenColor(screen, color.White)
	squares := g.model.Board()
	for y := 2; y < len(squares); y++ {
		for x := 0; x < len(squares[y]); x++ {
			ebitenutil.DrawRect(
				screen,
				float64(x)*g.size+g.leftMargin+g.interPieceMargin*float64(x),
				float64(y-2)*g.size+g.upMargin+g.interPieceMargin*float64(y),
				g.size,
				g.size,
				squares[y][x].Color(),
			)
		}
	}
}

func setScreenColor(screen *ebiten.Image, color color.Color) {
	if err := screen.Fill(color); err != nil {
	}
}

func (g *TetrisApp) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
