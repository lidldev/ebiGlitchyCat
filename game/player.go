package game

import "github.com/hajimehoshi/ebiten/v2"

type char struct {
	x  int
	y  int
	vx int
	vy int
}

func (c *char) update() {

}

type player struct {
	player *char
}

func NewPlayer() *player {
	p := &player{}

	return p
}

func (p *player) Update() error {
	return nil
}

func (p *player) Draw(screen *ebiten.Image) {

}
