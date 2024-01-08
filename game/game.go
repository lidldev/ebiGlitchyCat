package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player     Player
	camera     camera
	background Background
}

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
	g.camera.setPos(-450, -300)
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.intialBackground(screen)
	g.camera.clear()
	g.player.Draw(screen)

	g.camera.render(screen)
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
