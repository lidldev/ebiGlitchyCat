package game

import (
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	screen.DrawImage(assets.MainSprite, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
