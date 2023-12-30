package game

import (
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x int
	y int

	camera *camera
}

func NewPlayer() *Player {
	p := &Player{}

	return p
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.camera.draw(assets.MainSprite, &ebiten.DrawImageOptions{})
	p.camera.render(screen)
}
