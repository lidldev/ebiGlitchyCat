package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1000
	screenHeight = 600
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
	g.camera.clear()

	g.background.intialBackground(screen)

	g.camera.draw(assets.MainSprite, &ebiten.DrawImageOptions{})
	g.camera.render(screen)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.x += 10
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.x -= 10
	}
	g.player.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
