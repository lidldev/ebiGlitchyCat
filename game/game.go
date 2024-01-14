package game

import (
	"main/assets"

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
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.8, 0.8)
	op.GeoM.Translate(float64(g.player.player.x)/unit, float64(g.player.player.y)/unit)
	g.camera.draw(assets.MainSprite, op)
	screen.DrawImage(assets.MainSprite, &ebiten.DrawImageOptions{})
	//g.player.Draw(screen)
	g.camera.render(screen)
}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-300, g.player.player.y/unit-400)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
